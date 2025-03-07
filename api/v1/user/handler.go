package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gin-app/errors"
	"gin-app/models"
	"gin-app/responses"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userRepo models.UserRepository
}

// NewUserHandler creates a new UserHandler with the provided repository
func NewUserHandler(userRepo models.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// CreateUserRequest represents the request body for creating a user
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
	Username string `json:"username" binding:"omitempty,min=3,max=50"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password" binding:"omitempty,min=6"`
}

// RegisterRoutes registers all user-related routes
func (h *UserHandler) RegisterRoutes(router *gin.RouterGroup) {
	users := router.Group("/users")
	{
		users.POST("", h.CreateUser)
		users.GET("", h.GetAllUsers)
		users.GET("/:id", h.GetUser)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
	}
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User information"
// @Success 201 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Failure 409 {object} responses.Response
// @Failure 500 {object} responses.Response
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body: "+err.Error())
		return
	}

	// Check if username already exists
	_, err := h.userRepo.GetByUsername(req.Username)
	if err == nil {
		responses.Conflict(c, "Username already taken")
		return
	}

	// Check if email already exists
	_, err = h.userRepo.GetByEmail(req.Email)
	if err == nil {
		responses.Conflict(c, "Email already in use")
		return
	}

	// Create user
	user := &models.User{
		ID:        uuid.New().String(),
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password, // In a real application, you would hash this
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.userRepo.Create(user); err != nil {
		responses.InternalServerError(c, "Failed to create user: "+err.Error())
		return
	}

	responses.Created(c, "User created successfully", user)
}

// GetUser retrieves a user by ID
// @Summary Get a user by ID
// @Description Get a user by their ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} responses.Response
// @Failure 404 {object} responses.Response
// @Failure 500 {object} responses.Response
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userRepo.GetByID(id)
	if err != nil {
		appErr := errors.NotFound("User not found", map[string]string{"id": id})
		responses.Error(c, appErr.StatusCode, appErr.Message, appErr.Details)
		return
	}

	responses.Success(c, "User retrieved successfully", user)
}

// GetAllUsers retrieves all users
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {object} responses.Response
// @Failure 500 {object} responses.Response
// @Router /api/v1/users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userRepo.GetAll()
	if err != nil {
		responses.InternalServerError(c, "Failed to retrieve users: "+err.Error())
		return
	}

	responses.Success(c, "Users retrieved successfully", users)
}

// UpdateUser updates an existing user
// @Summary Update a user
// @Description Update a user's information
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body UpdateUserRequest true "User information"
// @Success 200 {object} responses.Response
// @Failure 400 {object} responses.Response
// @Failure 404 {object} responses.Response
// @Failure 409 {object} responses.Response
// @Failure 500 {object} responses.Response
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	// Check if user exists
	user, err := h.userRepo.GetByID(id)
	if err != nil {
		appErr := errors.NotFound("User not found", map[string]string{"id": id})
		responses.Error(c, appErr.StatusCode, appErr.Message, appErr.Details)
		return
	}

	// Bind request
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.BadRequest(c, "Invalid request body: "+err.Error())
		return
	}

	// Update user fields if provided
	if req.Username != "" {
		// Check if username is already taken by another user
		if user.Username != req.Username {
			existingUser, err := h.userRepo.GetByUsername(req.Username)
			if err == nil && existingUser.ID != id {
				responses.Conflict(c, "Username already taken")
				return
			}
		}
		user.Username = req.Username
	}

	if req.Email != "" {
		// Check if email is already in use by another user
		if user.Email != req.Email {
			existingUser, err := h.userRepo.GetByEmail(req.Email)
			if err == nil && existingUser.ID != id {
				responses.Conflict(c, "Email already in use")
				return
			}
		}
		user.Email = req.Email
	}

	if req.Password != "" {
		// In a real application, you would hash this password
		user.Password = req.Password
	}

	user.UpdatedAt = time.Now()

	// Save updated user
	if err := h.userRepo.Update(user); err != nil {
		responses.InternalServerError(c, "Failed to update user: "+err.Error())
		return
	}

	responses.Success(c, "User updated successfully", user)
}

// DeleteUser deletes a user by ID
// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags users
// @Param id path string true "User ID"
// @Success 204 "No Content"
// @Failure 404 {object} responses.Response
// @Failure 500 {object} responses.Response
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	// Check if user exists
	_, err := h.userRepo.GetByID(id)
	if err != nil {
		appErr := errors.NotFound("User not found", map[string]string{"id": id})
		responses.Error(c, appErr.StatusCode, appErr.Message, appErr.Details)
		return
	}

	// Delete user
	if err := h.userRepo.Delete(id); err != nil {
		responses.InternalServerError(c, "Failed to delete user: "+err.Error())
		return
	}

	responses.NoContent(c)
}

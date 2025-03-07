package models

import (
	"errors"
	"sync"
	"time"
)

// User represents a user in the system
type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Password is not exposed in JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository defines the interface for User data operations
type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetAll() ([]*User, error)
	Update(user *User) error
	Delete(id string) error
}

// InMemoryUserRepository implements the UserRepository interface with in-memory storage
type InMemoryUserRepository struct {
	users map[string]*User
	mutex sync.RWMutex
}

// NewInMemoryUserRepository creates a new instance of InMemoryUserRepository
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*User),
	}
}

// Create adds a new user to the repository
func (r *InMemoryUserRepository) Create(user *User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Check if user with same ID already exists
	if _, exists := r.users[user.ID]; exists {
		return errors.New("user with this ID already exists")
	}

	// Check if username is already taken
	for _, u := range r.users {
		if u.Username == user.Username {
			return errors.New("username already taken")
		}
		if u.Email == user.Email {
			return errors.New("email already in use")
		}
	}

	// Set timestamps
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Store user
	r.users[user.ID] = user
	return nil
}

// GetByID retrieves a user by ID
func (r *InMemoryUserRepository) GetByID(id string) (*User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// GetByUsername retrieves a user by username
func (r *InMemoryUserRepository) GetByUsername(username string) (*User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetByEmail retrieves a user by email
func (r *InMemoryUserRepository) GetByEmail(email string) (*User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetAll retrieves all users
func (r *InMemoryUserRepository) GetAll() ([]*User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]*User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

// Update updates an existing user
func (r *InMemoryUserRepository) Update(user *User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return errors.New("user not found")
	}

	// Check if username or email conflicts with another user
	for id, u := range r.users {
		if id == user.ID {
			continue
		}
		if u.Username == user.Username {
			return errors.New("username already taken")
		}
		if u.Email == user.Email {
			return errors.New("email already in use")
		}
	}

	// Update timestamp
	user.UpdatedAt = time.Now()

	// Update user
	r.users[user.ID] = user
	return nil
}

// Delete removes a user by ID
func (r *InMemoryUserRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(r.users, id)
	return nil
}

package models

import (
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
    repo := NewInMemoryUserRepository()
    
    user := &User{
        ID:       "1", 
        Username: "testuser", 
        Email:    "test@example.com",
        Password: "password123",
    }
    
    err := repo.Create(user)
    assert.Nil(t, err)
    
    // Check created user has timestamps set
    assert.False(t, user.CreatedAt.IsZero())
    assert.False(t, user.UpdatedAt.IsZero())
    
    // Verify user is stored
    storedUser, err := repo.GetByID("1")
    assert.Nil(t, err)
    assert.Equal(t, user.ID, storedUser.ID)
    assert.Equal(t, user.Username, storedUser.Username)
    assert.Equal(t, user.Email, storedUser.Email)
}

func TestCreateUserDuplicateID(t *testing.T) {
    repo := NewInMemoryUserRepository()
    
    user1 := &User{ID: "1", Username: "user1", Email: "user1@example.com"}
    user2 := &User{ID: "1", Username: "user2", Email: "user2@example.com"}
    
    err := repo.Create(user1)
    assert.Nil(t, err)
    
    err = repo.Create(user2)
    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "already exists")
}

func TestCreateUserDuplicateUsername(t *testing.T) {
    repo := NewInMemoryUserRepository()
    
    user1 := &User{ID: "1", Username: "testuser", Email: "user1@example.com"}
    user2 := &User{ID: "2", Username: "testuser", Email: "user2@example.com"}
    
    err := repo.Create(user1)
    assert.Nil(t, err)
    
    err = repo.Create(user2)
    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "username already taken")
}

func TestCreateUserDuplicateEmail(t *testing.T) {
    repo := NewInMemoryUserRepository()
    
    user1 := &User{ID: "1", Username: "user1", Email: "test@example.com"}
    user2 := &User{ID: "2", Username: "user2", Email: "test@example.com"}
    
    err := repo.Create(user1)
    assert.Nil(t, err)
    
    err = repo.Create(user2)
    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "email already in use")
}

func TestGetUserByID(t *testing.T) {
    repo := NewInMemoryUserRepository()
    
    // Create test user
    user := &User{ID: "1", Username: "testuser", Email: "test@example.com"}
    _ = repo.Create(user)
    
    // Test successful retrieval
    found, err := repo.GetByID("1")
    assert.Nil(t, err)
    assert.Equal(t, user.ID, found.ID)
    
    // Test non-existent user
    _, err = repo.GetByID("999")
    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "not found")
}

func TestGetUserByUsername(t *testing.T) {
    repo := NewInMemoryUserRepository()
    
    // Create test user
    user := &User{ID: "1", Username: "testuser", Email: "test@example.com"}
    _ = repo.Create(user)
    
    // Test successful retrieval
    found, err := repo.GetByUsername("testuser")
    assert.Nil(t, err)
    assert.Equal(t, user.ID, found.ID)
    
    // Test non-existent user
    _, err = repo.GetByUsername("nonexistent")
    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "not found")
}

func TestGetUserByEmail(t *testing.T) {
    repo := NewInMemoryUserRepository()
    
    // Create test user
    user := &User{ID: "1", Username: "testuser", Email: "test@example.com"}
    _ = repo.Create(user)
    
    // Test successful retrieval
    found, err := repo.GetByEmail("test@example.com")
    assert.Nil(t, err)
    assert.Equal(t, user.ID, found.ID)
    
    // Test non-existent user
    _, err = repo.GetByEmail("nonexistent@example.com")
    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "not found")
}

func TestGetAllUsers(t *testing.T) {
    repo := NewInMemoryUserRepository()
    
    // Test empty repository
    users, err := repo.GetAll()
    assert.Nil(t, err)
    assert.Len(t, users, 0)
    
    // Add users
    user1 := &User{ID: "1", Username: "user1", Email: "user1@example.com"}
    user2 := &User{ID: "2", Username: "user2", Email: "user2@example.com"}
    
    _ = repo.Create(user1)
    _ = repo.Create(user2)
    
    // Test retrieval of all users
    users, err = repo.GetAll()
    assert.Nil(t, err)
    assert.Len(t, users, 2)
}

func TestUpdateUser(t *testing.T) {
    repo := NewInMemoryUserRepository()
    
    // Create test user
    originalTime := time.Now().Add(-1 * time.Hour) // 1 hour ago
    user := &User{
        ID:        "1", 
        Username:  "testuser", 
        Email:     "test@example.com",
        CreatedAt: originalTime,
        UpdatedAt: originalTime,
    }
    _ = repo.Create(user)
    
    // Update user
    updatedUser := &User{
        ID:        "1",
        Username:  "updateduser",
        Email:     "updated@example.com",
        CreatedAt: originalTime, // CreatedAt should not change
        UpdatedAt: originalTime, // UpdatedAt will be updated by the repo
    }
    
    err := repo.Update(updatedUser)
    assert.Nil(t, err)
    
    // Verify update
    found, _ := repo.GetByID("1")
    assert.Equal(t, "updateduser", found.Username)
    assert.Equal(t, "updated@example.com", found.Email)
    assert.Equal(t, originalTime, found.CreatedAt)
    assert.True(t, found.UpdatedAt.After(originalTime))
}

func TestUpdateNonExistentUser(t *testing.T) {
    repo := NewInMemoryUserRepository()
    
    user := &User{ID: "999", Username: "testuser", Email: "test@example.com"}
    
    err := repo.Update(user)
    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "not found")
}

func TestUpdateUserDuplicateUsername(t *testing.T) {
    repo := NewInMemoryUserRepository()
    
    user1 := &User{ID: "1", Username: "user1", Email: "user1@example.com"}
    user2 := &User{ID: "2", Username: "user2", Email: "user2@example.com"}
    
    _ = repo.Create(user1)
    _ = repo.Create(user2)
    
    // Try to update user2 with username that already exists in user1
    user2.Username = "user1"
    
    err := repo.Update(user2)
    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "username already taken")
}

func TestDeleteUser(t *testing.T) {
    repo := NewInMemoryUserRepository()
    
    // Create test user
    user := &User{ID: "1", Username: "testuser", Email: "test@example.com"}
    _ = repo.Create(user)
    
    // Delete user
    err := repo.Delete("1")
    assert.Nil(t, err)
    
    // Verify deletion
    _, err = repo.GetByID("1")
    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "not found")
}

func TestDeleteNonExistentUser(t *testing.T) {
    repo := NewInMemoryUserRepository()
    
    err := repo.Delete("999")
    assert.NotNil(t, err)
    assert.Contains(t, err.Error(), "not found")
}

package domain

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidEmail    = errors.New("invalid email format")
	ErrInvalidUsername = errors.New("invalid username format")
	ErrWeakPassword    = errors.New("password does not meet requirements")
	ErrUserInactive    = errors.New("user account is inactive")
)

// User represents a user entity in the domain
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` // Never expose password in JSON
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserRepository defines the interface for user persistence
type UserRepository interface {
	Save(user *User) error
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	Update(user *User) error
	Delete(id string) error
}

// UserService defines the interface for user business logic
type UserService interface {
	CreateUser(email, username, password, firstName, lastName string) (*User, error)
	AuthenticateUser(email, password string) (*User, error)
	GetUserByID(id string) (*User, error)
	UpdateUser(id string, updates map[string]interface{}) (*User, error)
	DeleteUser(id string) error
}

// NewUser creates a new user with validation
func NewUser(email, username, password, firstName, lastName string) (*User, error) {
	user := &User{
		ID:        uuid.New().String(),
		Email:     email,
		Username:  username,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

// Validate validates the user data
func (u *User) Validate() error {
	if !isValidEmail(u.Email) {
		return ErrInvalidEmail
	}

	if !isValidUsername(u.Username) {
		return ErrInvalidUsername
	}

	if !isValidPassword(u.Password) {
		return ErrWeakPassword
	}

	return nil
}

// isValidEmail validates email format
func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// isValidUsername validates username format
func isValidUsername(username string) bool {
	if len(username) < 3 || len(username) > 20 {
		return false
	}
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return usernameRegex.MatchString(username)
}

// isValidPassword validates password strength
func isValidPassword(password string) bool {
	return len(password) >= 8
}

// SetPassword sets a new password for the user
func (u *User) SetPassword(password string) error {
	if !isValidPassword(password) {
		return ErrWeakPassword
	}
	u.Password = password
	u.UpdatedAt = time.Now()
	return nil
}

// Activate activates the user account
func (u *User) Activate() {
	u.IsActive = true
	u.UpdatedAt = time.Now()
}

// Deactivate deactivates the user account
func (u *User) Deactivate() {
	u.IsActive = false
	u.UpdatedAt = time.Now()
}

// IsActivated returns true if the user is active
func (u *User) IsActivated() bool {
	return u.IsActive
}

package user

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                     uuid.UUID  `json:"id" redis:"id"`
	FirstName              string     `json:"first_name" redis:"first_name"`
	LastName               string     `json:"last_name" redis:"last_name"`
	UserName               string     `json:"user_name" redis:"user_name"`
	Email                  string     `json:"email" redis:"email"`
	EmailVerified          bool       `json:"email_verified" redis:"email_verified"`
	Password               string     `json:"password" redis:"password"`
	UserImage              UserImage  `json:"user_image" redis:"user_image"`
	VerificationTokenHash  string     `json:"-" redis:"verification_token_hash"`
	VerificationExpiresAt  time.Time  `json:"-" redis:"verification_expires_at"`
	PasswordResetTokenHash string     `json:"-" redis:"password_reset_token_hash"`
	PasswordResetExpiresAt time.Time  `json:"-" redis:"password_reset_expires_at"`
	LastLogin              *time.Time `json:"last_login,omitempty" redis:"last_login"`
	FailedLoginAttempts    int        `json:"failed_login_attempts" redis:"failed_login_attempts"`
	LockedUntil            *time.Time `json:"locked_until,omitempty" redis:"locked_until"`
	CreatedAt              time.Time  `json:"created_at" redis:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at" redis:"updated_at"`
}

// SanitizePassword removes the password from the user struct for security reasons.
func (u *User) SanitizePassword() {
	u.Password = ""
}

// PrepareCreate processes the user data before creating a new user record.
func (u *User) PrepareCreate() {
	// Trim whitespace
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

// PrepareUpdate processes the user data before updating an existing user record.
func (u *User) PrepareUpdate() {
	// Trim whitespace and normalize
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Password = strings.TrimSpace(u.Password)
	if u.Email != "" {
		u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	}

	u.UpdatedAt = time.Now()
}

type AuthResponse struct {
	User        User     `json:"user"`
	Permissions []string `json:"permissions,omitempty"`
}

// Role represents an RBAC role record.
type Role struct {
	ID          int       `json:"id" redis:"id"`
	Name        string    `json:"name" redis:"name"`
	Description string    `json:"description,omitempty" redis:"description"`
	CreatedAt   time.Time `json:"created_at" redis:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" redis:"updated_at"`
}

// Permission represents a single permission that can be assigned to roles.
type Permission struct {
	ID          int       `json:"id" redis:"id"`
	Name        string    `json:"name" redis:"name"`
	Description string    `json:"description,omitempty" redis:"description"`
	CreatedAt   time.Time `json:"created_at" redis:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" redis:"updated_at"`
}

// RolePermission is the pivot between roles and permissions (many-to-many).
type RolePermission struct {
	RoleID       int `json:"role_id" redis:"role_id"`
	PermissionID int `json:"permission_id" redis:"permission_id"`
}

// UserRole represents a user's assignment to a role (many-to-many join table).
type UserRole struct {
	UserID uuid.UUID `json:"user_id" redis:"user_id"`
	RoleID int       `json:"role_id" redis:"role_id"`
}

// UserImage represents a user's profile image stored in the system.
type UserImage struct {
	PublicId string    `json:"publicId" redis:"publicId"`
	Url      string    `json:"url" redis:"url"`
	UserId   uuid.UUID `json:"userId" redis:"userId"`
}

package user

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// RedisRepository describes the Redis-backed cache / index operations
// used by the user domain. Implementations live in the infrastructure
// layer and must avoid embedding business logic — only cache/index
// operations and short-lived token mappings belong here.
type RedisRepository interface {
	// Basic user cache
	SetUser(ctx context.Context, u *User, ttl time.Duration) error
	GetUserByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error

	// Verification token mapping (store only token hash -> userID)
	SetVerificationTokenHash(ctx context.Context, userID uuid.UUID, tokenHash string, ttl time.Duration) error
	GetUserIDByVerificationTokenHash(ctx context.Context, tokenHash string) (uuid.UUID, error)
	DeleteVerificationTokenHash(ctx context.Context, tokenHash string) error

	// Password reset token mapping (store only token hash -> userID)
	SetPasswordResetTokenHash(ctx context.Context, userID uuid.UUID, tokenHash string, ttl time.Duration) error
	GetUserIDByPasswordResetTokenHash(ctx context.Context, tokenHash string) (uuid.UUID, error)
	DeletePasswordResetTokenHash(ctx context.Context, tokenHash string) error

	// Permissions cache (precomputed permissions for fast RBAC checks)
	SetPermissions(ctx context.Context, userID uuid.UUID, permissions []string, ttl time.Duration) error
	GetPermissions(ctx context.Context, userID uuid.UUID) ([]string, error)
	DeletePermissions(ctx context.Context, userID uuid.UUID) error

	// Per-user sessions index (helps implement single-sign-out and list sessions)
	AddSessionForUser(ctx context.Context, userID uuid.UUID, sessionID string, ttl time.Duration) error
	RemoveSessionForUser(ctx context.Context, userID uuid.UUID, sessionID string) error
	ListSessionsForUser(ctx context.Context, userID uuid.UUID) ([]string, error)
	DeleteAllSessionsForUser(ctx context.Context, userID uuid.UUID) error

	// Invalidate all caches related to a user (composite helper)
	InvalidateUserCache(ctx context.Context, userID uuid.UUID) error
}

// NOTE: implementations should favor simple, atomic Redis operations
// (SETEX, HSET, SADD, SREM, EXPIRE, DEL) and keep TTLs short for
// sensitive mappings (tokens, permissions). Do not persist long-lived
// secrets in Redis; store only hashed values for tokens.

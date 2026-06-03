package session

import "github.com/google/uuid"


type Session struct {
	ID        string `json:"id" redis:"id"`
	UserID uuid.UUID `json:"user_id" redis:"user_id"`
	Roles	 []string `json:"roles,omitempty" redis:"roles"`
	Permissions []string `json:"permissions,omitempty" redis:"permissions"`
	IssuedAt  int64  `json:"issued_at" redis:"issued_at"`
	ExpiresAt int64  `json:"expires_at" redis:"expires_at"`
}
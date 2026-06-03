package session

import "time"


type Repo interface {
	Create(ctx, s *Session, ttl time.Duration) error
	Get(ctx, id string) (*Session, error)
	Delete(ctx, id string) error
	Refresh(ctx, id string, ttl time.Duration) error
}
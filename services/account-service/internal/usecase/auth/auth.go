package auth

import (
	"github.com/jofosuware/bizhelm/account-service/internal/domain/user"
	"github.com/jofosuware/bizhelm/config"
	"github.com/jofosuware/bizhelm/logger"
)


type authenticationUC struct {
	cfg *config.Config
	userRepo user.Repository
	redisRepo user.RedisRepository
	logger logger.Logger
}

func NewAuthencationUC(cfg *config.Config, userRepo user.Repository, redisRepo user.RedisRepository, logger logger.Logger) *authenticationUC {
	return &authenticationUC{
		cfg: cfg,
		userRepo: userRepo,
		redisRepo: redisRepo,
		logger: logger,
	}
}
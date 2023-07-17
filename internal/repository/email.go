package repository

import (
	"namecard-gateway/internal/config"
	"namecard-gateway/internal/infrastructure/database"
	"namecard-gateway/internal/util/log"
)

type EmailRepository struct {
	config *config.Configuration
	db     *database.DB
	logger *log.Logger
}

type IEmailRepository interface {
}

func NewEmailRepository(
	config *config.Configuration,
	db *database.DB,
	logger *log.Logger,
) IEmailRepository {
	return &EmailRepository{
		config: config,
		logger: logger,
		db:     db,
	}
}

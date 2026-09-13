package services

import (
	"user-service/internal/config"
	"user-service/internal/mailer"
)

type service struct {
	repository repository
	mailer     mailer.Mailer
	cfg        config.Config
}

func NewUserService(repo repository, mail mailer.Mailer, config config.Config) *service {
	return &service{
		repository: repo,
		mailer:     mail,
		cfg:        config,
	}
}

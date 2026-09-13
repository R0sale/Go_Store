package services

import (
	"sync"
	"user-service/internal/config"
	"user-service/internal/mailer"
)

type service struct {
	repository repository
	mailer     mailer.Mailer
	cfg        config.Config
	wg         *sync.WaitGroup
}

func NewUserService(repo repository, mail mailer.Mailer, config config.Config, wg *sync.WaitGroup) *service {
	return &service{
		repository: repo,
		mailer:     mail,
		cfg:        config,
		wg:         wg,
	}
}

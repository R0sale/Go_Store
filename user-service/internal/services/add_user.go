package services

import (
	"context"
	"fmt"
	"user-service/internal/dto"
	"user-service/internal/models"
)

func (s service) AddUser(ctx context.Context, user dto.CreateUserDto) error {
	repoUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		ImageUrl: user.ImageUrl,
		Password: user.Password,
	}

	err := s.repository.AddUser(ctx, repoUser)
	if err != nil {
		return err
	}

	s.wg.Add(1)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("couldnt send the email err: %s\n", err)
			}
		}()

		defer s.wg.Done()

		err := s.mailer.Send(user.Email, "user_welcome.tmpl", user)
		if err != nil {
			fmt.Printf("couldnt send the email err: %s\n", err.Error())
		}
	}()

	return nil
}

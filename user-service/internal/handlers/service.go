package handlers

import "user-service/internal/dto"

type service interface {
	AddUser(user dto.CreateUserDto) error
}

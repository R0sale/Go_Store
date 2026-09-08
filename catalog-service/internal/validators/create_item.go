package validators

import (
	"catalog-service/internal/dto"

	"github.com/go-playground/validator/v10"
)

type ItemValidator struct {
	validate *validator.Validate
}

func (v *ItemValidator) ValidateCreateItem(item dto.CreateItemDto) error {
	return v.validate.Struct(item)
}

func NewItemValidator(validator *validator.Validate) *ItemValidator {
	return &ItemValidator{
		validate: validator,
	}
}

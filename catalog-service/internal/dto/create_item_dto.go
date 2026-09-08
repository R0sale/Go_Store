package dto

type CreateItemDto struct {
	Name     string  `validate:"required,min=2,max=50"`
	Price    float32 `validate:"gt=0"`
	ImageUrl string  `validate:"omitempty,url"`
}

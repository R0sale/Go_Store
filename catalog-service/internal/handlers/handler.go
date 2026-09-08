package handlers

import "catalog-service/internal/validators"

type catalogHandler struct {
	service       Service
	itemValidator *validators.ItemValidator
}

func NewCatalogHandler(serv Service, validator *validators.ItemValidator) *catalogHandler {
	return &catalogHandler{
		service:       serv,
		itemValidator: validator,
	}
}

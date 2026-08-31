package handlers

type catalogHandler struct {
	service Service
}

func NewCatalogHandler(serv Service) *catalogHandler {
	return &catalogHandler{
		service: serv,
	}
}

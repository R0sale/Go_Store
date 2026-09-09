package handlers

type userHandler struct {
	service service
}

func NewUserHandler(serv service) *userHandler {
	return &userHandler{
		service: serv,
	}
}

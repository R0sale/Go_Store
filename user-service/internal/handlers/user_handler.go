package handlers

import (
	"net/http"
)

type userHandler struct {
	service service
}

func NewUserHandler(serv service) *userHandler {
	return &userHandler{
		service: serv,
	}
}

func (h userHandler) AddJwtHttpOnlyCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		MaxAge:   3600,
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	http.SetCookie(w, &cookie)
}

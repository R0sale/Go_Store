package handlers

import (
	"net/http"
	"time"
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
		MaxAge:   time.Now().Hour(),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
}

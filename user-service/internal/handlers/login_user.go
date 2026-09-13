package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"user-service/internal/dto"
)

func (h userHandler) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	var user dto.LoginUserDto

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbUser, token, err := h.service.LoginUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		MaxAge:   time.Now().Hour(),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbUser)
}

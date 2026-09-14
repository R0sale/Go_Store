package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"user-service/internal/dto"
)

func (h userHandler) HandleAddUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserDto

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	newUser, token, err := h.service.AddUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.AddJwtHttpOnlyCookie(w, token)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

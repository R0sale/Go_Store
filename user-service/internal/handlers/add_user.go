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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = h.service.AddUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

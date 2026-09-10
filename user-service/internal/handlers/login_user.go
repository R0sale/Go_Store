package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"user-service/internal/dto"
	"user-service/internal/responses"
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

	token, err := h.service.LoginUser(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses.Token{Token: token})
}

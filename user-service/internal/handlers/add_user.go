package handlers

import (
	"encoding/json"
	"net/http"
	"user-service/internal/dto"
)

func (h userHandler) HandleAddUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserDto

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.AddUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

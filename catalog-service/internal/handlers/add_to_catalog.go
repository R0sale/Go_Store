package handlers

import (
	"catalog-service/internal/dto"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (h catalogHandler) HandleAddToCatalog(w http.ResponseWriter, r *http.Request) {
	var item dto.CreateItemDto

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.itemValidator.ValidateCreateItem(item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = h.service.AddToCatalog(ctx, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

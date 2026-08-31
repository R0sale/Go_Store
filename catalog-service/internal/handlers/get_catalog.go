package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func (h catalogHandler) HandleGetCatalog(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("userId")

	userId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Couldn't parse user id. It must be integer", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	items, err := h.service.GetCatalog(ctx, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

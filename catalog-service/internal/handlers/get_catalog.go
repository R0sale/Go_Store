package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (h catalogHandler) HandleGetCatalog(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	items, err := h.service.GetCatalog(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

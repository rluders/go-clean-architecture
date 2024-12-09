package task

import (
	"encoding/json"
	"net/http"
)

type RestHandler struct {
	service *Service
}

func NewRestHandler(service *Service) *RestHandler {
	return &RestHandler{service: service}
}

func (h *RestHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.service.CreateTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(task)
}

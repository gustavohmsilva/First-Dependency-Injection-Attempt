package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gustavohmsilva/test-dependency-injection/model"
	"github.com/gustavohmsilva/test-dependency-injection/service"
)

// Handler ...
type Handler struct {
	us service.UserService
}

// NewHandler ...
func NewHandler(us service.UserService) Handler {
	return Handler{us: us}
}

// GetLatestUser ...
func (h Handler) GetLatestUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	latestUser, err := h.us.GetLatestUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	luj, err := json.Marshal(latestUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(luj)
}

// SetLatestUser ...
func (h Handler) SetLatestUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.us.InsertUser(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

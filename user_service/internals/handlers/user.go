package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"user_service/internals/models"
	"user_service/internals/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "could not decode json", err.Error())
		return
	}

	err := h.service.CreateUser(context.TODO(), user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not create user:", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) FindUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not parse id:", err)
	}

	user, err := h.service.FindUser(context.TODO(), id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "could not find user:", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

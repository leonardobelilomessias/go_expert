package handlers

import (
	"encoding/json"
	"net/http"

	"goexpert.com/module07/internal/dto"
	"goexpert.com/module07/internal/entity"
	"goexpert.com/module07/internal/infra/database"
)

type userHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *userHandler {
	return &userHandler{UserDB: userDB}
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

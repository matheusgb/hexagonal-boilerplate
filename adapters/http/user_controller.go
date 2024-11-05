package http

import (
	"encoding/json"
	"hexagonal/adapters/http/dtos"
	"hexagonal/application"
	domain "hexagonal/domain/entities"
	"net/http"
)

type UserController struct {
	service *application.UserService
}

func NewUserController(service *application.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var userDTO dtos.UserInputDTO
	if err := json.NewDecoder(r.Body).Decode(&userDTO); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	user := domain.User{
		ID:    "uuid",
		Name:  userDTO.Name,
		Email: userDTO.Email,
	}

	if err := c.service.Create(user); err != nil {
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (c *UserController) GetById(w http.ResponseWriter, r *http.Request) {
}

func (c *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := c.service.GetAll()
	if err != nil {
		http.Error(w, "error fetching users", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(users)
}

package main

import (
	controller "hexagonal/adapters/http"
	"hexagonal/adapters/in_memory"
	"hexagonal/application"
	"log"
	"net/http"
)

func main() {
	userRepo := in_memory.NewInMemoryUserRepo()
	userService := application.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	http.HandleFunc("/users", userController.GetAll)
	http.HandleFunc("/users/create", userController.Create)
	// Adicione manipuladores para as outras operações

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

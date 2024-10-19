package main

import (
	"log"
	"net/http"

	"user_service/cmd/config"
	"user_service/internals/db"
	"user_service/internals/handlers"
	"user_service/internals/repository"
	"user_service/internals/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Could not load config: ", err)
	}

	mongoClient, err := db.NewMongoClient(cfg.MongoURI, cfg.DatabaseName)

	if err != nil {
		log.Fatal("Could not connect to mongodb: ", err)
	}

	userRepo := repository.NewUserRepository(mongoClient, cfg.DatabaseName)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	http.HandleFunc("POST /users", userHandler.CreateUser)
	http.HandleFunc("GET /user/{id}", userHandler.FindUser)

	log.Println("Server started at :80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}

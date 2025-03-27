package main

import (
	"log"
	"merchant-bank-api/api"
	"merchant-bank-api/repository"
	"merchant-bank-api/service"
	"net/http"
)

func main() {
	repo := &repository.Repository{}
	svc := service.NewService(repo)
	api := api.NewAPI(svc)
	router := api.SetupRoutes()

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

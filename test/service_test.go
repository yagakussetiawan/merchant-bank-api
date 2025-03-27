package test

import (
	"merchant-bank-api/repository"
	"merchant-bank-api/service"
	"testing"
)

func TestLogin(t *testing.T) {
	repo := &repository.Repository{}
	svc := service.NewService(repo)
	token, err := svc.Login("customer1", "pass123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if token == "" {
		t.Fatal("Expected token, got empty string")
	}
}

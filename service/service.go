package service

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"merchant-bank-api/models"
	"merchant-bank-api/repository"
	"time"
)

type Service struct {
	repo     *repository.Repository
	sessions map[string]models.Session // In-memory session store
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo, sessions: make(map[string]models.Session)}
}

func (s *Service) Login(username, password string) (string, error) {
	customers, err := s.repo.GetCustomers()
	if err != nil {
		return "", err
	}
	for _, c := range customers {
		if c.Username == username && c.Password == password {
			token := generateToken()
			s.sessions[token] = models.Session{CustomerID: c.ID, Token: token}
			s.repo.SaveHistory(models.History{
				ID:         generateID(),
				CustomerID: c.ID,
				Action:     "login",
				Timestamp:  time.Now().Format(time.RFC3339),
			})
			return token, nil
		}
	}
	return "", errors.New("invalid credentials")
}

func (s *Service) Payment(token string, amount int) error {
	session, exists := s.sessions[token]
	if !exists {
		return errors.New("unauthorized")
	}
	return s.repo.SaveHistory(models.History{
		ID:         generateID(),
		CustomerID: session.CustomerID,
		Action:     "payment",
		Amount:     amount,
		Timestamp:  time.Now().Format(time.RFC3339),
	})
}

func (s *Service) Logout(token string) error {
	session, exists := s.sessions[token]
	if !exists {
		return errors.New("unauthorized")
	}
	delete(s.sessions, token)
	return s.repo.SaveHistory(models.History{
		ID:         generateID(),
		CustomerID: session.CustomerID,
		Action:     "logout",
		Timestamp:  time.Now().Format(time.RFC3339),
	})
}

func generateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func generateID() string {
	b := make([]byte, 8)
	rand.Read(b)
	return hex.EncodeToString(b)
}

package repository

import (
	"encoding/json"
	"io/ioutil"
	"merchant-bank-api/models"
	"sync"
)

type Repository struct {
	mu sync.Mutex
}

func (r *Repository) GetCustomers() ([]models.Customer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	data, err := ioutil.ReadFile("data/customers.json")
	if err != nil {
		return nil, err
	}
	var customers []models.Customer
	json.Unmarshal(data, &customers)
	return customers, nil
}

func (r *Repository) SaveHistory(history models.History) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	data, err := ioutil.ReadFile("data/history.json")
	if err != nil {
		return err
	}
	var histories []models.History
	json.Unmarshal(data, &histories)
	histories = append(histories, history)
	updatedData, _ := json.MarshalIndent(histories, "", "  ")
	return ioutil.WriteFile("data/history.json", updatedData, 0644)
}

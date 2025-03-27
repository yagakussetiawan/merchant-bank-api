package models

type Customer struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Merchant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type History struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Action     string `json:"action"`
	Amount     int    `json:"amount,omitempty"`
	Timestamp  string `json:"timestamp"`
}

type Session struct {
	CustomerID string
	Token      string
}

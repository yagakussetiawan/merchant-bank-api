package api

import (
	"encoding/json"
	"merchant-bank-api/service"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	svc *service.Service
}

func NewAPI(svc *service.Service) *API {
	return &API{svc: svc}
}

func (a *API) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&creds)
	token, err := a.svc.Login(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (a *API) PaymentHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	var payment struct {
		Amount int `json:"amount"`
	}
	json.NewDecoder(r.Body).Decode(&payment)
	if err := a.svc.Payment(token, payment.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Write([]byte("Payment successful"))
}

func (a *API) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if err := a.svc.Logout(token); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Write([]byte("Logout successful"))
}

func (a *API) SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", a.LoginHandler).Methods("POST")
	r.HandleFunc("/payment", a.PaymentHandler).Methods("POST")
	r.HandleFunc("/logout", a.LogoutHandler).Methods("POST")
	return r
}

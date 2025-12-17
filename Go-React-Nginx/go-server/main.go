package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	HOST = "0.0.0.0"
	PORT = "8080"
)

type HttpResponse struct {
	Msg  string `json:"message"`
	Data any    `json:"data,omitempty"`
}

// User represents a simple user model
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// in-memory store
var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HttpResponse{
		Msg:  "Users list fetched",
		Data: users,
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", healthHandler)
	mux.HandleFunc("GET /api/users", usersHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", HOST, PORT),
		Handler: mux,

		WriteTimeout: 20 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	fmt.Printf("[Info] Server starting at %s:%s\n", HOST, PORT)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("[Error] Server start error\n")
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/chainpulse/backend/api-gateway/routes/handlers"
)

func main() {
	// Start HTTP server (REST or proxy)
	http.HandleFunc("/register", handlers.CreateBuyerHandler)
	// http.HandleFunc("/login", handlers.LoginBuyerHandler)

	log.Println("API Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

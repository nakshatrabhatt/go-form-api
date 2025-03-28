package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to PostgreSQL
	ConnectDB()

	// Setup Router
	r := mux.NewRouter()

	// Define Routes
	r.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API is running!"))
	}).Methods("GET")

	// Start Server
	port := "8080"
	log.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

package main

import (
	"fmt"
	"net/http"
)

func insecureAuthHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	// CWE-346: Improper Verification of Cryptographic Signature
	if username == "admin" && password == "password123" { // Insecure hardcoded check
		fmt.Fprintf(w, "Welcome, %s!", username)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func main() {
	http.HandleFunc("/login", insecureAuthHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}


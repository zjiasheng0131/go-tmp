package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// CVE-2008-0128: Predictable Session ID vulnerability
		// The session ID is generated using a predictable method (e.g., timestamp)
		sessionID := fmt.Sprintf("session-%d", r.ContentLength) // Insecure session ID generation
		
		// Store the session ID in a cookie
		cookie := &http.Cookie{
			Name:  "SESSIONID",
			Value: sessionID,
			Path:  "/",
		}
		http.SetCookie(w, cookie)

		fmt.Fprintf(w, "Session started with ID: %s", sessionID)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

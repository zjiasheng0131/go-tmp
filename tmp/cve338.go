package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateInsecureToken() string {
	rand.Seed(time.Now().UnixNano()) // Weak seeding
	token := ""
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < 16; i++ {
		token += string(chars[rand.Intn(len(chars))]) // CWE-338: Weak PRNG
	}
	return token
}

func main() {
	token := generateInsecureToken()
	fmt.Println("Generated Token:", token)
} 


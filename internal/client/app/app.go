package app

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"net"

	"github.com/hmdsefi/word-of-wisdom/config"
)

func Run(cfg *config.ClientConfig) {
	fmt.Println("Starting client...")
	conn, err := net.Dial("tcp", cfg.ServerAddress)
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())
		return
	}
	defer conn.Close()

	// Receive the challenge from the server
	challenge := make([]byte, 1024)
	n, err := conn.Read(challenge)
	if err != nil {
		fmt.Println("Error receiving challenge:", err.Error())
		return
	}

	// Compute the response using HMAC
	secretKey := []byte(cfg.SecretKey)
	hmac := hmac.New(sha256.New, secretKey)
	hmac.Write(challenge[:n])
	response := hmac.Sum(nil)

	// Send the response to the server
	conn.Write(response)

	// Receive the welcome message from the server
	welcome := make([]byte, 1024)
	n, err = conn.Read(welcome)
	if err != nil {
		fmt.Println("Error receiving welcome message:", err.Error())
		return
	}

	fmt.Println(string(welcome[:n]))
}

package app

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"github.com/hmdsefi/word-of-wisdom/internal/server/qoutes"
	"net"
	"reflect"

	"github.com/hmdsefi/word-of-wisdom/config"
)

func Run(cfg *config.ServerConfig) {
	fmt.Println("Starting server...")
	ln, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}

		go handleConnection(cfg, conn)
	}
}

func handleConnection(cfg *config.ServerConfig, conn net.Conn) {
	fmt.Println("New client connected:", conn.RemoteAddr().String())

	// Generate a challenge and a secret key
	secret := make([]byte, 10)
	rand.Read(secret)

	challenge := base32.StdEncoding.EncodeToString(secret)
	secretKey := []byte(cfg.SecretKey)

	// Send the challenge to the client
	conn.Write([]byte(challenge))

	// Receive the response from the client
	response := make([]byte, 64)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Println("Error receiving response:", err.Error())
		return
	}

	// Verify the response using HMAC
	hmac := hmac.New(sha256.New, secretKey)
	hmac.Write([]byte(challenge))
	expectedMAC := hmac.Sum(nil)

	if !reflect.DeepEqual(response[:n], expectedMAC) {
		fmt.Println("Invalid response. Authentication failed.")
		conn.Close()
		return
	}

	// Authentication successful, send a quote message to the client
	conn.Write([]byte(qoutes.RandomQuote()))
	fmt.Println("Client authenticated successfully:", conn.RemoteAddr().String())
}

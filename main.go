package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: jwt-cli <token>")
		os.Exit(1)
	}

	tokenString := os.Args[1]

	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		fmt.Println("Invalid JWT: A valid token must have exactly three parts separated by dots.")
		os.Exit(1)
	}

	headerB64 := parts[0]
	payloadB64 := parts[1]

	headerBytes, err := base64.RawURLEncoding.DecodeString(headerB64)
	if err != nil {
		fmt.Printf("Failed to decode header: %v\n", err)
		os.Exit(1)
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(payloadB64)
	if err != nil {
		fmt.Printf("Failed to decode payload: %v\n", err)
		os.Exit(1)
	}

	var headerData map[string]interface{}
	var payloadData map[string]interface{}

	if err := json.Unmarshal(headerBytes, &headerData); err != nil {
		fmt.Printf("Failed to parse header JSON: %v\n", err)
		os.Exit(1)
	}

	if err := json.Unmarshal(payloadBytes, &payloadData); err != nil {
		fmt.Printf("Failed to parse payload JSON: %v\n", err)
		os.Exit(1)
	}

	prettyHeader, _ := json.MarshalIndent(headerData, "", "  ")
	prettyPayload, _ := json.MarshalIndent(payloadData, "", "  ")

	fmt.Println("=== HEADER ===")
	fmt.Println(string(prettyHeader))
	fmt.Println("=== PAYLOAD ===")
	fmt.Println(string(prettyPayload))
	fmt.Println("=== STATUS ===")

	if expClaim, exists := payloadData["exp"]; exists {

		if expFloat, ok := expClaim.(float64); ok {

			expTime := time.Unix(int64(expFloat), 0)

			if time.Now().After(expTime) {
				fmt.Printf("TOKEN EXPIRED (expired at: %v)\n", expTime.Format(time.RFC1123))
			} else {
				fmt.Printf("TOKEN VALID (expires at: %v)\n", expTime.Format(time.RFC1123))
			}
		} else {
			fmt.Println("INVALID EXPIRY: 'exp claim is not a valid number")
		}
	} else {
		fmt.Println("NO EXPIRY: Token does not contain an 'exp' claim")
	}
}

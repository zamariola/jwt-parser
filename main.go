package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Error reading from stdin:", err)
		os.Exit(1)
	}

	tokenString := string(input)
	parser := jwt.NewParser()
	token, _, err := parser.ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		fmt.Println("Error parsing token:", err)
		os.Exit(1)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		headerJSON, err := json.MarshalIndent(token.Header, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling header to JSON:", err)
			os.Exit(1)
		}
		fmt.Println("Header:")
		fmt.Println(string(headerJSON))

		claimsJSON, err := json.MarshalIndent(claims, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling claims to JSON:", err)
			os.Exit(1)
		}
		fmt.Println("Payload:")
		fmt.Println(string(claimsJSON))
	} else {
		fmt.Println("Invalid claims format")
		os.Exit(1)
	}
}

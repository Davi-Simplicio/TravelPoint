package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Mailtrap API Request Body Struct
type MailtrapRequest struct {
	From struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"from"`
	To []struct {
		Email string `json:"email"`
	} `json:"to"`
	Subject  string `json:"subject"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

// SendCode sends a verification email using Mailtrap API
func SendCode(email string) (string, error) {
	// Generate a random 4-digit code
	rand.Seed(time.Now().UnixNano())
	code := strconv.Itoa(rand.Intn(10000))

	// Create JSON Payload Struct
	requestBody := MailtrapRequest{
		From: struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		}{
			Email: "hello@demomailtrap.com",
			Name:  "Varification Code",
		},
		To: []struct {
			Email string `json:"email"`
		}{
			{Email: email}, // Uses the provided email dynamically
		},
		Subject:  "Your Verification Code",
		Text:     "Your verification code is: " + code,
		Category: "Verification",
	}

	// Convert Struct to JSON
	payload, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return "", err
	}

	// Convert []byte payload to io.Reader using bytes.NewReader
	url := "https://send.api.mailtrap.io/api/send"
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	// Set Headers
	req.Header.Add("Authorization", "Bearer 1f0046916cfcac63ecfa533809c389a1")
	req.Header.Add("Content-Type", "application/json")

	// Make HTTP Request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "", err
	}
	defer res.Body.Close()

	// Read Response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return "", err
	}

	// Print Response
	fmt.Println("Mailtrap Response:", string(body))
	fmt.Println("Verification Code Sent:", code)

	return code, nil
}

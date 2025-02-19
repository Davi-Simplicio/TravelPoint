package utils

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// SendCode sends a verification email using Gmail SMTP through a proxy
func SendCode(toEmail string) (string, error) {
	// Generate a random 4-digit code
	rand.Seed(time.Now().UnixNano())
	code := strconv.Itoa(rand.Intn(10000))

	payload := fmt.Sprintf(`{"email": "%s", "code": "%s"}`, toEmail, code)
	resp, err := http.Post("http://127.0.0.1:5000/sendCode", "application/json", strings.NewReader(payload))
	if err != nil {
		return "", err
	}
	fmt.Println("Response Status:", resp.Status)
	
	fmt.Println("Verification Code Sent:", code)

	return code, nil
}

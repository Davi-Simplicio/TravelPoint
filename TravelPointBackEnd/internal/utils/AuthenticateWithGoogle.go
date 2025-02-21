package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config
var oauthStateString = "randomstate"
var CLIENT_SECRET = os.Getenv("CLIENT_SECRET_FILE")

// ✅ Initialize Google OAuth2 Config
func InitGoogleOAuth() {
	b, err := os.ReadFile("C:/Users/davii/TravelPointPackage/client_secret.json") // Load credentials
	if err != nil {
		fmt.Println("❌ Unable to read client secret file:", err)
		return
	}

	// Parse OAuth2 credentials
	googleOauthConfig, err = google.ConfigFromJSON(b, "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile")
	if err != nil {
		fmt.Println("❌ Unable to parse client secret file:", err)
		return
	}
    fmt.Println("✅ Google OAuth initialized")
}

// ✅ Step 1: Redirect to Google Login
func GoogleLogin(c *gin.Context) {
	if googleOauthConfig == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Google OAuth is not initialized"})
		return
	}

	// ✅ Set correct callback URL
	authURL  := googleOauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, authURL )
}

// ✅ Step 2: Handle Google OAuth Callback
func GoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OAuth state"})
		return
	}

	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	// Get user info from Google
	userInfo, err := GetGoogleUserInfo(token.AccessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	// ✅ Here, store user info in session or DB (Example: just returning the user)
	c.JSON(http.StatusOK, gin.H{
		"message":   "Google login successful!",
		"user_info": userInfo,
	})
}

// ✅ Step 3: Fetch Google User Info
func GetGoogleUserInfo(accessToken string) (map[string]interface{}, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var userInfo map[string]interface{}
	json.Unmarshal(body, &userInfo)

	return userInfo, nil
}

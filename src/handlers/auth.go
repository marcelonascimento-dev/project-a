package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"project-a/util"
)

func CreateToken(userID uint64) (string, error) {
	credentials, ok := util.GetGoogleOAuthCredentials()
	if !ok {
		return "", errors.New("google oauth secret key not found")
	}

	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        time.Now().Add(time.Minute * 15).Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(credentials.ClientSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusForbidden)
			return
		}

		credentials, ok := util.GetGoogleOAuthCredentials()
		if !ok {
			http.Error(w, "Google OAuth secret key not found", http.StatusInternalServerError)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 {
			http.Error(w, "Malformed token", http.StatusForbidden)
			return
		}

		tokenPart := parts[1]
		token, err := jwt.Parse(tokenPart, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(credentials.ClientSecret), nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		if !token.Valid {
			http.Error(w, "Token is not valid", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getGoogleOAuthConfig() (*oauth2.Config, error) {
	credentials, ok := util.GetGoogleOAuthCredentials()
	if !ok {
		return nil, errors.New("google oauth credentials not properly configured")
	}

	return &oauth2.Config{
		ClientID:     credentials.ClientID,
		ClientSecret: credentials.ClientSecret,
		RedirectURL:  credentials.RedirectURL,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}, nil
}

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	conf, err := getGoogleOAuthConfig()
	if err != nil {
		http.Error(w, "Server environment variables not properly configured", http.StatusInternalServerError)
		return
	}
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	conf, err := getGoogleOAuthConfig()
	if err != nil {
		http.Error(w, "Server environment variables not properly configured", http.StatusInternalServerError)
		return
	}

	code := r.FormValue("code")
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	client := conf.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	var userInfo map[string]interface{}
	json.Unmarshal(data, &userInfo)

	fmt.Fprintf(w, "User info: %s", userInfo)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/browser"
)

const (
	REDIRECT_URI = "http://localhost:8080/callback"
	AUTH_URL     = "https://id.twitch.tv/oauth2/authorize"
	TOKEN_URL    = "https://id.twitch.tv/oauth2/token"
)

var codeChan = make(chan string, 1)

type tokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

func authenticate() {
	browser.OpenURL(fmt.Sprintf(
		"%s?client_id=%s&redirect_uri=%s&response_type=code&scope=%s",
		AUTH_URL,
		CLIENT_ID,
		REDIRECT_URI,
		url.QueryEscape("chat:read chat:edit moderator:manage:banned_users")))

	mux := http.NewServeMux()
	mux.HandleFunc("/callback", handleCallback)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	authCode := <-codeChan
	token, err := getToken(authCode)
	if err != nil {
		log.Fatal("error exchanging auth code for token: ", err)
	} else if token.AccessToken == "" {
		log.Fatal("error getting token: empty")
	}

	ACCESS_TOKEN = token.AccessToken

	close(codeChan)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code != "" {
		_, err := w.Write([]byte("Authorization code received. You can close this tab now."))
		if err != nil {
			log.Println("Failed to write response:", err)
		}
	} else {
		if _, err := w.Write([]byte("Failed to receive authorization code.")); err != nil {
			log.Println("Failed to write response:", err)
		}
	}

	go func() { codeChan <- code }()
}

func getToken(authCode string) (*tokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", CLIENT_ID)
	data.Set("client_secret", CLIENT_SECRET)
	data.Set("code", authCode)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", REDIRECT_URI)

	resp, err := http.PostForm(TOKEN_URL, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var token tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

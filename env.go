package main

import (
	"log"
	"os"
)

func loadEnv() {
	if CHANNEL = os.Getenv("GLIB_CHANNEL"); CHANNEL == "" {
		log.Fatal("channel cannot be blank")
	}

	if USERNAME = os.Getenv("GLIB_USERNAME"); USERNAME == "" {
		log.Fatal("username cannot be blank")
	}

	if CLIENT_ID = os.Getenv("GLIB_CLIENT_ID"); CLIENT_ID == "" {
		log.Fatal("client ID cannot be blank")
	}

	if CLIENT_SECRET = os.Getenv("GLIB_CLIENT_SECRET"); CLIENT_SECRET == "" {
		log.Fatal("client secret cannot be blank")
	}
}

package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gempir/go-twitch-irc/v4"
)

var (
	CHANNEL       string
	USERNAME      string
	CLIENT_ID     string
	CLIENT_SECRET string
	ACCESS_TOKEN  string

	IRC *twitch.Client
	UI  *tea.Program
)

func init() {
	loadEnv()
}

func main() {
	UI = tea.NewProgram(model{})
	if err := UI.Start(); err != nil {
		log.Fatal("error starting program: ", err)
	}
}

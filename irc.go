package main

import (
	"fmt"
	"log"

	"github.com/gempir/go-twitch-irc/v4"
)

func listen() {
	IRC = twitch.NewClient(USERNAME, "oauth:"+ACCESS_TOKEN)

	IRC.OnPrivateMessage(func(m twitch.PrivateMessage) {
		UI.Send(fmt.Sprintf("\033[34m%s:\033[0m %s", m.User.Name, m.Message))
	})

	IRC.OnNoticeMessage(func(m twitch.NoticeMessage) {
		log.Println("\033[31merror in notice callback:\033[0m ", m.Message)
	})

	IRC.Join(CHANNEL)
	if err := IRC.Connect(); err != nil {
		log.Fatal("error connecting to channel: ", err)
	}
}

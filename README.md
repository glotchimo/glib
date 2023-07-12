# glib

A Twitch chat client in the terminal written in Go using Charm's BubbleTea and
Muesli's reflow for formatting.

## Setup

Environment variables are used to run the OAuth flow and get a token for IRC.

- `GLIB_CHANNEL`: The channel to connect to
- `GLIB_USERNAME`: Twitch username to connect with
- `GLIB_CLIENT_ID`: Twitch app client ID
- `GLIB_CLIENT_SECRET`: Twitch app client secret

Once those are set (or not if you want the defaults), just build & run it:

	go build
	./glib


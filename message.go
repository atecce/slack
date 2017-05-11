package slack

import (
	"net/http"
	"os"
)

type Message struct {
	Ts      string `json:"ts"`
	Channel string `json:"channel"`
	Text    string `json:"text"`
	User    string `json:"user"`
}

func PostMessage(channel, text string) {

	api := getMethodCall("chat.postMessage", map[string]string{
		"token":   os.Getenv("MY_KEY"),
		"channel": channel,
		"text":    text,
	})

	http.Get(api.String())
}

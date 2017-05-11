package slack

import (
	"fmt"
	"log"
	"net/url"

	"github.com/atecce/greece/uraniborg"

	"golang.org/x/net/websocket"
)

const (
	scheme = "https"
	host   = "slack.com"
)

func getMethodCall(method string, args map[string]string) url.URL {

	var api = url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   fmt.Sprintf("api/%s", method),
	}

	q := api.Query()
	for k, v := range args {
		q.Set(k, v)
	}
	api.RawQuery = q.Encode()

	return api
}

func RtmHandshake(token string) *websocket.Conn {

	api := getMethodCall("rtm.start", map[string]string{
		"token": token,
	})

	var response struct {
		Url string `json:"url"`
	}
	uraniborg.Observe(api, &response)

	ws, err := websocket.Dial(response.Url, "", "https://api.slack.com/")
	if err != nil {
		log.Fatalf("Failed to obtain websocket connection for URL %s\n", response.Url)
		log.Fatal(err)
	}

	return ws
}

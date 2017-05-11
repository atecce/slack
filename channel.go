package slack

import (
	"fmt"
	"os"

	"github.com/atecce/greece/uraniborg"
)

const count = 1000

type Channel struct {
	Id string `json:"id"`
}

type ChannelList struct {
	Channels []Channel `json:"channels"`
}

func GetChannels() []Channel {

	api := getMethodCall("channels.list", map[string]string{
		"token": os.Getenv("MY_KEY"),
	})

	var channels ChannelList
	uraniborg.Observe(api, &channels)

	return channels.Channels

}

type ChannelHistory struct {
	Latest   string    `json:"latest"`
	Messages []Message `json:"messages"`
	Has_more bool      `json:"has_more"`
	Oldest   string    `json:"oldest"`

	Channel string
}

func GetChannelHistory(channel, latest string) *ChannelHistory {

	api := getMethodCall("channels.history", map[string]string{
		"token":   os.Getenv("MY_KEY"),
		"channel": channel,
		"latest":  latest,
		"count":   fmt.Sprintf("%d", count),
	})

	var history ChannelHistory
	uraniborg.Observe(api, &history)
	history.Channel = channel

	return &history
}

func (history *ChannelHistory) Next() *ChannelHistory {

	if history.Has_more {
		return GetChannelHistory(history.Channel, history.Messages[count-1].Ts)
	}

	return nil
}

package slack

type Event struct {
	Type    string `json:"type"`
	Ts      string `json:"ts"`
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
	Error   struct {
		Msg string
	} `json:"error"`
}

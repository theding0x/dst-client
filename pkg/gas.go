package dst_client

type Gas struct {
	Network   string `json:"network"`
	Nickname  string `json:"nickname"`
	Frequency int    `json:"frequency"`
	Token     string `json:"discord_bot_token"`
}

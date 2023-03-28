package dst_client

type Board struct {
	Items     []string `json:"items"`
	Name      string   `json:"name"`
	Header    string   `json:"header"`
	Nickname  bool     `json:"nickname"`
	Color     bool     `json:"color"`
	Crypto    bool     `json:"crypto"`
	Arrows    bool     `json:"arrows"`
	Frequency int      `json:"frequency"`
	Token     string   `json:"discord_bot_token"`
}

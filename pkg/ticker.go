package dst_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

type Ticker struct {
	Ticker         string `json:"ticker"`
	Name           string `json:"name"`
	Nickname       bool   `json:"nickname"`
	Frequency      int    `json:"frequency"`
	Color          bool   `json:"color"`
	Decorator      string `json:"decorator"`
	Currency       string `json:"currency"`
	CurrencySymbol string `json:"currency_symbol"`
	Decimals       int    `json:"decimals"`
	Activity       string `json:"activity"`
	Pair           string `json:"pair"`
	PairFlip       bool   `json:"pair_flip"`
	Multiplier     int    `json:"multiplier"`
	ClientID       string `json:"client_id"`
	Crypto         bool   `json:"crypto"`
	Token          string `json:"discord_bot_token"`
	TwelveDataKey  string `json:"twelve_data_key"`
	Exrate         int    `json:"exrate"`
}
type TickerList struct {
	Count   int                 `json:"count"`
	Tickers []map[string]Ticker `json:"tickers"`
}

func (c *Client) AddTicker(botJson string) error {
	// curl -X POST -H "Content-type: application/json" -d @bots/crypto/btc.json localhost:8080/ticker
	var ticker Ticker
	err := json.Unmarshal([]byte(botJson), &ticker)

	if err != nil {
		log.Fatal().Err(err).Msg("Error unmarshalling ticker data")
	}
	log.Info().
		Str("name", ticker.Name).
		Bool("crypto", ticker.Crypto).
		Str("ticker", ticker.Ticker).
		Str("base_url", BaseURL).
		Msg("Adding ticker from JSON")

	req, err := http.NewRequest("POST", c.BaseURL+"/ticker", bytes.NewBuffer([]byte(botJson)))
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating request")
	}
	_, err = c.HTTPClient.Do(req)
	return err
}
func (c *Client) NewTicker(name string, crypto bool, ticker string, color bool, decorator string, currency string, currencySymbol string, pair string, pairFlip bool, activity string, decimals int, nickname bool, frequency int, token string, deploy bool) error {
	tickerObj := Ticker{
		Ticker:         ticker,
		Name:           name,
		Nickname:       nickname,
		Frequency:      frequency,
		Color:          color,
		Crypto:         crypto,
		Decorator:      decorator,
		Currency:       currency,
		CurrencySymbol: currencySymbol,
		Decimals:       decimals,
		Activity:       activity,
		Pair:           pair,
		PairFlip:       pairFlip,
		Token:          token,
	}
	tickerJson, err := json.Marshal(tickerObj)
	if err != nil {
		log.Fatal().Err(err).Msg("Error marshalling ticker data")
	}
	err = os.WriteFile(fmt.Sprintf("bots/tickers/%s.json", tickerObj.Ticker), tickerJson, 0644)
	if deploy {
		err = c.AddTicker(string(tickerJson))
	}

	return err
}

func (c *Client) GetTicker(ticker string) (Ticker, error) {
	var tickerObj Ticker
	req, err := http.NewRequest("GET", c.BaseURL+"/ticker/"+ticker, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating request")
	}
	err = c.SendRequest(req, &tickerObj)
	return tickerObj, err
}

func (c *Client) GetTickers() (map[string]Ticker, error) {
	response, err := http.Get(BaseURL + "/ticker")
	if err != nil {
		return make(map[string]Ticker), fmt.Errorf("The HTTP request failed with error %s\n", err)
	}
	defer response.Body.Close()

	var tickerList map[string]Ticker

	err = json.NewDecoder(response.Body).Decode(&tickerList)
	if err != nil {
		return make(map[string]Ticker), fmt.Errorf("Error decoding JSON response: %s\n", err)
	}

	return tickerList, nil
}

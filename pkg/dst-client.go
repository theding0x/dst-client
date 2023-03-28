package dst_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	BaseURL = "http://localhost:8080"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}
type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    BaseURL,
	}
}

func (c *Client) TestConnection() error {
	res, err := c.HTTPClient.Get(c.BaseURL + "/ticker")
	if err != nil {
		log.Fatal().Err(err).Msg("Error testing connection to DST service")
	}
	log.Info().Str("status", res.Status).Msg("Test connection")
	return err
}

func (c *Client) SendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal().Err(err).Msg("Error sending request")
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return errors.New(errRes.Message)
		}
		return fmt.Errorf("error %d: %s", errRes.Code, errRes.Message)
	}
	fullResponse := successResponse{
		Data: v,
	}
	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}
	return nil
}

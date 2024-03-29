package discordclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/context/ctxhttp"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SendMessage(ctx context.Context, webhookURL string, message Message) error {
	var buf bytes.Buffer
	// err := json.NewEncoder(&buf).Encode(Test{X: 1})
	err := json.NewEncoder(&buf).Encode(message)
	if err != nil {
		return err
	}

	resp, err := ctxhttp.Post(ctx, http.DefaultClient, webhookURL, "application/json", &buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		defer resp.Body.Close()

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf(string(responseBody))
	}

	return nil
}

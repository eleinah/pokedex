package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	HTTP http.Client
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		HTTP: http.Client{
			Timeout: timeout,
		},
	}
}

package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	HTTP http.Client
}

func NewClient(timeout time.Duration) *Client {
	if timeout < 1 {
		return &Client{
			HTTP: http.Client{
				Timeout: 0,
			},
		}
	}
	return &Client{
		HTTP: http.Client{
			Timeout: timeout * time.Second,
		},
	}
}

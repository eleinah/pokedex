package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	http.Client
}

func NewClient(timeout time.Duration) *Client {
	if timeout < 1 {
		return &Client{
			http.Client{
				Timeout: 0,
			},
		}
	}
	return &Client{
		http.Client{
			Timeout: timeout * time.Second,
		},
	}
}

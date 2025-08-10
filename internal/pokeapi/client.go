package pokeapi

import (
	"net/http"
	"time"
)

func NewClient(timeout time.Duration) *http.Client {
	if timeout < 1 {
		return &http.Client{
			Timeout: 0,
		}
	}
	return &http.Client{
		Timeout: timeout * time.Second,
	}
}

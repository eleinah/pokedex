package pokeapi

import (
	"net/http"
	"time"

	"github.com/eleinah/pokedex/internal/pokecache"
)

type Client struct {
	Cache pokecache.Cache
	HTTP  http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		Cache: pokecache.NewCache(cacheInterval),
		HTTP: http.Client{
			Timeout: timeout,
		},
	}
}

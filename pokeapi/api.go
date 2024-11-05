package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/wolv89/gokedex/pokecache"
)

const API_BASE_URL = "https://pokeapi.co/api/v2"

type APIQuery interface{}

type API struct {
	cache pokecache.Cache
}

var pkapi = &API{
	pokecache.NewCache(5 * time.Second),
}

func (api *API) Call(url string, query APIQuery) error {

	if len(url) == 0 {
		return nil
	}

	var (
		body []byte
		err  error
		ok   bool
	)

	body, ok = api.cache.Get(url)

	if !ok {
		res, err := http.Get(url)

		if err != nil {
			return err
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
	}

	err = json.Unmarshal(body, &query)
	if err != nil {
		return err
	}

	api.cache.Add(url, body)

	return nil

}

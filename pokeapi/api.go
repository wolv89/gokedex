package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const API_BASE_URL = "https://pokeapi.co/api/v2"

type APIQuery interface{}

func Call(url string, query APIQuery) error {

	if len(url) == 0 {
		return nil
	}

	res, err := http.Get(url)

	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &query)
	if err != nil {
		return err
	}

	return nil

}

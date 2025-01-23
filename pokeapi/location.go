package pokeapi

import "fmt"

// Generated from https://mholt.github.io/json-to-go/ <3
type LocationAreaQuery struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				ConditionValues []any `json:"condition_values"`
				Chance          int   `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetPokemon(location string) ([]Pokemon, error) {

	if len(location) == 0 {
		return nil, fmt.Errorf("No location entered")
	}

	url := API_BASE_URL + "/location-area/" + location
	var query LocationAreaQuery

	err := pkapi.Call(url, &query)
	if err != nil {
		return nil, err
	}

	if len(query.PokemonEncounters) == 0 {
		return nil, fmt.Errorf("No Pokemon found here?!")
	}

	foundPokemon := make([]Pokemon, 0)

	for _, encounter := range query.PokemonEncounters {
		foundPokemon = append(foundPokemon, Pokemon{
			encounter.Pokemon.Name,
			encounter.Pokemon.URL,
		})
	}

	return foundPokemon, nil

}

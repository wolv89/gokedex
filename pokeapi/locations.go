package pokeapi

type LocationQuery struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocations(url string) ([]Location, string, string, error) {

	var query LocationQuery

	if len(url) == 0 {
		url = API_BASE_URL + "/location/"
	}

	err := pkapi.Call(url, &query)
	if err != nil {
		return nil, "", "", err
	}

	var next, prev string
	if query.Next != nil {
		next = *query.Next
	}
	if query.Previous != nil {
		prev = *query.Previous
	}

	return query.Results, next, prev, nil

}

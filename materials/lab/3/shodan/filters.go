package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *Client) Filters() ([]string, error) {
	res, err := http.Get(fmt.Sprintf("%s/shodan/host/search/filters?key=%s", BaseURL, s.apiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var arr []string

	if err := json.NewDecoder(res.Body).Decode(&arr); err != nil {
		return nil, err
	}
	return arr, nil
}

func (s *Client) Facets() ([]string, error) {
	res, err := http.Get(fmt.Sprintf("%s/shodan/host/search/facets?key=%s", BaseURL, s.apiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var arr []string

	if err := json.NewDecoder(res.Body).Decode(&arr); err != nil {
		return nil, err
	}
	return arr, nil
}

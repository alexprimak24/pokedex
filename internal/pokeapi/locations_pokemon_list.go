package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocPokemons(area []string) (RespLocation, error) {
	url := baseURL + "/location-area/" + area[0]

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespLocation{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespLocation{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}

	locationsResp := RespLocation{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespLocation{}, err
	}

	return locationsResp, nil
}
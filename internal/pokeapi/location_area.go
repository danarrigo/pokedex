package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"

)

type LocationResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationResponse struct {
	Count    int              `json:"count"`
	Next     *string          `json:"next"`
	Previous *string          `json:"previous"`
	Results  []LocationResult `json:"results"`
}

func (c *Client) FetchLocations(pageURL *string) (LocationResponse, error) {
	
	url := "https://pokeapi.co/api/v2/location-area/"
	if pageURL != nil {
		url = *pageURL
	}
	var response LocationResponse

	if val,ok:=c.cache.Get(url);ok{
		if err:=json.Unmarshal(val,&response); err!=nil{
			return LocationResponse{},err
		}
	}else{
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return LocationResponse{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationResponse{}, err
		}
		defer resp.Body.Close()
	
		data, err := io.ReadAll(resp.Body)
			if err != nil {
			    return LocationResponse{}, err
			}
			c.cache.Add(url, data)
			if err := json.Unmarshal(data, &response); err != nil {
			
			    return LocationResponse{}, err
			
			}
	}

	return response, nil
}

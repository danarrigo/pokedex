package pokeapi



import (

	"encoding/json"

	"net/http"

)



type LocationResponse struct {

	Count    int     `json:"count"`

	Next     *string `json:"next"`

	Previous *string `json:"previous"`

	Results  []struct {

		Name string `json:"name"`

		URL  string `json:"url"`

	} `json:"results"`

}



func (c *Client) FetchLocations(pageURL *string) (LocationResponse, error) {

	url := "https://pokeapi.co/api/v2/location-area/"

	if pageURL != nil {

		url = *pageURL

	}



	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {

		return LocationResponse{}, err

	}



	resp, err := c.httpClient.Do(req)

	if err != nil {

		return LocationResponse{}, err

	}

	defer resp.Body.Close()



	var response LocationResponse

	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&response); err != nil {

		return LocationResponse{}, err

	}



	return response, nil

}

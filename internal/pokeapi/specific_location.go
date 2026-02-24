package pokeapi

import("errors"
		"net/http"
		"io"
		"encoding/json"
		"fmt")

type SpecificLocationData struct {
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



func (c *Client) FetchSpecificLocationInfo(locationName string)(SpecificLocationData,error){
	var data SpecificLocationData
	baseUrl:="https://pokeapi.co/api/v2/location-area/"
	if locationName==""{
		return SpecificLocationData{},errors.New("Error: must pass a valid locationName")
	}
	fullUrl:=baseUrl+locationName+"/"

	if val, ok := c.cache.Get(fullUrl); ok {
		// Cache hit - unmarshal from cached value
		if err := json.Unmarshal(val, &data); err != nil {
			return SpecificLocationData{}, err
		}
	}else{
		req,err:=http.NewRequest(http.MethodGet,fullUrl,nil)
		if err!=nil{
			return SpecificLocationData{},err
		}
		res,err:=c.httpClient.Do(req)
		if err!=nil{
			return SpecificLocationData{},err
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
				return SpecificLocationData{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
		}
		body,err:=io.ReadAll(res.Body)
		if err!=nil{
			return SpecificLocationData{},err
		}
		c.cache.Add(fullUrl, body)
		if err=json.Unmarshal(body,&data);err!=nil{
			return SpecificLocationData{},err
		}
	}
	return data	,nil
}

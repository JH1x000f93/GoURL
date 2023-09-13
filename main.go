package GoURL

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ShortRequest struct {
	URL    string `json:"url"`
	APIKey string `json:"apikey"`
}

type ShortResponse struct {
	Status    string `json:"status"`
	Permalink string `json:"permalink,omitempty"`
	Short     string `json:"short,omitempty"`
	Info      string `json:"info,omitempty"`
}

// Short("https://example.com/long", "678GHJghj78IG0osh7GJi7hmK")
func Short(url, api_key string) (ShortResponse, error) {
	apiUrl := "https://n9.cl/api/short"

	requestData := ShortRequest{
		URL:    url,
		APIKey: api_key,
	}

	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return ShortResponse{}, err
	}

	response, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return ShortResponse{}, err
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		var short ShortResponse
		if err := json.NewDecoder(response.Body).Decode(&short); err != nil {
			return ShortResponse{}, err
		}
		return short, nil
	} else {
		return ShortResponse{}, fmt.Errorf("Status code %d", response.StatusCode)
	}
}

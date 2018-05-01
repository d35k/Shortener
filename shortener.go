package shortener

import (
	"encoding/json"
	"github.com/parnurzeal/gorequest"
)

//GoogleAPIKey For GOOGLE API KEY
var GoogleAPIKey string

//GoogleShortResponse for Google Response
type GoogleShortResponse struct {
	Kind    string `json:"kind"`
	ID      string `json:"id"`
	LongURL string `json:"longUrl"`
}

//SetAPIKey for Setting API Key
func SetAPIKey(apiKey string) error {
	GoogleAPIKey = apiKey
	return nil
}

// Short function for URL
func Short(url string) string {
	var GoogleShortResponse *GoogleShortResponse
	googleShortenRequest := gorequest.New()
	if GoogleAPIKey == "" {
		return "You have to set Google API Key"
	}
	if url == "" {
		return "You have to set URL for Shortener"
	}
	googleShortenURL := "https://www.googleapis.com/urlshortener/v1/url?key=" + GoogleAPIKey

	googleRequestResponse, _, googleResponseError := googleShortenRequest.Post(googleShortenURL).
		Set("Accept", "application/json").
		Set("Content-Type", "application/json").
		Send(`{"longUrl":"` + url + `"}`).End()

	if googleResponseError != nil {
		return "Google Response Error"
	}

	if googleRequestResponse.Status == "200 OK" {

		if err := json.NewDecoder(googleRequestResponse.Body).Decode(&GoogleShortResponse); err != nil {
			return "Decore Error!"
		}

		return GoogleShortResponse.ID
	}

	return "Some error occurred, please try again later"
}

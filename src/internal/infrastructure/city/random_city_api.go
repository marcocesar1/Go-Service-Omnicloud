package city

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RandomCityResponse struct {
	City string `json:"city"`
}

type RandomCityApi struct {
	url string
}

func NewRandomCityApi() *RandomCityApi {
	return &RandomCityApi{
		url: "https://random-city-api.vercel.app/api/random-city",
	}
}

func (r *RandomCityApi) GetCityName() (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, r.url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error status: %d", resp.StatusCode)
	}

	var data RandomCityResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return data.City, nil
}

package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetCotyCountry(latitude, longitude float64, apiKey string) (string, string, error) {
	url := fmt.Sprintf("https://api.geoapify.com/v1/geocode/reverse?lat=%v&lon=%v&apiKey=%v", latitude, longitude, apiKey)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return "", "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}
	var data RevGeoType
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", "", err
	}

	return data.Features[0].Properties.City, strings.ToUpper(data.Features[0].Properties.CountryCode), nil
}

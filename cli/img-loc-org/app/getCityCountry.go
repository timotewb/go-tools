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

	var cityName string
	var countryName string

	if data.Features[0].Properties.City == "" {
		cityName = "Unknown"
	} else {
		cityName = data.Features[0].Properties.City
	}

	if strings.ToUpper(data.Features[0].Properties.CountryCode) == "" {
		countryName = "UNK"
	} else {
		countryName = strings.ToUpper(data.Features[0].Properties.CountryCode)
	}

	return cityName, countryName, nil
}

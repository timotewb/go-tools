package imglocorg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/timotewb/go-tools/cmd/img-loc-org/app"
	"github.com/timotewb/go-tools/cmd/shared"
)

func App(help bool, apiKey, inDir, outDir string) error {

	//loop through image files in folder
	// Read the directory
	entries, err := os.ReadDir(inDir)
	if err != nil {
		return err
	}

	// Loop through each entry
	for _, entry := range entries {
		// Check if the entry is a file
		if !entry.IsDir() {
			if shared.IsImageFile(filepath.Join(inDir, entry.Name())) {

				// get lat lon from image
				latitude, longitude, err := app.GetLatLon(filepath.Join(inDir, entry.Name()))
				if err != nil {
					return err
				} else if latitude == 0 && longitude == 0 && err == nil {
					fmt.Println("No Lat Lon data found in image file.")
				} else {

					fmt.Printf("Latitude: %v, Longitude: %v\n", latitude, longitude)

					// get city name and country code from coordinates
					cityName, countryCode, err := app.GetCotyCountry(latitude, longitude, apiKey)
					if err != nil {
						return err
					} else {
						fmt.Println(cityName, countryCode)
					}
				}
			}
		}
	}

	return nil
}

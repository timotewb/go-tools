package imglocorg

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/timotewb/go-tools/cli/img-loc-org/app"
	"github.com/timotewb/go-tools/cli/shared"
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
		fmt.Println(entry.Name())
		// Check if the entry is a file
		if entry.IsDir() {
			err = App(help, apiKey, filepath.Join(inDir, entry.Name()), outDir)
			if err != nil {
				return err
			}
		} else {
			if shared.IsImageFile(filepath.Join(inDir, entry.Name())) {

				// get lat lon from image
				latitude, longitude, err := app.GetLatLon(filepath.Join(inDir, entry.Name()))
				if err != nil {
					return err
				} else if latitude == 0 && longitude == 0 && err == nil {
					fmt.Println("   No Lat Lon data found in image file ")
				} else {

					time.Sleep(1 * time.Second)

					// get city name and country code from coordinates
					cityName, countryCode, err := app.GetCotyCountry(latitude, longitude, apiKey)
					if err != nil {
						return err
					} else {
						newPath := filepath.Join(outDir, cityName+", "+countryCode)
						// Use os.MkdirAll to create the directory and any necessary parents
						err := os.MkdirAll(newPath, 0755)
						if err != nil {
							return err
						}
						// Use os.Rename to move the file
						err = os.Rename(filepath.Join(inDir, entry.Name()), filepath.Join(newPath, entry.Name()))
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}

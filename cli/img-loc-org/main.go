package imglocorg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/timotewb/go-tools/cli/img-loc-org/app"
	"github.com/timotewb/go-tools/cli/shared"
)

func Main(help bool, apiKey, inDir, outDir string) error {

	//loop through image files in folder
	// Read the directory
	entries, err := os.ReadDir(inDir)
	if err != nil {
		return fmt.Errorf("%w: Main - os.ReadDir()", err)
	}

	// Loop through each entry
	for _, entry := range entries {
		fmt.Println(entry.Name())
		// Check if the entry is a file
		if entry.IsDir() {
			err = Main(help, apiKey, filepath.Join(inDir, entry.Name()), outDir)
			if err != nil {
				return fmt.Errorf("%w: Main - Main()", err)
			}
		} else {
			if shared.IsImageFile(filepath.Join(inDir, entry.Name()), 2) && !strings.HasPrefix(entry.Name(), "._") {

				// get lat lon from image
				latitude, longitude, err := app.GetLatLon(filepath.Join(inDir, entry.Name()))
				if err != nil {
					fmt.Println(fmt.Errorf("%w: Main - app.GetLatLon()", err))
				} else if latitude == 0 && longitude == 0 && err == nil {
					fmt.Println("   No Lat Lon data found in image file ")
				} else {

					time.Sleep(1 * time.Second)

					// get city name and country code from coordinates
					cityName, countryCode, err := app.GetCotyCountry(latitude, longitude, apiKey)
					if err != nil {
						return fmt.Errorf("%w: Main - app.GetCotyCountry()", err)
					} else {
						newPath := filepath.Join(outDir, cityName+", "+countryCode)

						// Check if file already exists with same anme in location
						if shared.FileExists(filepath.Join(newPath, entry.Name())) {
							newPath = filepath.Join(outDir, "Duplicates", cityName+", "+countryCode)
						}

						// Use os.MkdirAll to create the directory and any necessary parents
						err := os.MkdirAll(newPath, 0755)
						if err != nil {
							return fmt.Errorf("%w: Main - os.MkdirAll()", err)
						}

						// Move file
						err = os.Rename(filepath.Join(inDir, entry.Name()), filepath.Join(newPath, entry.Name()))
						if err != nil {
							return fmt.Errorf("%w: Main - os.Rename()", err)
						}
					}
				}
			}
		}
	}
	return nil
}

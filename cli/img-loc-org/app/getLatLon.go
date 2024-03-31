package app

import (
	"fmt"
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

func GetLatLon(imageFile string) (float64, float64, error) {
	// Open the image file
	file, err := os.Open(imageFile)
	if err != nil {
		return 0, 0, fmt.Errorf("%w: GetLatLon - os.Open()", err)
	}
	defer file.Close()

	// Parse the EXIF data
	exifData, err := exif.Decode(file)
	if err != nil {
		return 0, 0, fmt.Errorf("%w: GetLatLon - exif.Decode()", err)
	}

	// Get the latitude and longitude
	latitude, longitude, err := exifData.LatLong()
	if err != nil {
		return 0, 0, fmt.Errorf("%w: GetLatLon - exifData.LatLong()", err)
	}

	return latitude, longitude, err
}

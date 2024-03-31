package shared

import (
	"os"
	"path/filepath"
	"strings"
)

// contains checks if a given string is present in a slice of strings.
//
// Parameters:
//   - list: A slice of strings to search through.
//   - target: The string to search for in the list.
//
// Returns:
//   - bool: True if the target string is found in the list, false otherwise.
func Contains(list []string, target string) bool {
	for _, str := range list {
		if str == target {
			return true
		}
	}
	return false
}

// isImageFile checks if the given file extension is an image file extension
func IsImageFile(file string) bool {
	// Extract the file extension from the file name or path
	ext := filepath.Ext(file)

	// Convert the extension to lowercase to ensure case-insensitive comparison
	ext = strings.ToLower(ext)

	// Check if the extension is in the list of image extensions
	for _, imageExt := range ImgExt {
		if ext == imageExt {
			return true
		}
	}

	return false
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

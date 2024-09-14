package fileutils

import (
	"encoding/json"
	"os"
)

// LoadJSON reads a JSON file and unmarshals its content into the provided data structure.
func LoadJSON(filename string, v interface{}) error {
	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Unmarshal the JSON data into the provided data structure
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}

	return nil
}

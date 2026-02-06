package load

import (
	"encoding/json"
	"os"
	"strings"
)

// ReadLines reads all lines from the provided file path.
func ReadLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimRight(string(data), "\n"), "\n"), nil
}

// ReadAll reads the entire content of the provided file path as a single string.
func ReadAll(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(data), "\n"), nil
}

// Json reads the content of the provided file path and unmarshals it into the provided variable.
func Json(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

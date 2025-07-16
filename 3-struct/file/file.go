package file

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type FileReader interface {
	Read() ([]byte, error)
}

type JSONFileReader struct {
	filename string
}

func NewJSONFileReader(filename string) *JSONFileReader {
	return &JSONFileReader{
		filename: filename,
	}
}

func (f *JSONFileReader) Read() ([]byte, error) {
	if _, err := os.Stat(f.filename); os.IsNotExist(err) {
		return nil, err
	}

	if !f.isJsonFile() {
		return nil, fmt.Errorf("file must have .json extension")
	}

	data, err := os.ReadFile(f.filename)
	if err != nil {
		return nil, err
	}

	var jsonData any
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, fmt.Errorf("invalid JSON: %v", err)
	}

	return data, nil
}

func (f *JSONFileReader) isJsonFile() bool {
	return strings.HasSuffix(f.filename, ".json")
}

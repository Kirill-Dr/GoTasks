package file

import (
	"os"
	"strings"
)

type JsonFile struct {
	filename string
}

func NewJsonFile(filename string) *JsonFile {
	return &JsonFile{
		filename: filename,
	}
}

func (file *JsonFile) ReadFile() ([]byte, error) {
	data, err := os.ReadFile(file.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (file *JsonFile) IsJsonFile() bool {
	return strings.HasSuffix(file.filename, ".json")
}

func (file *JsonFile) GetFilename() string {
	return file.filename
}

package file

import (
	"fmt"
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

func (file *JsonFile) WriteFile(content []byte) {
	data, err := os.Create(file.filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer data.Close()
	_, err = data.Write(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully")
}

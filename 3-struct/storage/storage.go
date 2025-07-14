package storage

import (
	"encoding/json"
	"errors"
	"os"

	"3-struct/bins"
	"3-struct/file"
)

type Storage struct {
	jsonFile *file.JsonFile
}

func NewStorage(filename string) *Storage {
	return &Storage{
		jsonFile: file.NewJsonFile(filename),
	}
}

func (storage *Storage) SaveBins(binList *bins.BinList) error {
	if !storage.jsonFile.IsJsonFile() {
		return errors.New("file is not a JSON file")
	}

	data, err := json.Marshal(binList)
	if err != nil {
		return errors.New("error marshalling data")
	}

	err = os.WriteFile(storage.jsonFile.GetFilename(), data, 0644)
	if err != nil {
		return errors.New("error writing file")
	}

	return nil
}

func (storage *Storage) ReadBins() (*bins.BinList, error) {
	if !storage.jsonFile.IsJsonFile() {
		return nil, errors.New("file is not a JSON file")
	}

	data, err := storage.jsonFile.ReadFile()
	if err != nil {
		return nil, errors.New("error reading file")
	}

	var binList bins.BinList
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return nil, errors.New("error serializing JSON")
	}

	return &binList, nil
}

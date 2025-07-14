package storage

import (
	"encoding/json"
	"errors"

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

	storage.jsonFile.WriteFile(data)
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
		return nil, errors.New("error decoding JSON")
	}

	return &binList, nil
}

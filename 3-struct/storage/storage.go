package storage

import (
	"3-struct/bins"
	"encoding/json"
	"os"
	"time"
)

type Storage interface {
	SaveBins(binList *bins.BinList) error
	ReadBins() (*bins.BinList, error)
	GetFilename() string
}

type FileStorage struct {
	filename string
}

func NewStorage(filename string) *FileStorage {
	return &FileStorage{
		filename: filename,
	}
}

func (s *FileStorage) SaveBins(binList *bins.BinList) error {
	binList.UpdatedAt = time.Now()
	data, err := json.Marshal(binList)
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, data, 0644)
}

func (s *FileStorage) ReadBins() (*bins.BinList, error) {
	data, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, err
	}

	var binList bins.BinList
	json.Unmarshal(data, &binList)
	return &binList, nil
}

func (s *FileStorage) GetFilename() string {
	return s.filename
}

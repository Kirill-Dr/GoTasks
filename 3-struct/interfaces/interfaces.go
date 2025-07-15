package interfaces

import "3-struct/bins"

type FileService interface {
	ReadFile() ([]byte, error)
	IsJsonFile() bool
	GetFilename() string
}

type API interface {
}

type StorageService interface {
	SaveBins(binList *bins.BinList) error
	ReadBins() (*bins.BinList, error)
}

type BinService interface {
	NewBin(id string, private bool, name string) bins.Bin
	NewBinList() bins.BinList
	AddBinToList(binList *bins.BinList, bin bins.Bin)
	GenerateBinID() string
	GetBinData() (bool, string)
}

type ApplicationService interface {
	CreateBin() error
	FindBin(name string) error
	DeleteBin(name string) error
	ListAllBins() error
	Run() error
}

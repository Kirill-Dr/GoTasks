package services

import (
	"3-struct/bins"
	"3-struct/config"
	"3-struct/interfaces"
	"3-struct/storage"
	"fmt"
)

type StorageServiceImpl struct {
	storage *storage.Storage
}

func NewStorageService(filename string) interfaces.StorageService {
	return &StorageServiceImpl{
		storage: storage.NewStorage(filename),
	}
}

func (s *StorageServiceImpl) SaveBins(binList *bins.BinList) error {
	return s.storage.SaveBins(binList)
}

func (s *StorageServiceImpl) ReadBins() (*bins.BinList, error) {
	return s.storage.ReadBins()
}

type BinServiceImpl struct{}

func NewBinService() interfaces.BinService {
	return &BinServiceImpl{}
}

func (b *BinServiceImpl) NewBin(id string, private bool, name string) bins.Bin {
	return bins.NewBin(id, private, name)
}

func (b *BinServiceImpl) NewBinList() bins.BinList {
	return bins.NewBinList()
}

func (b *BinServiceImpl) AddBinToList(binList *bins.BinList, bin bins.Bin) {
	bins.AddBinToList(binList, bin)
}

func (b *BinServiceImpl) GenerateBinID() string {
	return bins.GenerateBinID()
}

func (b *BinServiceImpl) GetBinData() (bool, string) {
	return bins.GetBinData()
}

type ApplicationServiceImpl struct {
	storageService interfaces.StorageService
	binService     interfaces.BinService
	config         *config.Config
	binList        *bins.BinList
}

func NewApplicationService(storageService interfaces.StorageService, binService interfaces.BinService, cfg *config.Config) interfaces.ApplicationService {
	binList, err := storageService.ReadBins()
	if err != nil {
		newList := binService.NewBinList()
		binList = &newList
	}

	return &ApplicationServiceImpl{
		storageService: storageService,
		binService:     binService,
		config:         cfg,
		binList:        binList,
	}
}

func (app *ApplicationServiceImpl) CreateBin() error {
	private, name := app.binService.GetBinData()
	bin := app.binService.NewBin(app.binService.GenerateBinID(), private, name)
	app.binService.AddBinToList(app.binList, bin)

	err := app.storageService.SaveBins(app.binList)
	if err != nil {
		return fmt.Errorf("error saving: %v", err)
	}
	fmt.Println("Bin created successfully")
	return nil
}

func (app *ApplicationServiceImpl) FindBin(name string) error {
	foundBins := app.findBinsByName(name)
	if len(foundBins) == 0 {
		fmt.Println("No bins found")
		return nil
	}
	for index, bin := range foundBins {
		fmt.Printf("Bin #%d\n", index+1)
		fmt.Printf("ID: %s\n", bin.Id)
		fmt.Printf("Name: %s\n", bin.Name)
		fmt.Printf("Private: %t\n", bin.Private)
	}
	return nil
}

func (app *ApplicationServiceImpl) DeleteBin(name string) error {
	deleted := app.deleteBinByName(name)
	if deleted {
		err := app.storageService.SaveBins(app.binList)
		if err != nil {
			return fmt.Errorf("error saving: %v", err)
		}
		fmt.Println("Bin deleted successfully")
	} else {
		fmt.Println("Bin not found")
	}
	return nil
}

func (app *ApplicationServiceImpl) ListAllBins() error {
	if len(app.binList.Bins) == 0 {
		fmt.Println("No bins found")
		return nil
	}
	fmt.Printf("Total bins: %d\n", len(app.binList.Bins))
	for i, bin := range app.binList.Bins {
		fmt.Printf("%d. ID: %s, Name: %s, Private: %t\n",
			i+1, bin.Id, bin.Name, bin.Private)
	}
	return nil
}

func (app *ApplicationServiceImpl) Run() error {
	fmt.Println("--- Bin Manager ---")

	for {
		choice := promptData([]string{"1. Create bin", "2. Find bin", "3. Delete bin", "4. List all bins", "5. Exit", "Choose an option"})
		switch choice {
		case "1":
			if err := app.CreateBin(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		case "2":
			name := promptData([]string{"Enter name to find"})
			if err := app.FindBin(name); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		case "3":
			name := promptData([]string{"Enter name to delete"})
			if err := app.DeleteBin(name); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		case "4":
			if err := app.ListAllBins(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		default:
			return nil
		}
	}
}

func (app *ApplicationServiceImpl) findBinsByName(name string) []bins.Bin {
	var found []bins.Bin
	for _, bin := range app.binList.Bins {
		if bin.Name == name {
			found = append(found, bin)
		}
	}
	return found
}

func (app *ApplicationServiceImpl) deleteBinByName(name string) bool {
	for i, bin := range app.binList.Bins {
		if bin.Name == name {
			app.binList.Bins = append(app.binList.Bins[:i], app.binList.Bins[i+1:]...)
			return true
		}
	}
	return false
}

func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}

package main

import (
	"3-struct/bins"
	"3-struct/storage"
	"fmt"
)

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

type Application struct {
	storageService StorageService
	binService     BinService
	binList        *bins.BinList
}

func NewApplication(storageService StorageService, binService BinService) *Application {
	binList, err := storageService.ReadBins()
	if err != nil {
		newList := binService.NewBinList()
		binList = &newList
	}

	return &Application{
		storageService: storageService,
		binService:     binService,
		binList:        binList,
	}
}

func main() {
	fmt.Println("--- Bin Manager ---")

	storageService := storage.NewStorage("bins.json")
	binService := &BinServiceImpl{}
	app := NewApplication(storageService, binService)

Menu:
	for {
		choice := promptData([]string{"1. Create bin", "2. Find bin", "3. Delete bin", "4. List all bins", "5. Exit", "Choose an option"})
		switch choice {
		case "1":
			createBin(app)
		case "2":
			findBin(app)
		case "3":
			deleteBin(app)
		case "4":
			listAllBins(app)
		default:
			break Menu
		}
	}
}

func createBin(app *Application) {
	private, name := app.binService.GetBinData()
	bin := app.binService.NewBin(app.binService.GenerateBinID(), private, name)
	app.binService.AddBinToList(app.binList, bin)

	err := app.storageService.SaveBins(app.binList)
	if err != nil {
		fmt.Printf("Error saving: %v\n", err)
	} else {
		fmt.Println("Bin created successfully")
	}
}

func findBin(app *Application) {
	name := promptData([]string{"Enter name to find"})
	foundBins := findBinsByName(app.binList, name)
	if len(foundBins) == 0 {
		fmt.Println("No bins found")
		return
	}
	for index, bin := range foundBins {
		fmt.Printf("Bin #%d\n", index+1)
		fmt.Printf("ID: %s\n", bin.Id)
		fmt.Printf("Name: %s\n", bin.Name)
		fmt.Printf("Private: %t\n", bin.Private)
	}
}

func deleteBin(app *Application) {
	name := promptData([]string{"Enter name to delete"})
	deleted := deleteBinByName(app.binList, name)
	if deleted {
		err := app.storageService.SaveBins(app.binList)
		if err != nil {
			fmt.Printf("Error saving: %v\n", err)
		} else {
			fmt.Println("Bin deleted successfully")
		}
	} else {
		fmt.Println("Bin not found")
	}
}

func listAllBins(app *Application) {
	if len(app.binList.Bins) == 0 {
		fmt.Println("No bins found")
		return
	}
	fmt.Printf("Total bins: %d\n", len(app.binList.Bins))
	for i, bin := range app.binList.Bins {
		fmt.Printf("%d. ID: %s, Name: %s, Private: %t\n",
			i+1, bin.Id, bin.Name, bin.Private)
	}
}

func findBinsByName(binList *bins.BinList, name string) []bins.Bin {
	var found []bins.Bin
	for _, bin := range binList.Bins {
		if bin.Name == name {
			found = append(found, bin)
		}
	}
	return found
}

func deleteBinByName(binList *bins.BinList, name string) bool {
	for i, bin := range binList.Bins {
		if bin.Name == name {
			binList.Bins = append(binList.Bins[:i], binList.Bins[i+1:]...)
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

type BinServiceImpl struct{}

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

package main

import (
	"3-struct/api"
	"3-struct/config"
	"3-struct/file"
	"3-struct/storage"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("--- JSONBIN ---")

	create := flag.Bool("create", false, "create new bin")
	update := flag.Bool("update", false, "update existing bin")
	get := flag.Bool("get", false, "get bin by id")
	delete := flag.Bool("delete", false, "delete bin by id")
	list := flag.Bool("list", false, "list all bins")

	filename := flag.String("file", "", "filename")
	binName := flag.String("name", "", "bin name")
	binId := flag.String("id", "", "bin id")

	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	config := config.NewConfig()
	storage := storage.NewStorage("bins.json")

	api := api.NewAPI(config.Key)

	executeWithFlags(api, storage, create, update, get, delete, list, filename, binName, binId)
}

func executeWithFlags(api *api.API, storage *storage.FileStorage, create, update, get, delete, list *bool, filename, name, id *string) {
	if *create {
		if *filename == "" {
			panic("file is required")
		}

		fileReader := file.NewJSONFileReader(*filename)
		err := api.CreateBin(fileReader, *name, storage)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Bin created successfully")
		return
	}

	if *update {
		if *filename == "" || *id == "" {
			panic("file and id are required")
		}

		fileReader := file.NewJSONFileReader(*filename)
		err := api.UpdateBin(*id, fileReader)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Bin updated successfully")
		return
	}

	if *get {
		if *id == "" {
			panic("id is required")
		}

		err := api.GetBinList(*id)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Bin data fetched successfully")
		return
	}

	if *delete {
		if *id == "" {
			panic("id is required")
		}

		err := api.DeleteBin(*id, storage)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Bin deleted successfully")
		return
	}

	if *list {
		binList, err := storage.ReadBins()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("Bin list:")
		prettyJSON, err := json.MarshalIndent(binList, "", "  ")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(string(prettyJSON))
		fmt.Println("Bin list got successfully")
		return
	}
}

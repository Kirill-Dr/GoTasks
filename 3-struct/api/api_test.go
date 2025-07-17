package api_test

import (
	"3-struct/api"
	"3-struct/file"
	"3-struct/storage"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Failed to load .env file")
	}
}

func TestCreateBin(t *testing.T) {
	api := api.NewAPI(os.Getenv("KEY"))
	storage := storage.NewStorage("../test1_bins.json")
	fileReader := file.NewJSONFileReader("../test1_bins.json")
	binName := "Test Bin"

	resp, err := api.CreateBin(fileReader, binName, storage)
	if err != nil {
		t.Fatalf("Failed to create bin: %v", err)
	}

	respId := resp.Metadata.ID
	_, err = api.DeleteBinById(respId, storage)
	if err != nil {
		t.Fatalf("Failed to delete bin: %v", err)
	}
}

func TestGetBinById(t *testing.T) {
	api := api.NewAPI(os.Getenv("KEY"))
	storage := storage.NewStorage("../test2_bins.json")
	fileReader := file.NewJSONFileReader("../test2_bins.json")
	binName := "Test Bin"

	resp, err := api.CreateBin(fileReader, binName, storage)
	if err != nil {
		t.Fatalf("Failed to create bin: %v", err)
	}

	getResp, err := api.GetBinById(resp.Metadata.ID)
	if err != nil {
		t.Fatalf("Failed to get bin: %v", err)
	}

	if getResp.Metadata.ID != resp.Metadata.ID {
		t.Fatalf("Bin ID mismatch: expected %s, got %s", resp.Metadata.ID, getResp.Metadata.ID)
	}

	_, err = api.DeleteBinById(resp.Metadata.ID, storage)
	if err != nil {
		t.Fatalf("Failed to delete bin: %v", err)
	}
}

func TestUpdateBinById(t *testing.T) {
	api := api.NewAPI(os.Getenv("KEY"))
	storage := storage.NewStorage("../test3_bins.json")
	fileReader := file.NewJSONFileReader("../test3_bins.json")
	binName := "Test Bin"

	resp, err := api.CreateBin(fileReader, binName, storage)
	if err != nil {
		t.Fatalf("Failed to create bin: %v", err)
	}

	fileReader = file.NewJSONFileReader("../test3_updated_bins.json")
	updateData := `{
		"bins": [
			{
				"id": "1",
				"name": "Test Bin updated"
			}
		]
	}`
	err = os.WriteFile("../test3_updated_bins.json", []byte(updateData), 0644)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}

	_, err = api.UpdateBinById(resp.Metadata.ID, fileReader)
	if err != nil {
		t.Fatalf("Failed to update bin: %v", err)
	}

	os.Remove("../test3_updated_bins.json")

	_, err = api.DeleteBinById(resp.Metadata.ID, storage)
	if err != nil {
		t.Fatalf("Failed to delete bin: %v", err)
	}
}

func TestDeleteBinById(t *testing.T) {
	api := api.NewAPI(os.Getenv("KEY"))
	storage := storage.NewStorage("../test4_bins.json")
	fileReader := file.NewJSONFileReader("../test4_bins.json")
	binName := "Test Bin"

	resp, err := api.CreateBin(fileReader, binName, storage)
	if err != nil {
		t.Fatalf("Failed to create bin: %v", err)
	}

	_, err = api.DeleteBinById(resp.Metadata.ID, storage)
	if err != nil {
		t.Fatalf("Failed to delete bin: %v", err)
	}
}

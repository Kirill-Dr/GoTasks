package api

import (
	"3-struct/bins"
	"3-struct/file"
	"3-struct/storage"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var baseUrl = "https://api.jsonbin.io/v3/b/"

type API struct {
	Key string
}

type BinResponse struct {
	Metadata struct {
		ID string `json:"id"`
	} `json:"metadata"`
	Record any `json:"record"`
}

func NewAPI(key string) *API {
	return &API{
		Key: key,
	}
}

func (a *API) makeRequest(method, url string, headers map[string]string, body []byte) (*BinResponse, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(responseBody))
	}

	var response BinResponse
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	prettyJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Response:\n%s\n", string(prettyJSON))
	}

	return &response, nil
}

func (a *API) CreateBin(fileReader file.FileReader, binName string, storage *storage.FileStorage) (*BinResponse, error) {
	binList := bins.NewBinList()

	isPrivate := true
	name := binName

	newBin := bins.NewBin("1", isPrivate, name)
	binList.Bins = append(binList.Bins, *newBin)

	err := storage.SaveBins(binList)
	if err != nil {
		return nil, fmt.Errorf("failed to save bin locally: %v", err)
	}

	jsonData, err := fileReader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	resp, err := a.makeRequest("POST", baseUrl, map[string]string{
		"Content-Type": "application/json",
		"X-Master-Key": a.Key,
		"X-Bin-Name":   binName,
	}, jsonData)
	if err != nil {
		return nil, fmt.Errorf("failed to create bin: %v", err)
	}

	return resp, nil
}

func (a *API) GetBinById(binId string) (*BinResponse, error) {
	if binId == "" {
		return nil, fmt.Errorf("bin id is required")
	}

	url, err := url.Parse(baseUrl + binId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url")
	}

	resp, err := a.makeRequest("GET", url.String(), map[string]string{
		"X-Master-Key": a.Key,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get bin list: %v", err)
	}

	return resp, nil
}

func (a *API) UpdateBinById(binId string, fileReader file.FileReader) (*BinResponse, error) {
	if binId == "" {
		return nil, fmt.Errorf("bin id is required")
	}

	jsonData, err := fileReader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	resp, err := a.makeRequest("PUT", baseUrl+binId, map[string]string{
		"Content-Type": "application/json",
		"X-Master-Key": a.Key,
	}, jsonData)
	if err != nil {
		return nil, fmt.Errorf("failed to update bin: %v", err)
	}

	return resp, nil
}

func (a *API) DeleteBinById(binId string, storage *storage.FileStorage) (*BinResponse, error) {
	if binId == "" {
		return nil, fmt.Errorf("bin id is required")
	}

	err := os.Remove(storage.GetFilename())
	if err != nil {
		return nil, fmt.Errorf("failed to delete local file: %v", err)
	}

	resp, err := a.makeRequest("DELETE", baseUrl+binId, map[string]string{
		"X-Master-Key": a.Key,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to delete bin from server: %v", err)
	}

	return resp, nil
}

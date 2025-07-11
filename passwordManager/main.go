package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type BinList struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

func newBin(id string, private bool, name string) BinList {
	newBin := BinList{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}

	data, _ := os.ReadFile("BinList.json")
	var bins []BinList

	if len(data) > 0 {
		json.Unmarshal(data, &bins)
	}

	bins = append(bins, newBin)

	jsonData, _ := json.MarshalIndent(bins, "", "  ")
	os.WriteFile("BinList.json", jsonData, 0644)

	fmt.Println("Bin created and saved to BinList.json")
	return newBin
}

var binIdLetters = []rune("0123456789")

func generateBinID() string {
	binId := make([]rune, 10)
	for i := range binId {
		binId[i] = binIdLetters[rand.Intn(len(binIdLetters))]
	}
	return string(binId)
}

func getBinData() (private bool, name string) {
	var privateStr string
	fmt.Print("Enter private (true/false): ")
	fmt.Scanln(&privateStr)

	if privateStr == "true" {
		private = true
	} else {
		private = false
	}

	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	return private, name
}

func main() {
	fmt.Println("--- Password Manager ---")
	private, name := getBinData()
	bin := newBin(generateBinID(), private, name)
	fmt.Println(bin)
}

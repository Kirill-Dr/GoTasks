package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	Bins []Bin
}

func newBin(id string, private bool, name string) Bin {
	return Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func newBinList() BinList {
	return BinList{
		Bins: []Bin{},
	}
}

func addBinToList(binList *BinList, bin Bin) {
	binList.Bins = append(binList.Bins, bin)
}

func generateBinID() string {
	binId := make([]rune, 10)
	letters := []rune("0123456789")
	for i := range binId {
		binId[i] = letters[rand.Intn(len(letters))]
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
	fmt.Println("--- Password Manager CLI ---")

	binList := newBinList()

	private, name := getBinData()

	bin := newBin(generateBinID(), private, name)

	addBinToList(&binList, bin)

	fmt.Print(binList)
}

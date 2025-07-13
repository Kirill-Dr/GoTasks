package bins

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

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func NewBinList() BinList {
	return BinList{
		Bins: []Bin{},
	}
}

func AddBinToList(binList *BinList, bin Bin) {
	binList.Bins = append(binList.Bins, bin)
}

func GenerateBinID() string {
	binId := make([]rune, 10)
	letters := []rune("0123456789")
	for i := range binId {
		binId[i] = letters[rand.Intn(len(letters))]
	}
	return string(binId)
}

func GetBinData() (private bool, name string) {
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

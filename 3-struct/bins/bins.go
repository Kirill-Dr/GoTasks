package bins

import (
	"fmt"
	"math/rand"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
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

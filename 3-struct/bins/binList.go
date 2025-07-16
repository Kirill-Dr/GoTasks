package bins

import "time"

type BinList struct {
	Bins      []Bin     `json:"bins"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBinList() *BinList {
	return &BinList{
		Bins:      []Bin{},
		UpdatedAt: time.Now(),
	}
}

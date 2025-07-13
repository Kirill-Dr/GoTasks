package main

import (
	"3-struct/bins"
	"fmt"
)

func main() {
	fmt.Println("--- Password Manager CLI ---")

	binList := bins.NewBinList()

	private, name := bins.GetBinData()

	bin := bins.NewBin(bins.GenerateBinID(), private, name)

	bins.AddBinToList(&binList, bin)

	fmt.Print(binList)
}

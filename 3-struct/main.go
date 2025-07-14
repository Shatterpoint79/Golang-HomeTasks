package main

import (
	"api/binjson/bins"
	"fmt"
)

func main() {
	myBin := bins.NewBin("Secret Data", true)

	fmt.Printf("Создан Bin:\nID: %s\nName: %s\nPrivate: %t\nCreated: %s\n",
		myBin.ID,
		myBin.Name,
		myBin.Private,
		myBin.CreatedAt.Format("2006-01-02 15:04:05"),
	)

	binList := bins.NewBinList()

	binList.AddBin(myBin)
	binList.AddBin(bins.NewBin("Public Data", false))

	fmt.Printf("\nВ коллекции %d элементов:\n", len(binList.Bins))
	for i, b := range binList.Bins {
		fmt.Printf("%d: %s (ID: %s)\n", i+1, b.Name, b.ID)
	}
}

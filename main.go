package main

import (
	"cash/cashregister"
	"fmt"
)

var itemList []cashregister.Item

func init() {
	itemList = cashregister.AppendItem(itemList, "Elma", 0.75, 100.0)
	itemList = cashregister.AppendItem(itemList, "Armut", 1.0, 3.0)
	itemList = cashregister.AppendItem(itemList, "Karpuz", 13.5, 4.9)
	itemList = cashregister.AppendItem(itemList, "Karpuz", 100, 50)
	itemList = cashregister.AppendItem(itemList, "Elmas", 1000, 0.0)
}

func main() {

	for _, item := range itemList {
		fmt.Println(item.Description())
	}

	fmt.Printf("Toplam fiyat %.2f:", cashregister.TotalPrice(itemList))

}

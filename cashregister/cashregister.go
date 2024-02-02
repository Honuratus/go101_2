package cashregister

import (
	"fmt"
	"log"
)

type Describable interface {
	Description() string
}

type Item struct {
	Name     string
	Price    float64
	discount float64 // discount should be in range (0, 100)
}

// get function for discount
func (i Item) GetDiscount() float64 {
	return i.discount
}

// set function for discount
func (i *Item) SetDiscount(dc float64) {
	if dc < 0.0 || dc > 100.0 {
		log.Fatalf("You typed not valid discount amount\n")
	}
	(*i).discount = dc

}

func (i Item) Description() string {
	if isDiscountZero(i) {
		return fmt.Sprintf("%s - %v TL\n", i.Name, i.Price)
	}
	return fmt.Sprintf("%s - %v TL (%.2f%c indirimle %.2f TL)", i.Name, i.Price, i.GetDiscount(), '%', CalculatePrice(i))
}

// calcuting the discounted price of an item
func CalculatePrice(item Item) float64 {
	if isDiscountZero(item) { // checks if discount's value assigned to zero
		return item.Price
	}
	// calculations
	discountAmount := item.Price * (item.GetDiscount() / 100)
	calculatedPrice := item.Price - discountAmount
	return calculatedPrice
}

// calculating the total discounted price
func TotalPrice(items []Item) float64 {
	if len(items) == 0 {
		return 0
	}
	var sum float64
	sum += TotalPrice(items[1:]) + CalculatePrice(items[0])
	return sum
}

// item append function
func AppendItem(items []Item, name string, price float64, discount float64) []Item {
	newItem := Item{
		Name:  name,
		Price: price,
	}
	newItem.SetDiscount(discount)
	items = append(items, newItem)
	return items
}

func isDiscountZero(item Item) bool {
	if item.GetDiscount() == 0.0 {
		return true
	}
	return false
}

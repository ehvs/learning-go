package main

import "fmt"

// maps: map[KeyType]ElementType
// foo := map[string]int{}
// value, exists := foo["baz"]

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728}
	return units
}

func NewBill() map[string]int {
	return make(map[string]int)
}

func AddItem(bill, units map[string]int, item, unit string) bool {
	/*Return false if the given unit is not in the units map.
	Otherwise, add the item to the customer bill, indexed by the item name, then return true.
	If the item is already present in the bill, increase its quantity by the amount that belongs to the provided unit.
	*/
	if numUnit, exists := units[unit]; exists {
		fmt.Println("numUnit?", numUnit, "|| Exists or Not?", exists, "|| What is in units[unit]?", units[unit], " || What is in bill[item]? ", bill[item], "|| what is the content of map bill?", bill)
		// before bill["carrot":0] now: bill["carrot":12]
		bill[item] = bill[item] + numUnit
		fmt.Println("What is 'item'?", item, "|| What is in bill[item]?", bill[item], "|| what is the content of map bill?", bill)
		return true
	} else {
		return false
	}
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	/*Return false if the given item is not in the bill
	  Return false if the given unit is not in the units map.
	  Return false if the new quantity would be less than 0.
	  If the new quantity is 0, completely remove the item from the bill then return true.
	  Otherwise, reduce the quantity of the item and return true.*/
	numUnit, unitExists := units[unit]
	fmt.Println("What has inside units[unit]? ", units[unit], " || numUnit?", numUnit, " || unitExist?", unitExists)
	if !unitExists {
		return false
	}
	itemQuantity, itemExists := bill[item]
	fmt.Println("What has inside bill[item]? ", bill[item], " || itemQuantity?", itemQuantity, " || itemExists?", itemExists)
	if !itemExists {
		return false
	}
	itemQuantity -= numUnit
	if itemQuantity < 0 {
		return false
	}
	if itemQuantity == 0 {
		delete(bill, item)
		return true
	}
	bill[item] = itemQuantity
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if thingy, exists := bill[item]; exists {
		return thingy, exists
	}
	return 0, false
}

func main() {
	//units := Units()
	//bill := NewBill()
	//ok := AddItem(bill, units, "carrot", "dozen")
	//ok := RemoveItem(bill, units, "carrot", "dozen")
	//fmt.Println(ok, bill)
	bill := map[string]int{"carrot": 12, "grapes": 3}
	qty, ok := GetItem(bill, "carrot")
	fmt.Println(qty)
	// Output: 12
	fmt.Println(ok)
}

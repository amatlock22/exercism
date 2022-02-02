package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	u, ok := units[unit]
	if !ok {
		return false
	}

	bill[item] += u

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	quantity, itemOk := bill[item]
	u, unitOk := units[unit]

	if !itemOk || !unitOk {
		return false
	}

	newQuantity := quantity - u
	if newQuantity < 0 {
		return false
	}

	if newQuantity == 0 {
		delete(bill, item)

		return true
	}

	bill[item] = newQuantity

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if quantity, itemOk := bill[item]; !itemOk {
		return 0, false
	} else {
		return quantity, true
	}
}

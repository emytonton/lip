package gross

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

func NewBill() map[string]int {
	return map[string]int{}
}

func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, exists := units[unit]
	if !exists {
		return false
	}
	bill[item] += quantity
	return true
}

func RemoveItem(bill, units map[string]int, item, unit string) bool {
	quantity, exists := units[unit]
	current, itemExists := bill[item]

	if !exists || !itemExists {
		return false
	}

	if current < quantity {
		return false
	}

	bill[item] = current - quantity

	if bill[item] == 0 {
		delete(bill, item)
	}

	return true
}

func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	return quantity, exists
}

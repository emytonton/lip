package sublist

func contains(bigger, smaller []int) bool {
	lenBigger := len(bigger)
	lenSmaller := len(smaller)

	if lenSmaller == 0 {
		return true
	}

	if lenSmaller > lenBigger {
		return false
	}

	for i := 0; i <= lenBigger-lenSmaller; i++ {

		match := true
		for j := 0; j < lenSmaller; j++ {

			if bigger[i+j] != smaller[j] {
				match = false
				break
			}
		}

		if match {
			return true
		}
	}

	return false
}

func Sublist(l1, l2 []int) Relation {
	l1ContainsL2 := contains(l1, l2)
	l2ContainsL1 := contains(l2, l1)

	if l1ContainsL2 && l2ContainsL1 {
		return RelationEqual
	}

	if l2ContainsL1 {
		return RelationSublist
	}

	if l1ContainsL2 {
		return RelationSuperlist
	}

	return RelationUnequal
}

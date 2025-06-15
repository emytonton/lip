// Package listops implements basic list operations for a custom integer list type.
package listops

// IntList is an abstraction of a list of integers which we can define methods on.
type IntList []int

// binFunc is a binary function taking two integers and returning an integer.
type binFunc func(int, int) int

// predFunc is a predicate function taking an integer and returning a boolean.
type predFunc func(int) bool

// unaryFunc is a unary function taking an integer and returning an integer.
type unaryFunc func(int) int

// Foldl reduces the list from the left, applying the function to an accumulator and each element.
// The function is called as fn(accumulator, element).
func (s IntList) Foldl(fn binFunc, initial int) int {
	accumulator := initial
	for _, v := range s {
		accumulator = fn(accumulator, v)
	}
	return accumulator
}

// Foldr reduces the list from the right, applying the function to an accumulator and each element.
// The function is called as fn(element, accumulator).
func (s IntList) Foldr(fn binFunc, initial int) int {
	accumulator := initial
	// Iterate from the last element to the first.
	for i := len(s) - 1; i >= 0; i-- {
		accumulator = fn(s[i], accumulator)
	}
	return accumulator
}

// Filter returns a new list containing only the elements for which the predicate is true.
func (s IntList) Filter(fn predFunc) IntList {
	// Pre-allocate a slice with a capacity of 0, as the final length is unknown.
	filtered := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// Length returns the total number of items in the list.
func (s IntList) Length() int {
	// The built-in len() is the most idiomatic and efficient way to get the length of a slice.
	return len(s)
}

// Map returns a new list with the results of applying the function to each element.
func (s IntList) Map(fn unaryFunc) IntList {
	// Pre-allocate a slice of the same size to avoid reallocations.
	mapped := make(IntList, len(s))
	for i, v := range s {
		mapped[i] = fn(v)
	}
	return mapped
}

// Reverse returns a new list with the original items in reversed order.
func (s IntList) Reverse() IntList {
	length := len(s)
	reversed := make(IntList, length)
	for i, v := range s {
		// Place the element from the start of the original list
		// at the end of the new list, and so on.
		reversed[length-1-i] = v
	}
	return reversed
}

// Append adds all items from the second list to the end of the first list.
func (s IntList) Append(lst IntList) IntList {
	// The built-in append function is the standard way to do this in Go.
	return append(s, lst...)
}

// Concat combines all items from a series of lists into one flattened list.
func (s IntList) Concat(lists []IntList) IntList {
	result := s
	for _, lst := range lists {
		result = append(result, lst...)
	}
	return result
}

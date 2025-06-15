package strain

func Keep[T any](collection []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Discard[T any](collection []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range collection {
		if !predicate(item) {
			result = append(result, item)
		}
	}
	return result
}
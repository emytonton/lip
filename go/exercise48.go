package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	result := make(map[string]int)

	for score, letters := range in {
		for _, letter := range letters {
			lowerLetter := strings.ToLower(letter)
			result[lowerLetter] = score
		}
	}

	return result
}

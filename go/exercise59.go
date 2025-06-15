package anagram

import (
	"sort"
	"strings"
)

func sortRunes(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}


func Detect(subject string, candidates []string) []string {
	
	var anagrams []string

	
	lowerSubject := strings.ToLower(subject)
	sortedSubject := sortRunes(lowerSubject)

	
	for _, candidate := range candidates {
	
		lowerCandidate := strings.ToLower(candidate)

		
		if lowerCandidate == lowerSubject {
			continue
		}

		
		if len([]rune(lowerCandidate)) != len([]rune(lowerSubject)) {
			continue
		}

	
		if sortRunes(lowerCandidate) == sortedSubject {
			
			anagrams = append(anagrams, candidate)
		}
	}

	return anagrams
}

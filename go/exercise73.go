package bottlesong

import (
	"fmt"
	"unicode"
)


func Capitalize(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func numberToWord(n int) string {
	switch n {
	case 10:
		return "ten"
	case 9:
		return "nine"
	case 8:
		return "eight"
	case 7:
		return "seven"
	case 6:
		return "six"
	case 5:
		return "five"
	case 4:
		return "four"
	case 3:
		return "three"
	case 2:
		return "two"
	case 1:
		return "one"
	case 0:
		return "no" 
	default:
		return ""
	}
}


func getBottleString(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}


func generateVerse(n int) []string {
	verse := make([]string, 4)

	currentNumStr := numberToWord(n)
	currentBottleStr := getBottleString(n)
	line := fmt.Sprintf("%s green %s hanging on the wall,", Capitalize(currentNumStr), currentBottleStr)


	verse[0] = line
	verse[1] = line


	verse[2] = "And if one green bottle should accidentally fall,"

	
	nextNum := n - 1
	nextNumStr := numberToWord(nextNum)
	nextBottleStr := getBottleString(nextNum)
	verse[3] = fmt.Sprintf("There'll be %s green %s hanging on the wall.", nextNumStr, nextBottleStr)

	return verse
}


func Recite(startBottles, takeDown int) []string {

	song := make([]string, 0, takeDown*5)

	for i := 0; i < takeDown; i++ {
		currentBottles := startBottles - i
		verse := generateVerse(currentBottles)
		song = append(song, verse...)

		if i < takeDown-1 {
			song = append(song, "")
		}
	}

	return song
}

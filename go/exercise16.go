package logs

import "unicode/utf8"

func Application(log string) string {
	for _, char := range log {
		switch char {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

func Replace(log string, oldRune, newRune rune) string {
	var result []rune
	for _, char := range log {
		if char == oldRune {
			result = append(result, newRune)
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}

func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
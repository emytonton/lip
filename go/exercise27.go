package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	result := make([]string, len(lines))
	for i, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			result[i] = "[USR] " + matches[1] + " " + line
		} else {
			result[i] = line
		}
	}
	return result
}
package series

func All(n int, s string) []string {
	if n > len(s) || n <= 0 {
		return []string{}
	}

	result := make([]string, 0, len(s)-n+1)
	for i := 0; i <= len(s)-n; i++ {
		sub := s[i : i+n]
		result = append(result, sub)
	}
	return result
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) || n <= 0 {
		return "", false
	}
	return s[0:n], true
}
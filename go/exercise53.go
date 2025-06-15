package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	if shiftKey == 0 || shiftKey == 26 {
		return plain
	}

	result := make([]rune, len(plain))
	for i, char := range plain {
		if char >= 'A' && char <= 'Z' {
			shifted := int(char-'A') + shiftKey
			shifted = shifted % 26
			result[i] = 'A' + rune(shifted)
		} else if char >= 'a' && char <= 'z' {
			shifted := int(char-'a') + shiftKey
			shifted = shifted % 26
			result[i] = 'a' + rune(shifted)
		} else {
			result[i] = char
		}
	}
	return string(result)
}
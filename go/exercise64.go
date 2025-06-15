package cipher

import (
	"strings"
	"unicode"
)


type shift struct {
	distance int
}


func NewCaesar() Cipher {
	return NewShift(3)
}


func NewShift(distance int) Cipher {
	if distance == 0 || distance > 25 || distance < -25 {
		return nil
	}
	return shift{distance}
}


func (c shift) Encode(input string) string {
	return shiftString(input, c.distance)
}


func (c shift) Decode(input string) string {
	return shiftString(input, -c.distance)
}



type vigenere struct {
	key string
}


func NewVigenere(key string) Cipher {
	if len(key) == 0 {
		return nil
	}

	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
	}


	allA := true
	for _, r := range key {
		if r != 'a' {
			allA = false
			break
		}
	}
	if allA {
		return nil
	}

	return vigenere{key}
}


func (v vigenere) Encode(input string) string {
	return vigenereString(input, v.key, true)
}


func (v vigenere) Decode(input string) string {
	return vigenereString(input, v.key, false)
}


func shiftString(input string, shift int) string {
	var sb strings.Builder

	for _, r := range input {
		if unicode.IsLetter(r) {
			lower := unicode.ToLower(r)
			shifted := ((int(lower-'a') + shift + 26) % 26) + 'a'
			sb.WriteRune(rune(shifted))
		}
	}

	return sb.String()
}


func vigenereString(input, key string, encode bool) string {
	var sb strings.Builder
	keyLen := len(key)
	keyIndex := 0

	for _, r := range input {
		if unicode.IsLetter(r) {
			lower := unicode.ToLower(r)
			keyChar := key[keyIndex%keyLen]
			shift := int(keyChar - 'a')
			if !encode {
				shift = -shift
			}
			shifted := ((int(lower-'a') + shift + 26) % 26) + 'a'
			sb.WriteRune(rune(shifted))
			keyIndex++
		}
	}

	return sb.String()
}

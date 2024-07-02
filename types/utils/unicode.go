package utils

import "unicode/utf8"

// ReplaceInvalidUTF8 replaces invalid UTF-8 characters in a string with a replacement character.
func ReplaceInvalidUTF8(s string) string {
	result := make([]rune, 0, len(s))
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == utf8.RuneError && size == 1 {
			result = append(result, '?')
		} else {
			result = append(result, r)
		}
		i += size
	}
	return string(result)
}

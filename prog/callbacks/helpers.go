package callbacks

import "fmt"

func hasRepeatedChars(val string) bool {
	m := make(map[string]bool)
	for _, character := range val {
		char := fmt.Sprintf("%c", character)
		e := m[char]
		if e {
			return true
		}

		m[char] = true
	}

	return false
}

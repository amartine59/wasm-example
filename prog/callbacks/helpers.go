package callbacks

import (
	"fmt"
	"syscall/js"
)

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

func updateResultBoxByID(boxID, v string) {
	js.Global().Get("document").Call("getElementById", boxID).Set("value", v)
}

package callbacks

import (
	"fmt"
	"syscall/js"
)

func HasRepeatedChars(this js.Value, args []js.Value) any {
	if len(args) > 1 {
		fmt.Println("This callback only needs a single argument")
		return nil
	}

	val := args[0].String()

	r := hasRepeatedChars(val)
	if r {
		fmt.Println("Yes, it does")
		return nil
	}

	fmt.Println("No, it does not")
	return nil
}

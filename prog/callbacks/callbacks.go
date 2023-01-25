package callbacks

import (
	"fmt"
	"syscall/js"
)

func HasRepeatedChars(this js.Value, args []js.Value) any {
	if len(args) > 2 {
		fmt.Println("This callback only requires two arguments. One for the id of the input and another one for its result")
		return nil
	}

	resId := args[1].String()

	valId := args[0].String()

	inp := js.Global().Get("document").Call("getElementById", valId).Get("value").String()
	r := hasRepeatedChars(inp)
	if r {
		fmt.Println("Yes, there is at least character repeating once")
		updateResultBoxByID(resId, "Yes, there is at least character repeating once")

		return true
	}

	fmt.Println("No, none of the characters is repeating")
	updateResultBoxByID(resId, "No, none of the characters is repeating")
	return false
}

func Clear(this js.Value, args []js.Value) any {
	for _, arg := range args {
		if arg.IsUndefined() || arg.IsNull() {
			continue
		}

		updateResultBoxByID(arg.String(), "")
	}

	return nil
}

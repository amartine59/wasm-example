package main

import (
	"fmt"
	"github/amartine59/wasm-example/prog/callbacks"
	"syscall/js"
)

func main() {
	ch := make(chan struct{})
	fmt.Println("WASM Program Online!")

	js.Global().Set("hasRepeatedChars", js.FuncOf(callbacks.HasRepeatedChars))

	<-ch
}

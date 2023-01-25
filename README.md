## WASM Example 🛠️

This is a very simple calculator that can help us get a small taste of WebAssembly using Go.
I will be using this code in a post(I'm almost certain).

## References 📑
I was able to put together this exercise after looking at the following resources(and doing some research on the subject as well):
-  [Go WebAssembly Tutorial - TutorialEdge](https://tutorialedge.net/golang/go-webassembly-tutorial/)
- [Go Wiki - WebAssembly](https://github.com/golang/go/wiki/WebAssembly#webassembly)
- [TinyGo - WebAssembly](https://tinygo.org/docs/guides/webassembly/)

## Running the exercise ⚙️
It's probably obvious at this point but you need to have go installed in your machine.
Apart from that, if you're using VS Code like me you will also need to set the following environment variables:
- `GOOS=js`
- `GOARCH=wasm`

This is due to an experimental package we'll be using called `syscall/js` that allows us to export our functions and make them available through Javascript and also helps us do some DOM manipulation. Without those env variables I was getting the following error:

> imports syscall/js: build constraints exclude all Go files in /usr/local/go/src/syscall/js

That being said, the first thing you'll need to do is to compile the side-program to *WebAssembly* by running the following command:

```
GOOS=js GOARCH=wasm go build -o ./wasm.wasm ./prog/main.go
```
That'll generate a file called `wasm.wasm` which will then be consumed by the JS side as a module that can be called from the client side.
If you're curious you can explore the content of the resulting file using a WebAssembly extension like [this one](https://marketplace.visualstudio.com/items?itemName=dtsvet.vscode-wasm)

After that, you just need to start the server:
```
go run server.go
```
And then go to `localhost:8080`.

From that point you should be able to interact with the go program by pressing the button in the UI or by using the dev tools of your browser to call the function that was exported.

## Support ⚠️
For now, this code should only work on Chrome and Mozilla Firefox.

## Acknowledgements
Made with ❤️ by Wawandco

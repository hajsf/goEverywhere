//go:build wasm
// +build wasm

package main

// As the file name hasing the prefex _wasm, it will be compiled only if the GOARC is wasm, so the // +build wasm here can be excluded

// if using TinyGo for compiling: tinygo build -o wasm.wasm -target wasm ./main.go
//
//export hello
func hello() string {
	return "Hello"
}

func main() {
	c := make(chan int) // channel to keep the wasm running, it is not a library as in rust/c/c++, so we need to keep the binary running
	println("Hello wasm")
	<-c // pause the execution so that the resources we create for JS keep available
}

// compile to wasm:
// GOOS=js GOARCH=wasm go build -o www/wasm/main.wasm github.io/hajsf/wasm
// Copied the wasm_exec.js file to the same working folder as:
// cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./wasm

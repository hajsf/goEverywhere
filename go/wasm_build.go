//go:build !wasm
// +build !wasm

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	var package_name = "goApp"
	var pwd = "D:/goEverywhere/go"
	wasmOut := pwd + "/lib/wasm"

	// go tool dist list

	os.Setenv("GOOS", "js")
	os.Setenv("GOARCH", "wasm")
	cmd := exec.Command("go", "build", "-o", wasmOut+"/lib"+package_name+".wasm", package_name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %v:\n\noutput:\n\n%s\n", err, out)
	}
	// Copied the wasm_exec.js file to the same working folder as:
	// cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./wasm
	goEnv := "C:/Program Files/Go" // $(go env GOROOT)

	sourceFile := goEnv + "/misc/wasm/wasm_exec.js"
	destinationFile := wasmOut + "/wasm_exec.js"

	nBytes, err := copy(sourceFile, destinationFile)
	if err != nil {
		fmt.Printf("The copy operation failed %q\n", err)
	} else {
		fmt.Printf("Copied %d bytes!\n", nBytes)
	}
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

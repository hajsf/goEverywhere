//go:build tag_build
// +build tag_build

package main

import (
	"fmt"
	"os"
	"os/exec"
)

type compiles struct {
	sys  string
	arch string
}

func main() {
	var package_name = "goApp"
	var pwd = "D:/goEverywhere/go"
	var output string

	// go tool dist list
	systems := []compiles{
		compiles{
			sys:  "darwin",
			arch: "amd64",
		},
		compiles{
			sys:  "linux",
			arch: "amd64",
		},
		compiles{
			sys:  "windows",
			arch: "amd64",
		},
	}

	for _, s := range systems {
		os.Setenv("GOOS", s.sys)
		os.Setenv("GOARCH", s.arch)
		switch println(s.sys); s.sys {
		case "darwin":
			output = pwd + "/lib/desktop/darwin/" + package_name + ".dmg"
		case "windows":
			output = pwd + "/lib/desktop/windows/" + package_name + ".exe"
		case "linux":
			output = pwd + "/lib/desktop/linux/" + package_name
		default:
			println("Undefined")
		}
		cmd := exec.Command("go", "build", "-o", output, package_name)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("cmd.Run() failed with %v:\n\noutput:\n\n%s\n", err, out)
		}
	}
}

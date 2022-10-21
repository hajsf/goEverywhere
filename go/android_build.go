//go:build tag_build
// +build tag_build

package main

import (
	"fmt"
	"os"
	"os/exec"
)

type platform struct {
	arch     string
	compiler string
}

type system struct {
	name string
	platform
}

func main() {
	var package_name = "goApp"
	var pwd = "D:/goEverywhere/go"

	androidOut := pwd + "/lib/android"
	var androidSDK = "C:/Users/hasan_yousef/AppData/Local/Android/Sdk"
	var ndk = androidSDK + "/ndk/25.1.8937393/toolchains/llvm/prebuilt/windows-x86_64/bin"
	os.Setenv("GOOS", "android")
	os.Setenv("CGO_ENABLED", "1")
	// go tool dist list
	android := []system{
		system{
			name: "android-armv7a",
			platform: platform{
				arch:     "arm",
				compiler: "armv7a-linux-androideabi33-clang",
			},
		},
		system{
			name: "android-arm64",
			platform: platform{
				arch:     "arm64",
				compiler: "aarch64-linux-android33-clang",
			},
		},
		system{
			name: "android-x86",
			platform: platform{
				arch:     "386",
				compiler: "i686-linux-android33-clang",
			},
		},
		system{
			name: "android-x86_64",
			platform: platform{
				arch:     "amd64",
				compiler: "x86_64-linux-android33-clang",
			},
		},
	}

	for _, s := range android {
		switch s.name {
		case "android-armv7a":
			os.Setenv("GOARM", "7")
		default:
			os.Setenv("GOARM", "")
		}
		os.Setenv("GOARCH", s.platform.arch)
		os.Setenv("CC", ndk+"/"+s.platform.compiler)
		args := []string{"build", "-ldflags=-w",
			"-buildmode=c-shared", "-o",
			androidOut + "/" + s.name + "/lib" + package_name + ".so", package_name}
		//	cmd := exec.Command("go", "build", "-ldflags=-w", "-buildmode=c-shared", "-o", androidOut+"/"+s.name+"/lib"+package_name+".so", package_name)
		cmd := exec.Command("go", args...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("cmd.Run() failed with %v:\n\noutput:\n\n%s\n", err, out)
		}
	}
}

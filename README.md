To check supported compilation tools: `go tool dist list`
Cross compiling not working with `cgo`, i.e. not working with any file has `import "C"`

https://go.dev/misc/android/README

At Powershell
```powersehl
# For ARM
$env:GOOS = "linux" 
$env:GOARCH = "arm" 
$env:GOARM = "7" 
go build -o main main.go
# For darwin
$env:GOOS = "darwin" 
$env:GOARCH = "amd64" 
go build -o main.dmg main.go
# same for others
$env:GOOS = "windows" 
$env:GOARCH = "amd64" 
go build -o main.exe main.go
```

At CMD `Command Prompt`:
```bash
set GOOS=darwin
set GOARCH=amd64
go build -o main.dmg main.go
```

To do it in Linux or Mac and compiling to Win
```bash
GOOS=windows GOARCH=amd64 go build -o main.exe main.go
```

Cross compiling will silently rebuild most of standard library, and for this reason will be quite slow. To speed-up the process, you can install all the standard packages required for cross compiling on your system, for example to install at Linux/Mac the cross compiling requirements for `windows-amd64` use:
```bash
GOOS=windows GOARCH=amd64 go install
```
Similar for any other OS you need to cross compile for it at Windows

With useg `// +build tag_name` we can run using `go build -tags tag_name`
To run a single file not the full package use the file name like `go run build.go`
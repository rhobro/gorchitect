# gorchitect

#### Gorchitect takes a go program and compiles it for every available GOOS and GOARCH. This will allow you to do this in one go instead of repetitively compiling.

### Installation

- `go build -i gorchitect.go` to compile the program (run `go install -i gorchitect.go` or move the binary to your PATH
  to make it accessible from anywhere)

### Usage

- `./gorchitect --help` to list and see description of arguments
- `-n` flag controls the number of goroutines used to compile the programs.
- `-o` flag used to provide path to directory where binaries should be compiled to.
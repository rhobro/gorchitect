# gorchitect

#### Gorchitect takes a Go program and concurrently compiles it for every available GOOS and GOARCH. This will allow you to do this in one go instead of repetitively compiling.

### Installation

- `go build -i gorchitect.go` to compile the program (run `go install -i gorchitect.go` or move the binary to your PATH
  to make it accessible from anywhere)

### Usage

- `./gorchitect --help` to list and see description of arguments.
- `-o` flag used to provide path to directory where binaries should be compiled to.
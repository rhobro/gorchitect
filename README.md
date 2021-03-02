# gorchitect

Gorchitect takes a Go program and concurrently compiles it for every available GOOS and GOARCH. This will allow you to
do this in one go instead of repetitively compiling.

### Installation

- `go build -i gorchitect.go` to compile the program (run `go install -i gorchitect.go` or move the binary to your PATH
  to make it accessible from anywhere)

### Usage

- `./gorchitect --help` to list and see description of arguments.
- `-o` flag used to provide a path to directory where binaries should be compiled to.
- `-n` flag user to control number of concurrent compilations (defaults to number of CPUs).

### Contribute

If you have a suggestion or have an improvement, please feel free to open an issue or pull request and I'll gladly take
it into account.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var path string
var out string

func init() {
	// args and flags
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("can't get working directory")
	}
	flag.StringVar(&out, "o", wd, "output path to which Go binaries should be compiled to")
	flag.Parse()
	path = os.Args[len(os.Args)-1]
}

func main() {
	// get list of GOOSs and GOARCHes and compile for each
	rawList := strings.TrimSpace(execute(exec.Command("go", "tool", "dist", "list")))
	for _, combo := range strings.Split(rawList, "\n") {
		spl := strings.Split(combo, "/")
		goos, goarch := spl[0], spl[1]
		fmt.Printf("%s : %s\n", goos, goarch)
	}
}

func execute(c *exec.Cmd) (s string) {
	out, _ := c.StdoutPipe()
	c.Start()
	bd, _ := ioutil.ReadAll(out)

	return string(bd)
}

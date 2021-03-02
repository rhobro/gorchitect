package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

	// mkdir out
	err = os.MkdirAll(out, os.ModePerm)
	if err != nil {
		log.Fatal("can't create output directory")
	}
}

func main() {
	// get executable name
	name := filepath.Base(path)
	if strings.Contains(name, ".") {
		name = name[:strings.LastIndex(name, ".")]
	}

	// get list of GOOSs and GOARCHes and compile for each
	rawList := strings.TrimSpace(execute(exec.Command("go", "tool", "dist", "list")))
	for _, combo := range strings.Split(rawList, "\n") {
		spl := strings.Split(combo, "/")
		goOS, goArch := spl[0], spl[1]

		out := filepath.Join(out, fmt.Sprintf("%s_%s_%s", out, goOS, goArch))
		cmd := exec.Command("go", "build", "-o", out, path)
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "GOOS="+goOS)
		cmd.Env = append(cmd.Env, "GOARCH="+goArch)

		err := cmd.Run()
		if err != nil {
			log.Print(err)
		}
	}
}

func execute(c *exec.Cmd) (s string) {
	out, _ := c.StdoutPipe()
	c.Start()
	bd, _ := ioutil.ReadAll(out)
	c.Wait()

	return string(bd)
}

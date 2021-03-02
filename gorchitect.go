package main

import (
	"flag"
	"fmt"
	"github.com/rhobro/goutils/pkg/util"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var path string
var out string
var nProcesses int

func init() {
	// args and flags
	wd, err := os.Getwd()
	util.Check(err)
	flag.StringVar(&out, "o", wd, "output path to which Go binaries should be compiled to")
	flag.IntVar(&nProcesses, "n", 5, "number of goroutines to use to concurrently compile")
	flag.Parse()

	path = flag.Arg(0)
	semaphore = make(chan struct{}, nProcesses)

	// mkdir out
	util.Check(os.MkdirAll(out, os.ModePerm))
}

// semaphore to limit concurrency
var semaphore chan struct{}

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

		semaphore <- struct{}{}
		go func() {
			err := cmd.Run()
			if err != nil {
				log.Print(err)
			}
			<-semaphore
		}()
	}
}

func execute(c *exec.Cmd) string {
	out, err := c.StdoutPipe()
	util.Check(err)
	util.Check(c.Start())
	bd, _ := ioutil.ReadAll(out)
	util.Check(c.Wait())

	return string(bd)
}

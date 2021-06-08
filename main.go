package main

import (
    "os/exec"
    "flag"
    "strings"
    "fmt"
    "sync"
)

func main() {
    flag.Parse()
    var wg sync.WaitGroup
    cmd := flag.Arg(0)
    cmd = strings.Trim(cmd, "\"")
    cmds := strings.Split(cmd, " ")
    wg.Add(1)
    go func() {
	defer wg.Done()
	c := exec.Command(cmds[0], cmds[1:]...)
        c.Start()
	c.Wait()
	fmt.Println("EVENT: Process is finished")
    }()
    wg.Wait()
}

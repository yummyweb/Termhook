package main

import (
    "os/exec"
    "os"
    "os/signal"
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
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func(){
    	for range c {
	    fmt.Println("EVENT: Process interrupted")    
	}
    }()
    go func() {
	defer wg.Done()
	c := exec.Command(cmds[0], cmds[1:]...)
        c.Start()
	fmt.Println("EVENT: Process started")
	c.Wait()
	fmt.Println("EVENT: Process is finished")
    }()
    wg.Wait()
}

package main

import (
	"net"
	"os"
	"strings"
)

const socketAddr = "/var/run/browser.sock"

func main() {
	c, err := net.Dial("unix", socketAddr)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Write([]byte(strings.Join(os.Args[1:], " ")))
}

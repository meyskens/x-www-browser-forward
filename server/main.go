package main

import (
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

const socketAddr = "/var/run/browser.sock"

func serveSocket(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := string(buf[0:nr])
		exec.Command("x-www-browser", strings.Split(data, " ")...).Output()
	}
}

func main() {
	l, err := net.Listen("unix", socketAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer os.Remove(socketAddr)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go serveSocket(conn)
	}
}

package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

const defaultSocketAddr = "/var/run/browser.sock"
const defaultBrowserCmd = "x-www-browser"

type Config struct {
	socketAddr string
	browserCmd string
}

func serveSocket(c net.Conn, browserCmd string) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := string(buf[0:nr])
		exec.Command(browserCmd, strings.Split(data, " ")...).Output()
	}
}

// Check if a file exists directly or in PATH
func checkIfFileExistsInPath(file string) bool {
	// Check if the file exists as it
	_, err := os.Stat(file)
	if err == nil {
		return true
	}

	// Check if the file exists in PATH
	for _, path := range strings.Split(os.Getenv("PATH"), ":") {
		file := []string{path, "/", file}
		_, err = os.Stat(strings.Join(file, ""))
		if err == nil {
			return true
		}
	}

	return false
}

// Get command line arguments
func getArgs() Config {
	flags := flag.NewFlagSet("x-www-browser-forwarder", flag.ExitOnError)
	browserCmd := flags.String("browser-cmd", defaultBrowserCmd, "Command to open URL")
	socketFile := flags.String("socket-file", defaultSocketAddr, "Socket address")
	flags.Parse(os.Args[1:])

	if checkIfFileExistsInPath(*browserCmd) != true {
		log.Fatal("Browser command not found: ", *browserCmd)
	}

	conf := Config{
		socketAddr: *socketFile,
		browserCmd: *browserCmd,
	}

	return conf
}

func main() {
	config := getArgs()

	l, err := net.Listen("unix", config.socketAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer os.Remove(config.socketAddr)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go serveSocket(conn, config.browserCmd)
	}
}

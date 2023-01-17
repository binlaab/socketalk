package main

// TODO: canales (probablemente lo deje para más tarde)

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
)

const usage = `Usage:
	-a, --address: IP address where the server will listen for connections. Defaults to all interfaces (0.0.0.0).
	-p, --port: Port where the server will listen for connection. Defaults to port 8000.
`

const TYPE = "tcp"

var IP = "0.0.0.0"
var PORT = "8000"

func main() {

	flag.StringVar(&IP, "a", IP, "placeholder")
	flag.StringVar(&IP, "address", IP, "placeholder")

	flag.StringVar(&PORT, "p", PORT, "placeholder")
	flag.StringVar(&PORT, "port", PORT, "placeholder")
	flag.Usage = func() { fmt.Print(usage) }
	flag.Parse()

	addr := IP + ":" + PORT
	ln, err := net.Listen(TYPE, addr)
	if err != nil {
		fmt.Printf("[ERROR] Could not bind to port. See the full error here: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Listening on port", PORT)
	go eval_admin()

	go func() { // catches ctrl + c
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		close_conns()
		os.Exit(0)
	}()

	for {
		conn, err := ln.Accept()
		ip := strings.Split(conn.RemoteAddr().String(), ":")[0]
		if check_banned(ip) {
			conn.Close()
			fmt.Printf("\rConnection received from %s, but it is banned.\n>>> ", ip)
			continue
		}
		if err != nil {
			fmt.Printf("\r[ERROR] Could not accept connection. See the full error here: %s\n>>> ", err.Error())
		}
		fmt.Printf("\rConnection received from %s\n>>> ", ip)
		go handle_connection(conn)
	}
}

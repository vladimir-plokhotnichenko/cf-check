package main

import (
	"fmt"
	"net"
	"os"
)

func show(host string, ip net.IP, mode bool) {
	if mode {
		fmt.Println(host)
	} else {
		fmt.Println(ip)
	}
}

func isStdin() bool {
	f, e := os.Stdin.Stat()
	if e != nil {
		return false
	}

	if f.Mode()&os.ModeNamedPipe == 0 {
		return false
	}

	return true
}

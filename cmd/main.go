package main

import (
	"fmt"
	"github.com/billopark/iep.ee/internal/dnsserver"
	"github.com/billopark/iep.ee/internal/webserver"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	dnsHalt := make(chan bool)
	dnsserver.Start(dnsHalt)
	log.Println("DNS Server started")

	webHalt := make(chan bool)
	webserver.Start(webHalt)
	log.Println("Web Server started")

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig

	fmt.Printf("Signal (%s) received, stopping\n", s)
	dnsHalt <- true
	webHalt <- true
}

package main

import (
	"github.com/billopark/iep.ee/internal/dnsserver"
	_ "github.com/billopark/iep.ee/internal/log"
	"github.com/billopark/iep.ee/internal/webserver"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	dnsHalt := make(chan bool)
	dnsserver.Start(dnsHalt)
	log.Infoln("DNS Server started")

	webHalt := make(chan bool)
	webserver.Start(webHalt)
	log.Infoln("Web Server started")

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	s := <-sig

	log.Infof("Signal (%s) received, stopping\n", s)
	dnsHalt <- true
	webHalt <- true
}

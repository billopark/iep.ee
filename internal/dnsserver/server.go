package dnsserver

import (
	"fmt"
	"github.com/billopark/iep.ee/config"
	"github.com/miekg/dns"
	"log"
)

func Start(halt chan bool) {
	dns.HandleFunc(config.Get().Domain, handler)
	server := &dns.Server{
		Addr: fmt.Sprintf("[::]:%d", config.Get().DnsPort),
		Net:  "udp",
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to setup the udp dnsserver: %s\n", err.Error())
		}
	}()

	go func() {
		<-halt
		if err := server.Shutdown(); err != nil {
			panic(err)
		} else {
			log.Println("DNS Server halted")
		}
	}()

	return
}

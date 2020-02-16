package record

import (
	"fmt"
	"github.com/billopark/iep.ee/config"
	"github.com/miekg/dns"
	"strconv"
	"sync"
	"time"
)

var serial string
var soaOnce sync.Once

func getSerial() string {
	soaOnce.Do(func() {
		serial = config.Get().SOA.Serial
		if _, err := strconv.Atoi(serial); serial == "" || err != nil {
			now := time.Now()
			serial = fmt.Sprintf("%04d%02d%2d00", now.Year(), now.Month(), now.Day())
		}
	})

	return serial
}

func BuildSOA() (dns.RR, error) {
	domain := config.Get().Domain
	soaConfig := config.Get().SOA
	mname := dns.Fqdn(soaConfig.MName)
	rname := dns.Fqdn(soaConfig.RName)
	soa := fmt.Sprintf("%s IN SOA %s %s %s 21600 7200 604800 3600", domain, mname, rname, getSerial())

	return dns.NewRR(soa)
}

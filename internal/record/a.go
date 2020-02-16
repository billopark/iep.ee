package record

import (
	"errors"
	"github.com/miekg/dns"
	"net"
)

func BuildA(query string, ip net.IP) (dns.RR, error) {
	if ip == nil {
		ip = buildIP(query)
	}
	if ip == nil {
		return nil, errors.New("No proper ip found.\n")
	}

	isV4 := ip.To4() != nil
	if isV4 {
		return &dns.A{
			Hdr: dns.RR_Header{
				Name:   query,
				Rrtype: dns.TypeA,
				Class:  dns.ClassINET,
				Ttl:    0,
			},
			A: ip.To4(),
		}, nil
	} else {
		return &dns.AAAA{
			Hdr: dns.RR_Header{
				Name:   query,
				Rrtype: dns.TypeAAAA,
				Class:  dns.ClassINET,
				Ttl:    0,
			},
			AAAA: ip.To4(),
		}, nil
	}
}

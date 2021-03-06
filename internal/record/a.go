package record

import (
	"errors"
	"github.com/miekg/dns"
	"net"
)

func BuildA(query string, ip net.IP) (dns.RR, error) {
	if ip == nil {
		ip = buildIPv4(query)
	}
	if ip == nil {
		return nil, errors.New("No proper ip found.\n")
	}

	var ipv4 net.IP

	if ipv4 = ip.To4(); ipv4 == nil {
		return nil, errors.New("No proper ip found.\n")
	}

	return &dns.A{
		Hdr: dns.RR_Header{
			Name:   query,
			Rrtype: dns.TypeA,
			Class:  dns.ClassINET,
			Ttl:    0,
		},
		A: ipv4,
	}, nil

}

package record

import (
	"errors"
	"github.com/miekg/dns"
	"net"
)

func BuildAAAA(query string, ip net.IP) (dns.RR, error) {
	if ip == nil {
		ip = buildIPv6(query)
	}
	if ip == nil {
		return nil, errors.New("No proper ip found.\n")
	}

	var ipv6 net.IP

	if ipv6 = ip.To16(); ipv6 == nil {
		return nil, errors.New("No proper ip found.\n")
	}

	return &dns.AAAA{
		Hdr: dns.RR_Header{
			Name:   query,
			Rrtype: dns.TypeAAAA,
			Class:  dns.ClassINET,
			Ttl:    0,
		},
		AAAA: ipv6,
	}, nil

}

package record

import (
	"github.com/billopark/iep.ee/config"
	"github.com/miekg/dns"
)

func BuildNS(query string) dns.RR {
	return &dns.NS{
		Hdr: dns.RR_Header{
			Name:   config.Get().Domain,
			Rrtype: dns.TypeNS,
			Class:  dns.ClassINET,
			Ttl:    0,
		},
		Ns: query,
	}
}

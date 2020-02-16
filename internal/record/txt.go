package record

import (
	"github.com/billopark/iep.ee/config"
	"github.com/miekg/dns"
)

func BuildTXT() dns.RR {
	return &dns.TXT{
		Hdr: dns.RR_Header{
			Name:   config.Get().Domain,
			Rrtype: dns.TypeTXT,
			Class:  dns.ClassINET,
			Ttl:    0,
		},
		Txt: []string{"Hello! This is IEPEE service. For more detail, please visit http://iep.ee"},
	}
}

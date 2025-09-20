package dnslookup

import (
	"context"
	"fmt"
	"net"
)

// DNSRecords guarda os resultados para vários tipos DNS
type DNSRecords struct {
	A     []string
	AAAA  []string
	MX    []string
	TXT   []string
	NS    []string
	CNAME []string
}

// Lookup faz a consulta DNS avançada para o domínio informado
func Lookup(ctx context.Context, domain string) (*DNSRecords, error) {
	records := &DNSRecords{}

	// lookup A
	aRecords, err := net.DefaultResolver.LookupHost(ctx, domain)
	if err == nil {
		records.A = aRecords
	}

	// lookup AAAA
	aaaaRecords, err := net.DefaultResolver.LookupIPAddr(ctx, domain)
	if err == nil {
		for _, ip := range aaaaRecords {
			if ip.IP.To4() == nil { // só IPv6
				records.AAAA = append(records.AAAA, ip.IP.String())
			}
		}
	}

	// lookup MX
	mxRecords, err := net.DefaultResolver.LookupMX(ctx, domain)
	if err == nil {
		for _, mx := range mxRecords {
			records.MX = append(records.MX, fmt.Sprintf("%s (Pref: %d)", mx.Host, mx.Pref))
		}
	}

	// lookup TXT
	txtRecords, err := net.DefaultResolver.LookupTXT(ctx, domain)
	if err == nil {
		records.TXT = txtRecords
	}

	// lookup NS
	nsRecords, err := net.DefaultResolver.LookupNS(ctx, domain)
	if err == nil {
		for _, ns := range nsRecords {
			records.NS = append(records.NS, ns.Host)
		}
	}

	// lookup CNAME
	cname, err := net.DefaultResolver.LookupCNAME(ctx, domain)
	if err == nil {
		records.CNAME = []string{cname}
	}

	return records, nil
}

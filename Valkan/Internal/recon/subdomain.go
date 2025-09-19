package recon

import (
	"context" // IMPORTAR para usar context.Background()
	"fmt"
	"net"
	"sync"
	"time"
)

// Resultado de uma tentativa de resolução de subdomínio
type SubdomainResult struct {
	Subdomain string
	IP        string
	Alive     bool
}

// Lista básica de subdomínios comuns
var commonSubdomains = []string{
	"www", "mail", "ftp", "webmail", "localhost",
	"cpanel", "blog", "dev", "test", "api", "shop",
	"staging", "portal", "admin", "vpn", "cdn",
	"forum", "ns1", "ns2", "smtp",
}

// FindSubdomains tenta resolver subdomínios comuns para um domínio
func FindSubdomains(domain string, timeout time.Duration, concurrency int) []SubdomainResult {
	var wg sync.WaitGroup
	var mu sync.Mutex
	results := []SubdomainResult{}

	sem := make(chan struct{}, concurrency)

	for _, sub := range commonSubdomains {
		wg.Add(1)
		sem <- struct{}{}
		go func(sub string) {
			defer wg.Done()
			defer func() { <-sem }()

			full := fmt.Sprintf("%s.%s", sub, domain)
			res := resolveSubdomain(full, timeout)
			if res.Alive {
				mu.Lock()
				results = append(results, res)
				mu.Unlock()
			}
		}(sub)
	}

	wg.Wait()
	return results
}

// resolveSubdomain tenta resolver o subdomínio e retorna resultado
func resolveSubdomain(subdomain string, timeout time.Duration) SubdomainResult {
	r := SubdomainResult{Subdomain: subdomain, Alive: false}

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{Timeout: timeout}
			return d.DialContext(ctx, network, address)
		},
	}

	ips, err := resolver.LookupHost(context.Background(), subdomain)
	if err != nil {
		return r
	}

	if len(ips) > 0 {
		r.IP = ips[0]
		r.Alive = true
	}
	return r
}

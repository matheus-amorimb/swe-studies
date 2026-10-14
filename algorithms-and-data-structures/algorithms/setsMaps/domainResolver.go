package main

import (
	"fmt"
)

func main() {
	r := DomainResolver{
		ipsToDomains:        map[string]map[string]struct{}{},
		domainsToSubDomains: map[string]map[string]struct{}{},
	}
	r.registerDomain("192.168.1.1", "example.com")
	r.registerDomain("192.168.1.1", "example.org")
	r.registerDomain("192.168.1.2", "domain.com")
	r.registerSubdomain("example.com", "a")
	r.registerSubdomain("example.com", "b")
	fmt.Println(r.hasSubdomain("192.168.1.1", "example.com", "a"))
	fmt.Println(r.hasSubdomain("192.168.1.1", "example.com", "c"))
	fmt.Println(r.hasSubdomain("127.0.0.1", "example.com", "a"))
	fmt.Println(r.hasSubdomain("192.168.1.1", "example.org", "a"))
}

type DomainResolver struct {
	ipsToDomains        map[string]map[string]struct{}
	domainsToSubDomains map[string]map[string]struct{}
}

func (dr *DomainResolver) registerDomain(ip, d string) {
	if _, ok := dr.ipsToDomains[ip]; !ok {
		dr.ipsToDomains[ip] = map[string]struct{}{}
	}
	dr.ipsToDomains[ip][d] = struct{}{}
}

func (dr *DomainResolver) registerSubdomain(d, sd string) {
	if _, ok := dr.domainsToSubDomains[d]; !ok {
		dr.domainsToSubDomains[d] = map[string]struct{}{}
	}
	dr.domainsToSubDomains[d][sd] = struct{}{}
}

func (dr *DomainResolver) hasSubdomain(ip, d, sd string) bool {
	if _, ok := dr.ipsToDomains[ip][d]; !ok {
		return false
	}

	_, ok := dr.domainsToSubDomains[d][sd]

	return ok
}

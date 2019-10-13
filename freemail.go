package freemail

import (
	"strings"
)

//go:generate go run gen.go
//go:generate go fmt .

// IsFree returns true if the email domain is a known free email provider.
func IsFree(domain string) bool {
	return find(domain, free)
}

// IsDisposable returns true if the email domain is known to be a disposable
// email provider.
func IsDisposable(domain string) bool {
	return find(domain, disposable)
}

func find(domain string, domains map[string]bool) bool {
	if domains[domain] {
		return true
	}
	parts := strings.Split(domain, ".")
	for i := 0; i < len(parts); i++ {
		test := strings.Join(parts[i:], ".")
		if domains[test] {
			return true
		}
	}
	return false
}

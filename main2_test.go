package main

import (
	"testing"
)

func TestMainInvalidCert(t *testing.T) {
	userDetails = userCreds{username: "admin", password: "spider~heLm3t/auto", realm: "admin/default-idp-1"}
	userDetailsOrg = userCreds{username: "chrisp", password: "Alligat0r/clips", realm: "provider/ldap-local"}
	server = "apimdev1075.hursley.ibm.com"
	// orgs = []string{"chrisp"}
	orgs = []string{"chrisp", "chrisp2", "chrisp3"}
	main()
}

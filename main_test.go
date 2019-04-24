package main

import (
	"testing"
)

// noMember     int
// noTLSProfile int
// org          *[]organization
// azs          *[]az -- Done

func TestMain(m *testing.M) {
	userDetails = userCreds{username: "admin", password: "spider~heLm3t/auto", realm: "admin/default-idp-1"}
	userDetailsOrg = userCreds{username: "chrisp", password: "Alligat0r/clips", realm: "provider/ldap-local"}
	server = "apim.lts.apicww.cloud"
	orgs = []string{"chrisp", "chrisp2", "chrisp3"}
	main()
}

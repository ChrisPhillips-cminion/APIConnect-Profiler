package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

type TestData struct {
	UserDetails    map[string]string
	UserDetailsOrg map[string]string
	Server         string
	Orgs           []string
	Output         string
}

var ConstTestData TestData

func TestMain(m *testing.M) {
	Log("---TestMain---")
	Log("Loading Test Environmental Data")
	jsonFile, _ := os.Open("env.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &ConstTestData)

	setUp()

	m.Run()
}

func setUp() {
	Log("\t---Setup---")

	userDetails = userCreds{username: ConstTestData.UserDetails["username"], password: ConstTestData.UserDetails["password"], realm: ConstTestData.UserDetails["realm"]}
	userDetailsOrg = userCreds{username: ConstTestData.UserDetailsOrg["username"], password: ConstTestData.UserDetailsOrg["password"], realm: ConstTestData.UserDetailsOrg["realm"]}
	server = ConstTestData.Server
	orgs = ConstTestData.Orgs
	output = ConstTestData.Output
	debug = false
}

func TestMainMethod(t *testing.T) {

	Log("---TestMainMethod---")
	setUp()
	output = "unset"
	main()
}

func TestValid(t *testing.T) {
	Log("---TestValid---")
	setUp()
	mainRunner()
	// t.Errorf("Abs(-1) = %d; want 1", 1)
}

func TestInvalidServer(t *testing.T) {
	Log("---TestInvalidServer---")
	setUp()
	server = "rubbish.local"
	mainRunner()
	// t.Errorf("Abs(-1) = %d; want 1", 1)
}

func TestNotValidOrg(t *testing.T) {
	// debug = true
	Log("---TestNotValidOrg---")
	setUp()
	orgs = []string{"chrisp", "rubbish", "chrisp3"}
	mainRunner()
}

func TestNotValidCredsOrg(t *testing.T) {
	Log("---TestNotValidCredsOrg---")
	setUp()
	userDetailsOrg = userCreds{username: "chris", password: "Alligat0r/clips", realm: "provider/ldap-local"}

	mainRunner()
}

func TestNotValidCredsCM(t *testing.T) {
	Log("---TestNotValidCredsCM---")
	setUp()
	userDetails = userCreds{username: "aasdasddmin", password: "spider~heLm3t/auto", realm: "admin/default-idp-1"}

	mainRunner()
}

func TestNotValidCredsForOneOrg(t *testing.T) {
	Log("---TestNotValidCredsForOneOrg---")
	setUp()
	orgs = []string{"ben", "chrisp2", "chrisp3"}

	mainRunner()
}
func TestYaml(t *testing.T) {

	Log("---TestYaml---")
	setUp()
	output = "yaml"

	mainRunner()
}
func TestTable(t *testing.T) {

	Log("---Testtable---")
	setUp()
	output = "table"

	mainRunner()
}
func TestVerbose(t *testing.T) {

	Log("---TestVerbose---")
	setUp()
	output = "verbose"

	mainRunner()
}
func TestOutputError(t *testing.T) {

	Log("---TestOutputError---")
	setUp()
	output = "ERROR"

	mainRunner()
}
func TestDebug(t *testing.T) {

	Log("---TestDebug---")
	setUp()
	debug = true

	mainRunner()
}

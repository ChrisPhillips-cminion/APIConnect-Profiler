package main

import (
	"encoding/json"
	"github.com/chrisphillips-cminion/apiprofile/pkg/model"
	"github.com/chrisphillips-cminion/apiprofile/pkg/utils"
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
	utils.Log("---TestMain---")
	utils.Log("Loading Test Environmental Data")
	jsonFile, _ := os.Open("env.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &ConstTestData)

	setUp()

	m.Run()
}

func setUp() {
	utils.Log("\t---Setup---")

	utils.Vars.UserDetails = model.UserCreds{Username: ConstTestData.UserDetails["Username"], Password: ConstTestData.UserDetails["Password"], Realm: ConstTestData.UserDetails["Realm"]}
	utils.Vars.UserDetailsOrg = model.UserCreds{Username: ConstTestData.UserDetailsOrg["Username"], Password: ConstTestData.UserDetailsOrg["Password"], Realm: ConstTestData.UserDetailsOrg["Realm"]}
	utils.Vars.Server = ConstTestData.Server
	utils.Vars.Orgs = ConstTestData.Orgs
	utils.Vars.Output = ConstTestData.Output
	utils.Vars.Debug = false
}

func TestMainMethod(t *testing.T) {

	utils.Log("---TestMainMethod---")
	setUp()
	utils.Vars.Output = "unset"
	main()
}

func TestValid(t *testing.T) {
	utils.Log("---TestValid---")
	setUp()
	mainRunner()
	// t.Errorf("Abs(-1) = %d; want 1", 1)
}

func TestInvalidServer(t *testing.T) {
	utils.Log("---TestInvalidServer---")
	setUp()
	utils.Vars.Server = "rubbish.local"
	mainRunner()
	// t.Errorf("Abs(-1) = %d; want 1", 1)
}

func TestNotValidOrg(t *testing.T) {
	// debug = true
	utils.Log("---TestNotValidOrg---")
	setUp()
	utils.Vars.Orgs = []string{"chrisp", "rubbish", "chrisp3"}
	mainRunner()
}

func TestNotValidCredsOrg(t *testing.T) {
	utils.Log("---TestNotValidCredsOrg---")
	setUp()
	utils.Vars.UserDetailsOrg = model.UserCreds{Username: "chris", Password: "Alligat0r/clips", Realm: "provider/ldap-local"}

	mainRunner()
}

func TestNotValidCredsCM(t *testing.T) {
	utils.Log("---TestNotValidCredsCM---")
	setUp()
	utils.Vars.UserDetails = model.UserCreds{Username: "aasdasddmin", Password: "spider~heLm3t/auto", Realm: "admin/default-idp-1"}

	mainRunner()
}

func TestNotValidCredsForOneOrg(t *testing.T) {
	utils.Log("---TestNotValidCredsForOneOrg---")
	setUp()
	utils.Vars.Orgs = []string{"ben", "chrisp2", "chrisp3"}

	mainRunner()
}
func TestYaml(t *testing.T) {

	utils.Log("---TestYaml---")
	setUp()
	utils.Vars.Output = "yaml"

	mainRunner()
}
func TestTable(t *testing.T) {

	utils.Log("---Testtable---")
	setUp()
	utils.Vars.Output = "table"

	mainRunner()
}
func TestVerbose(t *testing.T) {

	utils.Log("---TestVerbose---")
	setUp()
	utils.Vars.Output = "verbose"

	mainRunner()
}
func TestOutputError(t *testing.T) {

	utils.Log("---TestOutputError---")
	setUp()
	utils.Vars.Output = "ERROR"

	mainRunner()
}
func TestDebug(t *testing.T) {

	utils.Log("---TestDebug---")
	setUp()
	utils.Vars.Debug = true

	mainRunner()
}

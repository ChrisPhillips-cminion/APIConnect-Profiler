package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

var server = "unset"

var token = "unset"
var userDetails = userCreds{}

//This is here to supprt the test
var userDetailsOrg = userCreds{}
var orgs = make([]string, 3)
var output = "unset"
var debug = false

func main() {
	TraceEnter("main")
	// output := "unset"
	outputTypePtr := flag.String("output", "table", "Dictates the output type for the script. Value must be one of [ table | json | yaml | verbose ] ")
	debugPtr := flag.Bool("debug", false, "Enable trace for this appication")
	serverPtr := flag.String("server", "unset", "APIConnect Cloud endpoint, if this is not set it is prompted")
	realmPtr := flag.String("CMrealm", "", "Realm for logging into the Cloud Manager Endpoint, if this is not set it is prompted")
	userPtr := flag.String("CMuser", "", "APIConnect User for logging into the Cloud Manager Endpoint, if this is not set it is prompted")
	passwordPtr := flag.String("CMpassword", "", "APIConnect Password for logging into the Cloud Manager Endpoint, if this is not set it is prompted")

	passwordManagerPtr := flag.String("APIMpassword", "", "APIConnect Password for logging into the API Manager Endpoint, if this is not set it is prompted")
	userManagerPtr := flag.String("APIMuser", "", "APIConnect User for logging into the API Manager Endpoint, if this is not set it is prompted")
	realmManagerPtr := flag.String("APIMrealm", "", "Realm for logging into the API Manager Endpoint, if this is not set it is prompted")
	orgPtr := flag.String("APIMorg", "", "Organiztion List to investigate. Please multiple orgs in csv, e.g. dev,test,chrisp,marketting ")
	flag.Parse()

	if server == "unset" {
		server = *serverPtr
	}
	if !debug {
		debug = *debugPtr
	}
	if *realmPtr != "" {
		userDetails = userCreds{
			username: *userPtr,
			password: *passwordPtr,
			realm:    *realmPtr}
	}

	if output == "unset" {
		output = *outputTypePtr
	}

	if *orgPtr != "" {
		orgs = strings.Split(*orgPtr, ",")
	}
	if *realmManagerPtr != "" {
		userDetailsOrg = userCreds{
			username: *userManagerPtr,
			password: *passwordManagerPtr,
			realm:    *realmManagerPtr}
	}

	mainRunner()
}

//This is split out because the test tool can not cope with the flags being set multiple times
func mainRunner() {
	server = PromptServer()
	creds := PromptCredentials("Admin")
	Log("Logging in")
	err := errors.New("Not an error")
	token, err = login(creds)
	if err != nil {
		Log(server)
		Log(err.Error())
	} else {
		Log("Gathering Data, this may take some time")
		data, err := GetTopLevel()
		if err != nil {
			Log(err.Error())
		} else {
			switch output {
			case "yaml":
				Yamlify(data)
				break
			case "json":
				JSONify(data)
				break
			case "table":
				Tablify(data)
				break
			case "verbose":
				printData(data)
				break
			default:
				fmt.Printf("Unexpected value '%s' please provide one of [ table | json | yaml | verbose ] .\n", output)
			}
		}
	}
	//
	TraceExit("main")
}

package main
/*
      Licensed Materials - Property of IBM
      © IBM Corp. 2019
*/
import (
	"errors"
	"flag"
	"fmt"
	"github.com/chrisphillips-cminion/apiprofile/pkg/model"
	"github.com/chrisphillips-cminion/apiprofile/pkg/utils"
	"strings"
)


func main() {
	utils.TraceEnter("main")
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

	if utils.Vars.Server == "" {
		utils.Vars.Server = *serverPtr
	}
	if !utils.Vars.Debug {
		utils.Vars.Debug = *debugPtr
	}
	utils.Vars.UserDetails = model.UserCreds{}
	if *realmPtr != "" {
		utils.Vars.UserDetails.Realm = *realmPtr
	}
	if *userPtr != "" {
		utils.Vars.UserDetails.Username = *userPtr
	}
	if *passwordPtr != "" {
		utils.Vars.UserDetails.Password = *passwordPtr
	}

	if utils.Vars.Output == "" {
		utils.Vars.Output = *outputTypePtr
	}

	if *orgPtr != "" {
		utils.Vars.Orgs = strings.Split(*orgPtr, ",")
	}
	utils.Vars.UserDetailsOrg = model.UserCreds{}
	if *realmManagerPtr != "" {
		utils.Vars.UserDetailsOrg.Realm = *realmManagerPtr
	}
	if *userManagerPtr != "" {
		utils.Vars.UserDetailsOrg.Username = *userManagerPtr
	}
	if *passwordManagerPtr != "" {
		utils.Vars.UserDetailsOrg.Password = *passwordManagerPtr
	}

	mainRunner()
}


func mainRunner() {
	utils.Vars.Server = utils.PromptServer()
	creds := utils.PromptCredentials("Admin", "admin")
	utils.Log("Logging in")
	err := errors.New("Not an error")
	utils.Vars.Token, err = utils.Login(creds)
	if err != nil {
		utils.Log(utils.Vars.Server)
		utils.Log(err.Error())
	} else {
		utils.Log("Gathering Data, this may take some time")
		data, err := utils.GetTopLevel()
		if err != nil {
			utils.Log(err.Error())
		} else {
			switch utils.Vars.Output {
			case "yaml":
				utils.Yamlify(data)
				break
			case "json":
				utils.JSONify(data)
				break
			case "table":
				utils.Tablify(data)
				break
			case "verbose":
				utils.PrintData(data)
				break
			default:
				fmt.Printf("Unexpected value '%s' please provide one of [ table | json | yaml | verbose ] .\n", utils.Vars.Output)
			}
		}
	}
	//
	utils.TraceExit("main")
}

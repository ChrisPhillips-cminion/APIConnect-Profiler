package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
)

var server = "unset"

var token = "unset"
var userDetails = userCreds{}

//This is here to supprt the test
var userDetailsOrg = userCreds{}
var orgs = make([]string, 3)

// noMember     int
// noTLSProfile int
// org          *[]organization
// azs          *[]az -- Done
func main() {
	TraceEnter("main")

	server = PromptServer()
	creds := PromptCredentials("Admin")
	Log("Logging in")
	token = login(creds)
	Log("Gathering Data, this may take some time")
	Tablify(GetTopLevel())
	TraceExit("main")
}

//
// noPortal     int
// noAnalytics  int
// noV5cgateway int
// noApigateway int
func printData(data topLevel) {
	TraceEnter("printData")
	Trace("************************************")
	Trace("")
	Trace("")
	br()
	log.Printf("Cloud Name = '%v'", data.name)
	br()
	log.Printf(" Cloud Manager Details (-1 means an issue retreiving the data)")
	log.Printf(" Number of Orgs                            -  %v ", data.noOrg)
	log.Printf(" Number of Members                         -  %v ", data.noMembers)
	log.Printf(" Number of Member Invites                  -  %v ", data.noMemberInvites)
	log.Printf(" Number of Oauth Providers                 -  %v ", data.noOauthP)
	log.Printf(" Number of User Registries                 -  %v ", data.noUserReg)
	log.Printf(" Number of Mail Servers                    -  %v ", data.noMailServers)
	br()
	log.Printf(" Number of Availability Zones %v", len(*data.azs))
	br()
	for i, v := range *data.azs {
		log.Printf("\t Availability Zone %v - %v", i, v.name)
		br()
		log.Printf("\t\t Number of V5C GW              -  %v ", v.noV5cgateway)
		log.Printf("\t\t Number of API GW              -  %v ", v.noApigateway)
		log.Printf("\t\t Number of Portal              -  %v ", v.noPortal)
		log.Printf("\t\t Number of Analytics           -  %v ", v.noAnalytics)
		br()
	}
	br()
	log.Printf("Number of Provider Organizations to investigate %v", len(*data.org))
	br()
	for i, v := range *data.org {
		log.Printf("\t Provider Organization %v - %v", i, v.name)
		br()
		log.Printf("\t\t POrg %v - %v ", i, v.name)
		log.Printf("\t\t Number of Members             -  %v ", v.noMembers)
		log.Printf("\t\t Number of Member Invitations  -  %v ", v.noMemberInvites)
		log.Printf("\t\t Number of DraftAPI            -  %v ", v.noDraftAPI)
		log.Printf("\t\t Average API Size (bytes)      -  %v ", v.avgAPISize)
		log.Printf("\t\t Max API Size (bytes)          -  %v ", v.maxAPISize)
		log.Printf("\t\t Number of Draft Products      -  %v ", v.noDraftProduct)
		log.Printf("\t\t Number of TLS Profiles        -  %v ", v.noTLSProfile)
		log.Printf("\t\t Number of OAuth Proivders     -  %v ", v.noOAuthProvider)
		log.Printf("\t\t Number of User Registries     -  %v ", v.userRegistries)
		log.Printf("\t\t Number of Key Stores          -  %v ", v.noKeyStore)
		log.Printf("\t\t Number of Trust Stores        -  %v ", v.noTrustStore)
		br()
		log.Printf("\t\t Number of catalogs            -  %v ", len(*v.catalog))
		br()
		for ci, cv := range *v.catalog {
			log.Printf("\t\t\t Catalog %v - %v", ci, cv.name)
			br()

			log.Printf("\t\t\t\t Number of Members                -  %v ", cv.noMember)
			log.Printf("\t\t\t\t Number of Member Invitations     -  %v ", cv.noMemberInvites)
			log.Printf("\t\t\t\t Number of APIs                   -  %v ", cv.noAPI)
			log.Printf("\t\t\t\t Number of Products               -  %v ", cv.noProduct)
			log.Printf("\t\t\t\t Average API Size                 -  %v ", cv.avgAPISize)
			log.Printf("\t\t\t\t Max API Size                     -  %v ", cv.maxAPISize)
			log.Printf("\t\t\t\t Number of Consumer Orgs          -  %v ", cv.noConsumerOrg)
			log.Printf("\t\t\t\t Portal Enabled                   -  %v ", cv.portal)
			log.Printf("\t\t\t\t Number of TLS Profiles           -  %v ", cv.noTLSProfile)
			log.Printf("\t\t\t\t Number of OAuth Providers        -  %v ", cv.noOAuthProvider)
			log.Printf("\t\t\t\t Number of User Registries        -  %v ", cv.userRegistries)
			log.Printf("\t\t\t\t Number of Spaces                 -  %v ", cv.noSpace)
		}
		br()
	}
	Trace("")
	Trace("")
	Trace("************************************")
	TraceExit("printData")
}

func br() {
	log.Printf("--------------------------------------------------------------------------------------------------------")
}
func Tablify(data topLevel) {
	TraceEnter("Tablify")
	Trace("************************************")
	Trace("")
	Trace("")
	br()
	log.Printf("Cloud Name = '%v'", data.name)
	br()
	log.Printf(" Cloud Manager Details (-1 means an issue retreiving the data or 0 found.)")
	log.Printf(" Number of Orgs                            -  %v ", data.noOrg)
	log.Printf(" Number of Members                         -  %v ", data.noMembers)
	log.Printf(" Number of Member Invites                  -  %v ", data.noMemberInvites)
	log.Printf(" Number of Oauth Providers                 -  %v ", data.noOauthP)
	log.Printf(" Number of User Registries                 -  %v ", data.noUserReg)
	log.Printf(" Number of Mail Servers                    -  %v ", data.noMailServers)
	br()
	log.Printf(" Number of Availability Zones %v", len(*data.azs))
	br()
	//
	title := []string{"Name", "Number of V5C GWs", "Number of API GWs", "Number of Portals", "Number of Analytics"}
	content := make([][]string, len(*data.azs))
	for i, v := range *data.azs {
		content[i] = []string{
			fmt.Sprintf("%v", v.name),
			fmt.Sprintf("%v", v.noV5cgateway),
			fmt.Sprintf("%v", v.noApigateway),
			fmt.Sprintf("%v", v.noPortal),
			fmt.Sprintf("%v", v.noAnalytics),
		}
	}
	RenderTable(title, content)

	br()
	log.Printf("Number of Provider Organizations to investigate %v", len(*data.org))
	br()

	title = []string{"Name", "Members", "MemberInvites", "DraftAPIs", "Avg API Size", "Max API Size", "DraftProducts", "TLS Profiles", "OAuth Proivders", "User Registries", "KeyStores", "TrustStores", "Catalogs"}
	content = make([][]string, len(*data.org))
	catCount := 0
	for i, v := range *data.org {
		content[i] = []string{
			fmt.Sprintf("%v", v.name),
			fmt.Sprintf("%v", v.noMembers),
			fmt.Sprintf("%v", v.noMemberInvites),
			fmt.Sprintf("%v", v.noDraftAPI),
			fmt.Sprintf("%v", v.avgAPISize),
			fmt.Sprintf("%v", v.maxAPISize),
			fmt.Sprintf("%v", v.noDraftProduct),
			fmt.Sprintf("%v", v.noTLSProfile),
			fmt.Sprintf("%v", v.noOAuthProvider),
			fmt.Sprintf("%v", v.userRegistries),
			fmt.Sprintf("%v", v.noKeyStore),
			fmt.Sprintf("%v", v.noTrustStore),
			fmt.Sprintf("%v", len(*v.catalog)),
		}
		catCount = catCount + len(*v.catalog)
	}
	RenderTable(title, content)
	br()
	title = []string{"org", "Name", "Members", "MemberInvites", "APIs", "Avg API Size", "Max API Size", "Products", "ConsumerOrgs", "Portal", "TLSProfiles", "OAuthProivders", "UserRegistries", "Spaces", "Apps"}
	content = make([][]string, catCount)
	cnt := 0
	log.Printf("\t\t Number of catalogs investigated - %v  ", catCount)
	br()
	for _, v := range *data.org {

		for _, cv := range *v.catalog {
			content[cnt] = []string{
				fmt.Sprintf("%v", v.name),
				fmt.Sprintf("%v", cv.name),
				fmt.Sprintf("%v", cv.noMember),
				fmt.Sprintf("%v", cv.noMemberInvites),
				fmt.Sprintf("%v", cv.noAPI),
				fmt.Sprintf("%v", cv.avgAPISize),
				fmt.Sprintf("%v", cv.maxAPISize),
				fmt.Sprintf("%v", cv.noProduct),
				fmt.Sprintf("%v", cv.noConsumerOrg),
				fmt.Sprintf("%v", cv.portal),
				fmt.Sprintf("%v", cv.noTLSProfile),
				fmt.Sprintf("%v", cv.noOAuthProvider),
				fmt.Sprintf("%v", cv.userRegistries),
				fmt.Sprintf("%v", cv.noSpace),
				fmt.Sprintf("%v", cv.applications)}
		}
		cnt++
	}
	RenderTable(title, content)
}

func RenderTable(title []string, obj [][]string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(title)
	for _, v := range obj {
		content := make([]string, len(title))
		for i, ev := range v {
			content[i] = ev
		}
		table.Append(content)
	}
	table.Render()
}

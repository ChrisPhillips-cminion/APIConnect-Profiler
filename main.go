package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

var server = "unset"

var token = "unset"
var userDetails = userCreds{}

//This is here to supprt the test
var userDetailsOrg = userCreds{}
var orgs = make([]string, 3)

func main() {
	TraceEnter("main")

	server = PromptServer()
	creds := PromptCredentials("Admin")
	Log("Logging in")
	token = login(creds)
	Log("Gathering Data, this may take some time")
	Tablify(GetTopLevel())
	// printData(GetTopLevel())
	TraceExit("main")
}

func printData(data topLevel) {
	TraceEnter("printData")
	Trace("************************************")
	Trace("")
	Trace("")
	br()
	fmt.Printf("\nCloud Name = '%v'", data.name)
	br()
	fmt.Printf("\n Cloud Manager Details (-1 means an issue retreiving the data)")
	fmt.Printf("\n Number of Orgs                            -  %v ", data.noOrg)
	fmt.Printf("\n Number of Members                         -  %v ", data.noMembers)
	fmt.Printf("\n Number of Member Invites                  -  %v ", data.noMemberInvites)
	fmt.Printf("\n Number of Oauth Providers                 -  %v ", data.noOauthP)
	fmt.Printf("\n Number of User Registries                 -  %v ", data.noUserReg)
	fmt.Printf("\n Number of Mail Servers                    -  %v ", data.noMailServers)
	br()
	fmt.Printf("\n Number of Availability Zones %v", len(*data.azs))
	br()
	for i, v := range *data.azs {
		fmt.Printf("\n\t Availability Zone %v - %v", i, v.name)
		br()
		fmt.Printf("\n\t\t Number of V5C GW              -  %v ", v.noV5cgateway)
		fmt.Printf("\n\t\t Number of API GW              -  %v ", v.noApigateway)
		fmt.Printf("\n\t\t Number of Portal              -  %v ", v.noPortal)
		fmt.Printf("\n\t\t Number of Analytics           -  %v ", v.noAnalytics)
		br()
	}
	br()
	fmt.Printf("\nNumber of Provider Organizations to investigate %v", len(*data.org))
	br()
	for i, v := range *data.org {
		fmt.Printf("\n\t Provider Organization %v - %v", i, v.name)
		br()
		fmt.Printf("\n\t\t POrg %v - %v ", i, v.name)
		fmt.Printf("\n\t\t Number of Members             -  %v ", v.noMembers)
		fmt.Printf("\n\t\t Number of Member Invitations  -  %v ", v.noMemberInvites)
		fmt.Printf("\n\t\t Number of DraftAPI            -  %v ", v.noDraftAPI)
		fmt.Printf("\n\t\t Average API Size (bytes)      -  %v ", v.avgAPISize)
		fmt.Printf("\n\t\t Max API Size (bytes)          -  %v ", v.maxAPISize)
		fmt.Printf("\n\t\t Number of Draft Products      -  %v ", v.noDraftProduct)
		fmt.Printf("\n\t\t Number of TLS Profiles        -  %v ", v.noTLSProfile)
		fmt.Printf("\n\t\t Number of OAuth PROIVDERS     -  %v ", v.noOAuthProvider)
		fmt.Printf("\n\t\t Number of User Registries     -  %v ", v.userRegistries)
		fmt.Printf("\n\t\t Number of Key Stores          -  %v ", v.noKeyStore)
		fmt.Printf("\n\t\t Number of Trust Stores        -  %v ", v.noTrustStore)
		br()
		fmt.Printf("\n\t\t Number of catalogs            -  %v ", len(*v.catalog))
		br()
		for ci, cv := range *v.catalog {
			fmt.Printf("\n\t\t\t Catalog %v - %v", ci, cv.name)
			br()

			fmt.Printf("\n\t\t\t\t Number of Members                -  %v ", cv.noMember)
			fmt.Printf("\n\t\t\t\t Number of Member Invitations     -  %v ", cv.noMemberInvites)
			fmt.Printf("\n\t\t\t\t Number of APIs                   -  %v ", cv.noAPI)
			fmt.Printf("\n\t\t\t\t Number of Products               -  %v ", cv.noProduct)
			fmt.Printf("\n\t\t\t\t Average API Size                 -  %v ", cv.avgAPISize)
			fmt.Printf("\n\t\t\t\t Max API Size                     -  %v ", cv.maxAPISize)
			fmt.Printf("\n\t\t\t\t Number of Consumer Orgs          -  %v ", cv.noConsumerOrg)
			fmt.Printf("\n\t\t\t\t Portal Enabled                   -  %v ", cv.portal)
			fmt.Printf("\n\t\t\t\t Number of TLS Profiles           -  %v ", cv.noTLSProfile)
			fmt.Printf("\n\t\t\t\t Number of OAuth Providers        -  %v ", cv.noOAuthProvider)
			fmt.Printf("\n\t\t\t\t Number of User Registries        -  %v ", cv.userRegistries)
			fmt.Printf("\n\t\t\t\t Number of Spaces                 -  %v ", cv.noSpace)
			fmt.Printf("\n\t\t\t\t Number of Subscriptions          -  %v ", cv.subscriptions)
		}
		br()
	}
	Trace("")
	Trace("")
	Trace("************************************")
	TraceExit("printData")
}

func br() {
	fmt.Printf("\n--------------------------------------------------------------------------------------------------------\n")
}
func Tablify(data topLevel) {
	TraceEnter("Tablify")
	Trace("************************************")
	Trace("")
	Trace("")
	br()
	fmt.Printf("\n# Cloud Name = '%v'", data.name)
	fmt.Printf("\n-1 means an issue retreiving the data\n\n")
	br()
	topTable := make([][]string, 2)
	fmt.Printf("\n Cloud Manager Details\n\n")
	topTable[0] = []string{"Number of Orgs", "Number of Members", "Number of Member Invites", "Number of Oauth Providers", "Number of User Registries", "Number of Mail Servers"}
	topTable[1] = []string{
		fmt.Sprintf("%v", data.noOrg),
		fmt.Sprintf("%v", data.noMembers),
		fmt.Sprintf("%v", data.noMemberInvites),
		fmt.Sprintf("%v", data.noOauthP),
		fmt.Sprintf("%v", data.noUserReg),
		fmt.Sprintf("%v", data.noMailServers)}

	topTable = Transpose(topTable)
	topTableTitle := topTable[0]
	topTable = append(topTable[:0], topTable[1:]...)

	RenderTable(topTableTitle, topTable)

	br()
	fmt.Printf("\n ## Number of Availability Zones %v", len(*data.azs))
	br()
	//
	title := []string{"Name", "Number of V5C GWs", "Number of API GWs", "Number of Portals", "Number of Analytics"}
	content := make([][]string, len(*data.azs)+1)
	content[0] = title
	for i, v := range *data.azs {
		//ignore the first line as that will be the title

		content[i+1] = []string{
			fmt.Sprintf("%v", v.name),
			fmt.Sprintf("%v", v.noV5cgateway),
			fmt.Sprintf("%v", v.noApigateway),
			fmt.Sprintf("%v", v.noPortal),
			fmt.Sprintf("%v", v.noAnalytics),
		}
	}
	content[0] = title
	content = Transpose(content)
	title = content[0]

	RenderTable(title, content)

	br()
	fmt.Printf("\n## Number of Provider Organizations to investigate %v", len(*data.org))
	br()

	title = []string{"Name", "Members", "MemberInvites", "DraftAPIs", "Avg API Size", "Max API Size", "DraftProducts", "TLS Profiles", "OAuth PROIVDERS", "User Registries", "KeyStores", "TrustStores", "Catalogs"}
	content = make([][]string, len(*data.org)+1)
	catCount := 0
	for i, v := range *data.org {
		content[i+1] = []string{
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
	content[0] = title
	content = Transpose(content)
	title = content[0]

	RenderTable(title, content)
	br()
	title = []string{"org",
		"Catalog Name",
		"Members",
		"MemberInvites",
		"APIs",
		"Avg API Size",
		"Max API Size",
		"Products",
		"ConsumerOrgs",
		"Portal",
		"TLSProfiles",
		"OAuthPros",
		"UserRegs",
		"Spaces",
		"Apps",
		"Subscriptions",
		"Webhooks"}
	content = make([][]string, catCount+1)
	cnt := 1
	fmt.Printf("\n## \t\t Number of catalogs investigated - %v  ", catCount)
	br()
	webhookSize := 0
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
				fmt.Sprintf("%v", cv.applications),
				fmt.Sprintf("%v", cv.subscriptions),
				fmt.Sprintf("%v", len(*cv.webhooks))}
			webhookSize = webhookSize + len(*cv.webhooks)
			cnt++
		}

	}
	content[0] = title
	content = Transpose(content)
	title = content[1]

	RenderTable(title, content)

	br()
	title = []string{
		"WebHook Id",
		"Organization",
		"Catalog",
		"State",
		"Level",
		"Title",
		"Created At",
		"Updated At"}
	content = make([][]string, webhookSize+1)
	cnt = 1
	fmt.Printf("\n## \t\t Number of Webhooks investigated - %v  ", webhookSize)
	br()

	for _, v := range *data.org {

		for _, cv := range *v.catalog {
			for _, wv := range *cv.webhooks {
				content[cnt] = []string{
					// organizationName string
					// catalogName      string
					// organization     string
					// catalog          string
					// state            string
					// level            string
					// title            string
					// created_at       string
					// updated_at       string
					fmt.Sprintf("%v", wv.webhookId),
					fmt.Sprintf("%v\n%v", wv.organizationName, wv.organization),
					fmt.Sprintf("%v\n%v", wv.catalogName, wv.catalog),
					fmt.Sprintf("%v", wv.state),
					fmt.Sprintf("%v", wv.level),
					fmt.Sprintf("%v", wv.title),
					fmt.Sprintf("%v", wv.created_at),
					fmt.Sprintf("%v", wv.updated_at),
				}
				cnt++
			}
		}

	}
	// content[0] = title
	// content = Transpose(content)
	// title = content[0]

	RenderTable(title, content)

}
func Transpose(a [][]string) [][]string {
	x := len(a)
	y := len(a[0])

	b := make([][]string, y)
	for i := 0; i < y; i++ {
		b[i] = make([]string, x)
		for j := 0; j < x; j++ {
			b[i][j] = a[j][i]
		}
	}
	return b
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
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.Render()
}

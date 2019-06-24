package utils
/*
      Licensed Materials - Property of IBM
      © IBM Corp. 2019
*/
import (
	"fmt"
	"github.com/chrisphillips-cminion/apiprofile/pkg/model"
	"github.com/olekukonko/tablewriter"
	"os"
)

func Tablify(data model.TopLevel) {
	TraceEnter("Tablify")
	Trace("************************************")
	Trace("")
	Trace("")
	br()
	fmt.Printf("\n# Cloud Name = '%v'", data.Name)
	fmt.Printf("\n-1 means an issue retreiving the data\n\n")
	br()
	topTable := make([][]string, 2)
	fmt.Printf("\n Cloud Manager Details\n\n")
	topTable[0] = []string{"Number of Orgs", "Number of Members", "Number of Member Invites", "Number of Oauth Providers", "Number of User Registries", "Number of Mail Servers"}
	topTable[1] = []string{
		fmt.Sprintf("%v", data.NoOrg),
		fmt.Sprintf("%v", data.NoMembers),
		fmt.Sprintf("%v", data.NoMemberInvites),
		fmt.Sprintf("%v", data.NoOauthP),
		fmt.Sprintf("%v", data.NoUserReg),
		fmt.Sprintf("%v", data.NoMailServers),
	}

	topTable = Transpose(topTable)
	topTableTitle := topTable[0]
	topTable = append(topTable[:0], topTable[1:]...)

	RenderTable(topTableTitle, topTable)

	br()
	fmt.Printf("\n ## Number of Availability Zones %v", len(*data.Azs))
	br()
	//
	title := []string{"Name", "Number of V5C GWs", "Number of API GWs", "Number of Portals", "Number of Analytics"}
	content := make([][]string, len(*data.Azs)+1)
	// content[0] = title
	for i, v := range *data.Azs {
		//ignore the first line as that will be the title

		content[i] = []string{
			fmt.Sprintf("%v", v.Name),
			fmt.Sprintf("%v", v.NoV5cgateway),
			fmt.Sprintf("%v", v.NoApigateway),
			fmt.Sprintf("%v", v.NoPortal),
			fmt.Sprintf("%v", v.NoAnalytics),
		}
	}
	// content[0] = title
	// content = Transpose(content)
	// title = content[0]

	RenderTable(title, content)

	br()
	fmt.Printf("\n## Number of Provider Organizations to investigate %v", len(*data.Org))
	br()

	title = []string{"Name", " ", "Members", "MemberInvites", "DraftAPIs", "Avg API Size", "Max API Size", "DraftProducts", "TLS Profiles", "OAuth PROIVDERS", "User Registries", "KeyStores", "TrustStores", "Catalogs"}
	content = make([][]string, len(*data.Org)+1)
	catCount := 0

	for i, v := range *data.Org {

		content[i] = []string{
			fmt.Sprintf("%v", v.Name),
			fmt.Sprintf("%v", v.Error),
			fmt.Sprintf("%v", v.NoMembers),
			fmt.Sprintf("%v", v.NoMemberInvites),
			fmt.Sprintf("%v", v.NoDraftAPI),
			fmt.Sprintf("%v", v.AvgAPISize),
			fmt.Sprintf("%v", v.MaxAPISize),
			fmt.Sprintf("%v", v.NoDraftProduct),
			fmt.Sprintf("%v", v.NoTLSProfile),
			fmt.Sprintf("%v", v.NoOAuthProvider),
			fmt.Sprintf("%v", v.UserRegistries),
			fmt.Sprintf("%v", v.NoKeyStore),
			fmt.Sprintf("%v", v.NoTrustStore),
			fmt.Sprintf("%v", len(*v.Catalog)),
		}
		catCount = catCount + len(*v.Catalog)
	}
	// content[0] = title
	// content = Transpose(content)
	// title = content[0]

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
	for _, v := range *data.Org {

		for _, cv := range *v.Catalog {
			content[cnt] = []string{
				fmt.Sprintf("%v", v.Name),
				fmt.Sprintf("%v", cv.Name),
				fmt.Sprintf("%v", cv.NoMember),
				fmt.Sprintf("%v", cv.NoMemberInvites),
				fmt.Sprintf("%v", cv.NoAPI),
				fmt.Sprintf("%v", cv.AvgAPISize),
				fmt.Sprintf("%v", cv.MaxAPISize),
				fmt.Sprintf("%v", cv.NoProduct),
				fmt.Sprintf("%v", cv.NoConsumerOrg),
				fmt.Sprintf("%v", cv.Portal),
				fmt.Sprintf("%v", cv.NoTLSProfile),
				fmt.Sprintf("%v", cv.NoOAuthProvider),
				fmt.Sprintf("%v", cv.UserRegistries),
				fmt.Sprintf("%v", cv.NoSpace),
				fmt.Sprintf("%v", cv.Applications),
				fmt.Sprintf("%v", cv.Subscriptions),
				fmt.Sprintf("%v", len(*cv.Webhooks))}
			webhookSize = webhookSize + len(*cv.Webhooks)
			cnt++
		}

	}
	// content[0] = title
	// content = Transpose(content)
	// title = content[1]

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

	for _, v := range *data.Org {

		for _, cv := range *v.Catalog {
			for _, wv := range *cv.Webhooks {
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
					fmt.Sprintf("%v", wv.WebhookId),
					fmt.Sprintf("%v (%v)", wv.OrganizationName, wv.Organization),
					fmt.Sprintf("%v (%v)", wv.CatalogName, wv.Catalog),
					fmt.Sprintf("%v", wv.State),
					fmt.Sprintf("%v", wv.Level),
					fmt.Sprintf("%v", wv.Title),
					fmt.Sprintf("%v", wv.Created_at),
					fmt.Sprintf("%v", wv.Updated_at),
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

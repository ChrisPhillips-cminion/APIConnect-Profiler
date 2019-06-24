package utils
/*
      Licensed Materials - Property of IBM
      © IBM Corp. 2019
*/
import (
	"fmt"
	"github.com/chrisphillips-cminion/apiprofile/pkg/model"
)

func PrintData(data model.TopLevel) {
	TraceEnter("printData")
	Trace("************************************")
	Trace("")
	Trace("")
	br()
	fmt.Printf("\n# Cloud Name = '%v'", data.Name)
	br()
	fmt.Printf("\n Cloud Manager Details (-1 means an issue retreiving the data)")
	fmt.Printf("\n Number of Orgs                            -  %v ", data.NoOrg)
	fmt.Printf("\n Number of Members                         -  %v ", data.NoMembers)
	fmt.Printf("\n Number of Member Invites                  -  %v ", data.NoMemberInvites)
	fmt.Printf("\n Number of Oauth Providers                 -  %v ", data.NoOauthP)
	fmt.Printf("\n Number of User Registries                 -  %v ", data.NoUserReg)
	fmt.Printf("\n Number of Mail Servers                    -  %v ", data.NoMailServers)
	br()
	fmt.Printf("\n Number of Availability Zones %v", len(*data.Azs))
	br()
	for i, v := range *data.Azs {
		fmt.Printf("\n\t ## Availability Zone %v - %v", i, v.Name)
		br()
		fmt.Printf("\n\t\t Number of V5C GW              -  %v ", v.NoV5cgateway)
		fmt.Printf("\n\t\t Number of API GW              -  %v ", v.NoApigateway)
		fmt.Printf("\n\t\t Number of Portal              -  %v ", v.NoPortal)
		fmt.Printf("\n\t\t Number of Analytics           -  %v ", v.NoAnalytics)
		br()
	}
	br()
	fmt.Printf("\nNumber of Provider Organizations to investigate %v", len(*data.Org))
	br()
	for i, v := range *data.Org {
		fmt.Printf("\n\t ## Provider Organization %v - %v", i, v.Name)
		br()
		fmt.Printf("\n\t\t POrg %v - %v ", i, v.Name)
		fmt.Printf("\n\t\t Number of Members             -  %v ", v.NoMembers)
		fmt.Printf("\n\t\t Number of Member Invitations  -  %v ", v.NoMemberInvites)
		fmt.Printf("\n\t\t Number of DraftAPI            -  %v ", v.NoDraftAPI)
		fmt.Printf("\n\t\t Average API Size (bytes)      -  %v ", v.AvgAPISize)
		fmt.Printf("\n\t\t Max API Size (bytes)          -  %v ", v.MaxAPISize)
		fmt.Printf("\n\t\t Number of Draft Products      -  %v ", v.NoDraftProduct)
		fmt.Printf("\n\t\t Number of TLS Profiles        -  %v ", v.NoTLSProfile)
		fmt.Printf("\n\t\t Number of OAuth PROIVDERS     -  %v ", v.NoOAuthProvider)
		fmt.Printf("\n\t\t Number of User Registries     -  %v ", v.UserRegistries)
		fmt.Printf("\n\t\t Number of Key Stores          -  %v ", v.NoKeyStore)
		fmt.Printf("\n\t\t Number of Trust Stores        -  %v ", v.NoTrustStore)
		br()
		fmt.Printf("\n\t\t Number of catalogs            -  %v ", len(*v.Catalog))
		br()
		for ci, cv := range *v.Catalog {
			fmt.Printf("\n\t\t\t ## Catalog %v - %v", ci, cv.Name)
			br()

			fmt.Printf("\n\t\t\t\t Number of Members                -  %v ", cv.NoMember)
			fmt.Printf("\n\t\t\t\t Number of Member Invitations     -  %v ", cv.NoMemberInvites)
			fmt.Printf("\n\t\t\t\t Number of APIs                   -  %v ", cv.NoAPI)
			fmt.Printf("\n\t\t\t\t Number of Products               -  %v ", cv.NoProduct)
			fmt.Printf("\n\t\t\t\t Average API Size                 -  %v ", cv.AvgAPISize)
			fmt.Printf("\n\t\t\t\t Max API Size                     -  %v ", cv.MaxAPISize)
			fmt.Printf("\n\t\t\t\t Number of Consumer Orgs          -  %v ", cv.NoConsumerOrg)
			fmt.Printf("\n\t\t\t\t Portal Enabled                   -  %v ", cv.Portal)
			fmt.Printf("\n\t\t\t\t Number of TLS Profiles           -  %v ", cv.NoTLSProfile)
			fmt.Printf("\n\t\t\t\t Number of OAuth Providers        -  %v ", cv.NoOAuthProvider)
			fmt.Printf("\n\t\t\t\t Number of User Registries        -  %v ", cv.UserRegistries)
			fmt.Printf("\n\t\t\t\t Number of Spaces                 -  %v ", cv.NoSpace)
			fmt.Printf("\n\t\t\t\t Number of Subscriptions          -  %v ", cv.Subscriptions)
		}
		br()
	}
	Trace("")
	Trace("")
	Trace("************************************")
	TraceExit("printData")
}

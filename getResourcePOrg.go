package main

import (
	"fmt"
	"strings"
)

func GetPOrg(orgDetails orgNameId, access_token string) organization {
	TraceEnter("Get Org")
	if access_token == "unset" {
		userDetails = userCreds{}
		token = login(PromptCredentials(orgDetails.name))
	} else {
		token = access_token
	}
	// promptToAnalyse()
	orgId := orgDetails.id
	no, max, avg := GetDraftAPIs(orgId)
	toReturn := organization{
		name:            orgDetails.name,
		catalog:         GetCatalogs(orgId),
		noMembers:       GetMembers(orgId),
		noMemberInvites: GetMemberInvites(orgId),
		noDraftAPI:      no,
		avgAPISize:      avg,
		maxAPISize:      max,
		noDraftProduct:  GetDraftProducts(orgId),
		noTLSProfile:    GetTLSProfilesPorg(orgId),
		noKeyStore:      GetKeyStorePorg(orgId),
		noTrustStore:    GetTrustStorePorg(orgId),
		userRegistries:  GetUserRegPOrg(orgId),
		noOAuthProvider: GetOAuthPOrg(orgId)}
	TraceExitReturn("Get Org", toReturn)
	return toReturn

}

func GetMembers(orgId string) int {
	TraceEnter("GetMembers")
	toReturn := getData("orgs", orgId, "members")
	TraceExitReturn("GetMembers", toReturn)
	return toReturn
}
func GetMemberInvites(orgId string) int {
	TraceEnter("GetMemberInvites")
	toReturn := getData("orgs", orgId, "member-invitations")
	TraceExitReturn("GetMemberInvites", toReturn)
	return toReturn
}
func GetCMembers(orgId string) int {
	TraceEnter("GetCMembers")
	toReturn := getData("catalogs", orgId, "members")
	TraceExitReturn("GetCMembers", toReturn)
	return toReturn
}
func GetCMemberInvites(orgId string) int {
	TraceEnter("GetCMemberInvites")
	toReturn := getData("catalogs", orgId, "member-invitations")
	TraceExitReturn("GetCMemberInvites", toReturn)
	return toReturn
}
func GetDraftProducts(orgId string) int {
	TraceEnter("GetDraftProducts")
	toReturn := getData("orgs", orgId+"/drafts", "draft-products")
	TraceExitReturn("GetDraftProducts", toReturn)
	return toReturn
}
func GetProducts(orgId string) int {
	TraceEnter("GetProducts")
	toReturn := getData("catalogs", orgId, "products")
	TraceExitReturn("GetProducts", toReturn)
	return toReturn
}
func GetCatalogs(orgId string) *[]catalog {
	TraceEnter("GetCatalogs")
	path := "orgs/" + orgId + "/catalogs"
	jsonObj := APICRequest(path)
	toReturn := make([]catalog, int(jsonObj["total_results"].(float64)))
	for i, v := range jsonObj["results"].([]interface{}) {
		Log("\tInvestigating Catalog " + v.(map[string]interface{})["title"].(string))
		catId := v.(map[string]interface{})["id"].(string)
		no, max, avg := GetAPIs("catalogs/" + orgId + "/" + catId)
		noc, noa := GetCOrgs(orgId + "/" + catId)
		toReturn[i] = catalog{
			name:            v.(map[string]interface{})["title"].(string),
			noMember:        GetCMembers(orgId + "/" + catId),
			noMemberInvites: GetCMemberInvites(orgId + "/" + catId),
			noAPI:           no,
			noProduct:       GetProducts(orgId + "/" + catId),
			avgAPISize:      avg,
			maxAPISize:      max,
			noSpace:         GetSpaces(orgId + "/" + catId),
			noConsumerOrg:   noc,
			portal:          GetPortal(orgId + "/" + catId),
			noTLSProfile:    GetCTLSProfilesPorg(orgId + "/" + catId),
			userRegistries:  GetCUserRegPOrg(orgId + "/" + catId),
			noOAuthProvider: GetCOAuthPOrg(orgId + "/" + catId),
			applications:    noa,
		}

		// size := APICRequestSize(v.(map[string]interface{})["url"].(string))
		// total = +size
		// if size > max {
		//   max = size
		// }
	}
	TraceExitReturn("GetCatalogs", toReturn)
	return &toReturn
}
func GetDraftAPIs(orgId string) (int, int, int) {
	TraceEnter("GetDraftAPIs")
	no, max, avg := ProcessAPI("orgs/" + orgId + "/drafts/draft-apis")

	TraceExitReturn("GetDraftAPIs", fmt.Sprintf("API Nos %v\t Max Size %v\t Avg Size %v", no, max, avg))
	return no, max, avg
}
func GetAPIs(orgId string) (int, int, int) {
	TraceEnter("GetAPIs")
	no, max, avg := ProcessAPI(orgId + "/apis")
	TraceExitReturn("GetAPIs", fmt.Sprintf("API Nos %v\t Max Size %v\t Avg Size %v", no, max, avg))
	return no, max, avg
}
func ProcessAPI(path string) (int, int, int) {
	TraceEnter("ProcessAPI")
	jsonObj := APICRequest(path)
	if jsonObj == nil {
		return -1, -1, -1
	}
	no := int(jsonObj["total_results"].(float64))
	max := 0
	total := 0
	for _, v := range jsonObj["results"].([]interface{}) {
		size := APICRequestSize(v.(map[string]interface{})["url"].(string))
		total = +size
		if size > max {
			max = size
		}
	}
	avg := 0
	if no != 0 {
		avg = total / no
	}

	TraceExitReturn("ProcessAPI", fmt.Sprintf("API Nos %v\t Max Size %v\t Avg Size %v", no, max, avg))
	return no, max, avg
}
func GetUserRegPOrg(orgId string) int {
	TraceEnter("GetUserRegPOrg")
	toReturn := getData("orgs", orgId, "user-registries")
	TraceExitReturn("GetUserRegPOrg", toReturn)
	return toReturn
}
func GetCUserRegPOrg(orgId string) int {
	TraceEnter("GetCUserRegPOrg")
	toReturn := getData("catalogs", orgId, "configured-api-user-registries")
	TraceExitReturn("GetCUserRegPOrg", toReturn)
	return toReturn
}
func GetOAuthPOrg(orgId string) int {
	TraceEnter("GetOAuthPOrg")
	toReturn := getData("orgs", orgId, "oauth-providers")
	TraceExitReturn("GetOAuthPOrg", toReturn)
	return toReturn
}
func GetCOAuthPOrg(orgId string) int {
	TraceEnter("GetCOAuthPOrg")
	toReturn := getData("catalogs", orgId, "configured-oauth-providers")
	TraceExitReturn("GetCOAuthPOrg", toReturn)
	return toReturn
}
func GetTLSProfilesPorg(orgId string) int {
	TraceEnter("GetTLSProfilesPorg")
	toReturn := getData("orgs", orgId, "tls-client-profiles")
	TraceExitReturn("GetTLSProfilesPorg", toReturn)
	return toReturn
}

func GetCTLSProfilesPorg(orgId string) int {
	TraceEnter("GetCTLSProfilesPorg")
	toReturn := getData("catalogs", orgId, "configured-tls-client-profiles")
	TraceExitReturn("GetCTLSProfilesPorg", toReturn)
	return toReturn
}
func GetSpaces(orgId string) int {
	TraceEnter("GetSpaces")
	toReturn := getData("catalogs", orgId, "spaces")
	TraceExitReturn("GetSpaces", toReturn)
	return toReturn
}
func GetCOrgs(orgId string) (int, int) {
	TraceEnter("GetCOrgs")
	path := "catalogs/" + orgId + "/consumer-orgs"
	jsonObj := APICRequest(path)
	if jsonObj == nil {
		TraceExitReturn("GetCOrgs", "ERROR")
		return -1, -1
	}
	no := int(jsonObj["total_results"].(float64))
	noa := 0
	for _, v := range jsonObj["results"].([]interface{}) {
		url := v.(map[string]interface{})["url"].(string) + "/apps"
		Trace(url)
		Trace(strings.Split(url, "/api/")[1])
		jsonObj = APICRequest(strings.Split(url, "/api/")[1])
		if jsonObj == nil {
			noa = -1
			break
		} else {
			noa = noa + int(jsonObj["total_results"].(float64))
		}
	}

	TraceExitReturn("GetCOrgs", fmt.Sprintf("ConsumerOrg %v \t Applications %v", no, noa))
	return no, noa
}

func GetKeyStorePorg(orgId string) int {
	TraceEnter("GetKeyStorePorg")
	toReturn := getData("orgs", orgId, "keystores")
	TraceExitReturn("GetKeyStorePorg", toReturn)
	return toReturn
}
func GetTrustStorePorg(orgId string) int {
	TraceEnter("GetTrustStorePorg")
	toReturn := getData("orgs", orgId, "truststores")
	TraceExitReturn("GetTrustStorePorg", toReturn)
	return toReturn
}
func GetPortal(orgId string) bool {
	TraceEnter("GetPortal")
	path := "catalogs/" + orgId + "/settings"
	jsonObj := APICRequest(path)
	if jsonObj == nil {
		TraceExitReturn("GetCOrgs", "ERROR")

		return false
	}
	toReturn := false
	if jsonObj["portal"].(map[string]interface{})["type"].(string) == "none" {
		toReturn = false
	} else {
		toReturn = true
	}
	TraceExitReturn("GetPortal", toReturn)
	return toReturn

}

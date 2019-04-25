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
	no, max, avg := ProcessAPI("orgs/"+orgId+"/drafts/draft-apis", "draft_api")
	toReturn := organization{
		name:            orgDetails.name,
		catalog:         GetCatalogs(orgDetails.name, orgId),
		noMembers:       getData("orgs", orgId, "members"),
		noMemberInvites: getData("orgs", orgId, "member-invitations"),
		noDraftAPI:      no,
		avgAPISize:      avg,
		maxAPISize:      max,
		noDraftProduct:  getData("orgs", orgId+"/drafts", "draft-products"),
		noTLSProfile:    getData("orgs", orgId, "tls-client-profiles"),
		noKeyStore:      getData("orgs", orgId, "keystores"),
		noTrustStore:    getData("orgs", orgId, "truststores"),
		userRegistries:  getData("orgs", orgId, "user-registries"),
		noOAuthProvider: getData("orgs", orgId, "oauth-providers")}
	TraceExitReturn("Get Org", toReturn)
	return toReturn

}
func GetSpaceIDs(orgId string) *[]string {
	TraceEnter("GetSpaceIDs")
	path := "catalogs/" + orgId + "/spaces"
	jsonObj := APICRequest(path)
	len := int(jsonObj["total_results"].(float64))
	toReturn := make([]string, len)
	for i, v := range jsonObj["results"].([]interface{}) {
		toReturn[i] = v.(map[string]interface{})["id"].(string)
	}
	TraceExitReturn("GetSpaceIDs", toReturn)
	return &toReturn
}

func GetCatalogs(orgName string, orgId string) *[]catalog {
	TraceEnter("GetCatalogs")
	path := "orgs/" + orgId + "/catalogs"
	jsonObj := APICRequest(path)
	len := int(jsonObj["total_results"].(float64))
	toReturn := make([]catalog, len)

	chanList := make(map[string]chan bool, len)

	for i, v := range jsonObj["results"].([]interface{}) {
		name := v.(map[string]interface{})["title"].(string)
		chanList[name] = make(chan bool)
		go AsyncGetCat(orgName, name, i, v, toReturn, orgId, chanList)
	}
	for _, v := range jsonObj["results"].([]interface{}) {
		name := v.(map[string]interface{})["title"].(string)
		<-chanList[name]
	}
	TraceExitReturn("GetCatalogs", toReturn)
	return &toReturn
}

func ProcessAPI(path string, apiType string) (int, int, int) {
	TraceEnter("ProcessAPI")
	jsonObj := APICRequest(path)
	if jsonObj == nil {
		return -1, -1, -1
	}
	no := int(jsonObj["total_results"].(float64))
	max := 0
	total := 0
	for _, v := range jsonObj["results"].([]interface{}) {
		size := APICRequestSize(v.(map[string]interface{})["url"].(string) + "?fields=add%28" + apiType + "%29")
		total = total + size
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

func GetCOrgs(orgId string) (int, int, int) {
	TraceEnter("GetCOrgs")
	path := "catalogs/" + orgId + "/consumer-orgs"
	jsonObj := APICRequest(path)
	if jsonObj == nil {
		TraceExitReturn("GetCOrgs", "ERROR")
		return -1, -1, -1
	}
	no := int(jsonObj["total_results"].(float64))
	noa := 0
	nos := 0
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
		for _, av := range jsonObj["results"].([]interface{}) {

			nos = nos + getData("apps", orgId+"/"+v.(map[string]interface{})["id"].(string)+"/"+av.(map[string]interface{})["id"].(string), "subscriptions")
		}

	}

	TraceExitReturn("GetCOrgs", fmt.Sprintf("ConsumerOrg %v \t Applications %v \t Subscriptions %v", no, noa, nos))
	return no, noa, nos
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
func AsyncGetCat(orgName string, name string, i int, v interface{}, toReturn []catalog, orgId string, chanList map[string]chan bool) {
	TraceEnter("AsyncGetCat")
	Log("\tInvestigating Catalog " + name + " in " + orgName)
	catId := v.(map[string]interface{})["id"].(string)
	no, max, avg := ProcessAPI("catalogs/"+orgId+"/"+catId+"/apis", "api")
	noc, noa, nos := GetCOrgs(orgId + "/" + catId)
	spaceNo := getData("catalogs", orgId+"/"+catId, "spaces")
	noTLSProfile := -1
	userRegistries := -1
	noOAuthProvider := -1

	if spaceNo < 1 {
		noTLSProfile = getData("catalogs", orgId+"/"+catId, "configured-tls-client-profiles")
		userRegistries = getData("catalogs", orgId+"/"+catId, "configured-api-user-registries")
		noOAuthProvider = getData("catalogs", orgId+"/"+catId, "configured-oauth-providers")
	} else {
		ids := GetSpaceIDs(orgId + "/" + catId)
		noTLSProfile = getDataSpaces("catalogs", orgId+"/"+catId, "configured-tls-client-profiles", ids)
		userRegistries = getDataSpaces("catalogs", orgId+"/"+catId, "configured-api-user-registries", ids)
		noOAuthProvider = getDataSpaces("catalogs", orgId+"/"+catId, "configured-oauth-providers", ids)
	}
	toReturn[i] = catalog{
		name:            name,
		noMember:        getData("catalogs", orgId+"/"+catId, "members"),
		noMemberInvites: getData("catalogs", orgId+"/"+catId, "member-invitations"),
		noAPI:           no,
		noProduct:       getData("catalogs", orgId+"/"+catId, "products"),
		avgAPISize:      avg,
		maxAPISize:      max,
		noSpace:         spaceNo,
		noConsumerOrg:   noc,
		portal:          GetPortal(orgId + "/" + catId),
		noTLSProfile:    noTLSProfile,
		userRegistries:  userRegistries,
		noOAuthProvider: noOAuthProvider,
		applications:    noa,
		subscriptions:   nos,
		webhooks:        GetWebhooks(orgName, orgId, name, catId),
		tasks:           GetTasks(orgName, orgId, name, catId),
	}

	chanList[name] <- true
	TraceExit("AsyncGetCat")
}

func GetWebhooks(orgName string, orgId string, catName string, catId string) *[]webhook {
	TraceEnter("GetWebhooks")
	path := "catalogs/" + orgId + "/" + catId + "/webhooks"
	jsonObj := APICRequest(path)
	if jsonObj == nil {
		return nil
	}
	toReturn := make([]webhook, int(jsonObj["total_results"].(float64)))

	for i, v := range jsonObj["results"].([]interface{}) {
		vMap := v.(map[string]interface{})
		toReturn[i] = webhook{
			webhookId:        vMap["id"].(string),
			catalogName:      catName,
			organization:     orgId,
			organizationName: orgName,
			catalog:          catId,
			state:            vMap["state"].(string),
			created_at:       vMap["created_at"].(string),
			updated_at:       vMap["updated_at"].(string),
			level:            vMap["level"].(string),
			title:            vMap["title"].(string),
		}
	}
	TraceExitReturn("GetWebhooks", fmt.Sprintf("webhooks %v", toReturn))
	return &toReturn
}

func GetTasks(orgName string, orgId string, catName string, catId string) *[]task {
	TraceEnter("GetTasks")
	path := "catalogs/" + orgId + "/" + catId + "/tasks"
	jsonObj := APICRequest(path)
	if jsonObj == nil {
		return nil
	}
	toReturn := make([]task, int(jsonObj["total_results"].(float64)))
	for i, v := range jsonObj["results"].([]interface{}) {
		vMap := v.(map[string]interface{})
		toReturn[i] = task{
			taskId:       vMap["id"].(string),
			catalogName:  catName,
			organization: orgId,
			name:         vMap["name"].(string),
			state:        vMap["state"].(string),
			created_at:   vMap["created_at"].(string),
			updated_at:   vMap["updated_at"].(string),
			content:      v.(map[string]interface{}),
		}

		Log(v.(map[string]interface{}))
	}

	TraceExitReturn("GetTasks", fmt.Sprintf("webhooks %v", toReturn))
	return &toReturn
}

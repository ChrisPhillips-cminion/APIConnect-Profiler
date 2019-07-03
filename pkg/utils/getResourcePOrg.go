package utils
/*
      Licensed Materials - Property of IBM
      © IBM Corp. 2019
*/
import (
	"errors"
	"fmt"
	"github.com/chrisphillips-cminion/apiprofile/pkg/model"
	"strings"
)

func GetPOrg(orgDetails orgNameId, access_token string) (model.Organization, error) {
	TraceEnter("Get Org")
	if access_token == "unset" {
		Vars.UserDetails = model.UserCreds{}
		err := errors.New("Not an error")
		Vars.Token, err = Login(PromptCredentials(orgDetails.Name, "provider"))

		if err != nil {
			TraceExit("GetPOrg ERROR")
			return model.Organization{}, err
		}
	} else {
		Vars.Token = access_token
	}
	// promptToAnalyse()
	orgId := orgDetails.id
	no, max, avg, _ := ProcessAPI("orgs/"+orgId+"/drafts/draft-apis", "draft_api")
	cat, err := GetCatalogs(orgDetails.Name, orgId)
	member, err := getData("orgs", orgId, "members")
	mi, err := getData("orgs", orgId, "member-invitations")
	dp, err := getData("orgs", orgId+"/drafts", "draft-products")
	NoTLSProfile, err := getData("orgs", orgId, "tls-client-profiles")
	NoKeyStore, err := getData("orgs", orgId, "keystores")
	NoTrustStore, err := getData("orgs", orgId, "truststores")
	UserRegistries, err := getData("orgs", orgId, "user-registries")
	NoOAuthProvider, err := getData("orgs", orgId, "oauth-providers")
	if err != nil {
		TraceExit("GetPOrg ERROR")
		return model.Organization{
			Name:            orgDetails.Name,
			Catalog:         &[]model.Catalog{},
			NoMembers:       -1,
			NoMemberInvites: -1,
			NoDraftAPI:      -1,
			AvgAPISize:      -1,
			MaxAPISize:      -1,
			NoDraftProduct:  -1,
			NoTLSProfile:    -1,
			NoKeyStore:      -1,
			NoTrustStore:    -1,
			UserRegistries:  -1,
			NoOAuthProvider: -1,
			Error:           err.Error(),
		}, err
	}
	toReturn := model.Organization{
		Name:            orgDetails.Name,
		Catalog:         cat,
		NoMembers:       member,
		NoMemberInvites: mi,
		NoDraftAPI:      no,
		AvgAPISize:      avg,
		MaxAPISize:      max,
		NoDraftProduct:  dp,
		NoTLSProfile:    NoTLSProfile,
		NoKeyStore:      NoKeyStore,
		NoTrustStore:    NoTrustStore,
		UserRegistries:  UserRegistries,
		NoOAuthProvider: NoOAuthProvider,
		Error:           ""}
	TraceExitReturn("Get Org", toReturn)
	return toReturn, nil

}
func GetSpaceIDs(orgId string) (*[]string, error) {
	TraceEnter("GetSpaceIDs")
	path := "catalogs/" + orgId + "/spaces"
	jsonObj, err := APICRequest(path)
	if err != nil {
		TraceExit("GetSpaceIDs ERROR")
		return nil, err
	}
	len := int(jsonObj["total_results"].(float64))
	toReturn := make([]string, len)
	for i, v := range jsonObj["results"].([]interface{}) {
		toReturn[i] = v.(map[string]interface{})["id"].(string)
	}
	TraceExitReturn("GetSpaceIDs", toReturn)
	return &toReturn, nil
}

func GetCatalogs(orgName string, orgId string) (*[]model.Catalog, error) {
	TraceEnter("GetCatalogs")
	path := "orgs/" + orgId + "/catalogs"
	jsonObj, err := APICRequest(path)
	if err != nil {
		TraceExit("GetCatalogs ERROR")
		return nil, err
	}
	len := int(jsonObj["total_results"].(float64))
	toReturn := make([]model.Catalog, len)

	chanList := make(map[string]chan bool, len)

	for i, v := range jsonObj["results"].([]interface{}) {
		name := v.(map[string]interface{})["name"].(string)
		chanList[name] = make(chan bool)
		go AsyncGetCat(orgName, name, i, v, toReturn, orgId, chanList)
	}
	for _, v := range jsonObj["results"].([]interface{}) {
		name := v.(map[string]interface{})["name"].(string)
		<-chanList[name]
	}
	TraceExitReturn("GetCatalogs", toReturn)
	return &toReturn, nil
}

func ProcessAPI(path string, apiType string) (int, int, int, error) {
	TraceEnter("ProcessAPI")
	jsonObj, err := APICRequest(path)
	if err != nil {
		TraceExit("GetCatalogs ERROR")
		return -1, -1, -1, err
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
	return no, max, avg, nil
}

func GetCOrgs(orgId string) (int, int, int, error) {
	TraceEnter("GetCOrgs")
	path := "catalogs/" + orgId + "/consumer-orgs"
	jsonObj, err := APICRequest(path)
	if err != nil {
		TraceExit("GetCOrgs ERROR")
		return -1, -1, -1, err
	}

	no := int(jsonObj["total_results"].(float64))
	noa := 0
	nos := 0
	errorFlag := false
	for _, v := range jsonObj["results"].([]interface{}) {
		url := v.(map[string]interface{})["url"].(string) + "/apps"
		Trace(url)
		Trace(strings.Split(url, "/api/")[1])
		jsonObj, err = APICRequest(strings.Split(url, "/api/")[1])
		if err != nil {
			noa = -1
			break
		} else {
			noa = noa + int(jsonObj["total_results"].(float64))
		}

		for _, av := range jsonObj["results"].([]interface{}) {
			v, err := getData("apps", orgId+"/"+v.(map[string]interface{})["id"].(string)+"/"+av.(map[string]interface{})["id"].(string), "subscriptions")
			if err != nil {
				errorFlag = true
			}
			nos = nos + v
		}

	}
	if errorFlag {
		nos = -1
	}

	TraceExitReturn("GetCOrgs", fmt.Sprintf("ConsumerOrg %v \t Applications %v \t Subscriptions %v", no, noa, nos))
	return no, noa, nos, nil
}

func GetPortal(orgId string) (bool, error) {
	TraceEnter("GetPortal")
	path := "catalogs/" + orgId + "/settings"
	jsonObj, err := APICRequest(path)
	if err != nil {
		TraceExit("GetPortal ERROR")
		return false, err
	}
	toReturn := false
	if jsonObj["portal"].(map[string]interface{})["type"].(string) == "none" {
		toReturn = false
	} else {
		toReturn = true
	}
	TraceExitReturn("GetPortal", toReturn)
	return toReturn, nil

}
func AsyncGetCat(orgName string, name string, i int, v interface{}, toReturn []model.Catalog, orgId string, chanList map[string]chan bool) error {
	TraceEnter("AsyncGetCat")
	Log("\tInvestigating Catalog " + name + " in " + orgName)
	catId := v.(map[string]interface{})["id"].(string)
	no, max, avg, _ := ProcessAPI("catalogs/"+orgId+"/"+catId+"/apis", "api")
	noc, noa, nos, _ := GetCOrgs(orgId + "/" + catId)
	spaceNo, _ := getData("catalogs", orgId+"/"+catId, "spaces")
	NoTLSProfile := -1
	userRegistries := -1
	NoOAuthProvider := -1

	if spaceNo < 1 {
		NoTLSProfile, _ = getData("catalogs", orgId+"/"+catId, "configured-tls-client-profiles")
		userRegistries, _ = getData("catalogs", orgId+"/"+catId, "configured-api-user-registries")
		NoOAuthProvider, _ = getData("catalogs", orgId+"/"+catId, "configured-oauth-providers")
	} else {
		ids, err := GetSpaceIDs(orgId + "/" + catId)
		if err != nil {
			chanList[name] <- true
			TraceExit("AsyncGetCat ERROR")
			return err
		}
		NoTLSProfile = getDataSpaces("catalogs", orgId+"/"+catId, "configured-tls-client-profiles", ids)
		userRegistries = getDataSpaces("catalogs", orgId+"/"+catId, "configured-api-user-registries", ids)
		NoOAuthProvider = getDataSpaces("catalogs", orgId+"/"+catId, "configured-oauth-providers", ids)
	}
	mem, _ := getData("catalogs", orgId+"/"+catId, "members")
	mi, _ := getData("catalogs", orgId+"/"+catId, "member-invitations")
	prod, _ := getData("catalogs", orgId+"/"+catId, "products")
	portal, err := GetPortal(orgId + "/" + catId)
	if err != nil {
		Log(fmt.Sprintf("Unable to get portal state for %v", name))
	}
	wh, _ := GetWebhooks(orgName, orgId, name, catId)
	t, _ := GetTasks(orgName, orgId, name, catId)
	toReturn[i] = model.Catalog{
		Name:            name,
		NoMember:        mem,
		NoMemberInvites: mi,
		NoAPI:           no,
		NoProduct:       prod,
		AvgAPISize:      avg,
		MaxAPISize:      max,
		NoSpace:         spaceNo,
		NoConsumerOrg:   noc,
		Portal:          portal,
		NoTLSProfile:    NoTLSProfile,
		UserRegistries:  userRegistries,
		NoOAuthProvider: NoOAuthProvider,
		Applications:    noa,
		Subscriptions:   nos,
		Webhooks:        wh,
		Tasks:           t,
	}

	chanList[name] <- true
	TraceExit("AsyncGetCat")
	return nil
}

func GetWebhooks(orgName string, orgId string, catName string, catId string) (*[]model.Webhook, error) {
	TraceEnter("GetWebhooks")
	path := "catalogs/" + orgId + "/" + catId + "/webhooks"
	jsonObj, err := APICRequest(path)
	if err != nil {
		Trace(err.Error())
		TraceExit("GetWebhooks")
		return nil, err
	}
	toReturn := make([]model.Webhook, int(jsonObj["total_results"].(float64)))

	for i, v := range jsonObj["results"].([]interface{}) {
		vMap := v.(map[string]interface{})
		toReturn[i] = model.Webhook{
			WebhookId:        vMap["id"].(string),
			CatalogName:      catName,
			Organization:     orgId,
			OrganizationName: orgName,
			Catalog:          catId,
			State:            vMap["state"].(string),
			Created_at:       vMap["created_at"].(string),
			Updated_at:       vMap["updated_at"].(string),
			Level:            vMap["level"].(string),
			Title:            vMap["title"].(string),
		}
	}
	TraceExitReturn("GetWebhooks", fmt.Sprintf("webhooks %v", toReturn))
	return &toReturn, nil
}

func GetTasks(orgName string, orgId string, catName string, catId string) (*[]model.Task, error) {
	TraceEnter("GetTasks")
	path := "catalogs/" + orgId + "/" + catId + "/tasks"
	jsonObj, err := APICRequest(path)
	if err != nil {
		Trace(err.Error())
		TraceExit("GetTasks")
		return nil, err
	}
	toReturn := make([]model.Task, int(jsonObj["total_results"].(float64)))
	for i, v := range jsonObj["results"].([]interface{}) {
		vMap := v.(map[string]interface{})
		toReturn[i] = model.Task{
			TaskId:       vMap["id"].(string),
			CatalogName:  catName,
			Organization: orgId,
			Name:         vMap["name"].(string),
			State:        vMap["state"].(string),
			Created_at:   vMap["created_at"].(string),
			Updated_at:   vMap["updated_at"].(string),
			Content:      v.(map[string]interface{}),
		}

		Log(v.(map[string]interface{}))
	}

	TraceExitReturn("GetTasks", fmt.Sprintf("webhooks %v", toReturn))
	return &toReturn, nil
}

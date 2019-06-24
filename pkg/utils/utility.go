package utils
/*
      Licensed Materials - Property of IBM
      © IBM Corp. 2019
*/

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/chrisphillips-cminion/apiprofile/pkg/model"
	"gopkg.in/AlecAivazis/survey.v1"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	// "strconv"
	// "os"
)

func PromptServer() string {
	TraceEnter("PromptServer")
	toReturn := Vars.Server
	if Vars.Server == "unset" {
		prompt := &survey.Input{Message: "APIC Manager or Admin Endpoint"}
		survey.AskOne(prompt, &toReturn, nil)
	}
	TraceExitReturn("PromptServer", toReturn)
	return toReturn
}
func GetRealm(scope string) error {

	path := "cloud/" + scope + "/identity-providers"
	jsonObj, err := APICRequest(path)

	if err != nil {
		Log(err.Error())
		return err
	}
	length := int(jsonObj["total_results"].(float64))
	idPs := make([]string, length)

	Trace(jsonObj)
	for i, v := range jsonObj["results"].([]interface{}) {
		idPs[i] = scope + "/" + v.(map[string]interface{})["name"].(string)
	}

	prompt := []*survey.Question{
		{
			Name: "idp",
			Prompt: &survey.Select{
				Message: "Pleas select your realm or identity provider:",
				Options: idPs,
			}}}
	// survey.AskOne(prompt, &, nil)
	survey.Ask(prompt, &Vars.UserDetails.Realm)
	return nil
}
func PromptCredentials(porg string, scope string) model.UserCreds {
	TraceEnter("PromptCredentials")
	if Vars.UserDetails.Username == "" {
		prompt := &survey.Input{Message: "Username"}
		survey.AskOne(prompt, &Vars.UserDetails.Username, nil)
	}
	if Vars.UserDetails.Password == "" {
		promptPass := &survey.Password{Message: "Password"}
		survey.AskOne(promptPass, &Vars.UserDetails.Password, nil)
	}
	if Vars.UserDetails.Realm == "" {

		GetRealm(scope)
	}
	TraceExit("PromptCredentials")
	return Vars.UserDetails
}
func Log(s interface{}) {
	fmt.Printf("\t %v\n", s)
}
func Trace(s interface{}) {
	if Vars.Debug {
		log.Printf("\t %v\n", s)
	}
}
func TraceExit(s string) {
	if Vars.Debug {
		log.Printf("------------------------------------\n")
		log.Printf("\t Exiting - %v\n", s)
		log.Printf("------------------------------------\n")
		log.Printf("------------------------------------\n")
	}
}
func TraceExitReturn(s string, i interface{}) {
	if Vars.Debug {
		log.Printf("------------------------------------\n")
		log.Printf("\t Exiting - %v\n", s)
		log.Printf("------------\n")
		log.Printf("Returning - %v\n", i)
		log.Printf("------------------------------------\n")
		log.Printf("------------------------------------\n")
	}
}
func TraceEnter(s string) {
	if Vars.Debug {
		log.Printf("------------------------------------\n")
		log.Printf("\t Entering - %v\n", s)
		log.Printf("------------------------------------\n")
	}
}

//
type gzreadCloser struct {
	*gzip.Reader
	io.Closer
}

func (gz gzreadCloser) Close() error {
	return gz.Closer.Close()
}

func APICRequest(path string) (map[string]interface{}, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	TraceEnter("APICRequest")
	Trace("path " + "https://" + Vars.Server + "/api/" + path)
	req, err := http.NewRequest("GET", "https://"+Vars.Server+"/api/"+path, nil)
	if err != nil {
		Trace(err.Error())
		TraceExit("APICRequest Error")
		return nil, err
	}
	// Trace(token)
	req.Header.Set("Authorization", "Bearer "+Vars.Token)
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Connection", "keep-alive")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// Log(err.Error())

		Trace(err.Error())
		TraceExit("APICRequest Error")
		return nil, err
	}
	if resp.Status != "200 OK" {
		Trace(fmt.Sprintf("HTTP Status '%v'", resp.Status))
		TraceExit("APICRequest ERROR")
		return nil, errors.New(fmt.Sprintf("HTTP Status '%v'", resp.Status))
	}
	defer resp.Body.Close()

	jsonObj := make(map[string]interface{})
	bodyBytes, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		Log(err2.Error())
		TraceExit("APICRequest Error")
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &jsonObj)
	if err != nil {
		Log(err.Error())
		TraceExit("APICRequest Error")
		return nil, err
	}
	TraceExit("APICRequest")
	return jsonObj, nil
}

func APICRequestSize(url string) int {
	TraceEnter("APICRequest")
	Trace("path " + url)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Trace(err.Error())
		return -1
	}
	// Trace(token)
	req.Header.Set("Authorization", "Bearer "+Vars.Token)
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Connection", "keep-alive")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Trace(err.Error())
		return -1

	}
	if resp.Status != "200 OK" {
		log.Printf("HTTP Status '%v'", resp.Status)
		return -1
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return len(bodyBytes)
}

func Login(UC model.UserCreds) (string, error) {
	TraceEnter("login")
	type Payload struct {
		Username     string `json:"username"`
		Password     string `json:"password"`
		Realm        string `json:"realm"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		GrantType    string `json:"grant_type"`
	}

	data := Payload{
		Username:     UC.Username,
		Password:     UC.Password,
		Realm:        UC.Realm,
		ClientID:     "599b7aef-8841-4ee2-88a0-84d49c4d6ff2",
		ClientSecret: "0ea28423-e73b-47d4-b40e-ddb45c48bb0c",
		GrantType:    "password",
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {

		Trace(err.Error())
		return "error", err

	}
	body := bytes.NewReader(payloadBytes)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("POST", "https://"+Vars.Server+"/api/token", body)

	if err != nil {
		Trace(err.Error())
		return "error", err

	}
	req.Header.Set("Origin", "https://apim.lts.apicww.cloud")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Referer", "https://apim.lts.apicww.cloud/auth/admin/sign-in/")
	req.Header.Set("Connection", "keep-alive")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Trace(err.Error())
		return "error", err
	}
	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		log.Printf("HTTP Status %v", resp.Status)
		// os.Exit(2)
		return "error", err
	}

	jsonObj := make(map[string]interface{})

	zr, err := gzip.NewReader(resp.Body)
	if err != nil {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		Trace(bodyBytes)
		Trace(err.Error())
	} else {
		gzr := gzreadCloser{zr, resp.Body}
		resp.Body = gzr
		defer gzr.Close()
	}

	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	// Trace(string(bodyBytes))
	if err2 != nil {
		Trace(err2.Error())
		TraceExit("login ERROR")
		return "error", err
	}
	err = json.Unmarshal(bodyBytes, &jsonObj)

	if err != nil {
		Trace(err.Error())
		TraceExit("login ERROR")
		return "error", err
	}
	TraceExit("login")
	return jsonObj["access_token"].(string), nil
}

func getData(type1 string, orgId string, keyword string) (int, error) {
	TraceEnter("getData")
	path := type1 + "/" + orgId + "/" + keyword
	jsonObj, err := APICRequest(path)
	if err != nil {
		TraceExit("getData ERROR")
		return -1, err
	}
	toReturn := int(jsonObj["total_results"].(float64))
	TraceExitReturn("getData", fmt.Sprintf("%v %v", keyword, toReturn))
	return toReturn, nil
}

//https://apim.lts.apicww.cloud/api/spaces/8643dbab-b431-4602-a083-8d8fa29d2f6e/42298b69-0edf-4cc3-83b4-232786976bbe/ae90bbac-9cc5-4b9e-9197-215efdacd3f8/configured-tls-client-profiles
func getDataSpaces(type1 string, orgId string, keyword string, spaces *[]string) int {
	TraceEnter("getDataSpaces")
	count := 0
	errorFlag := false
	for _, v := range *spaces {
		no, err := getData("spaces", orgId+"/"+v, keyword)
		if err != nil {
			errorFlag = true
			Trace("Error getting a Spaces Details")
			Trace(err.Error())
		}
		count = count + no
	}
	if errorFlag {
		count = -1
	}
	TraceExitReturn("getDataSpaces", fmt.Sprintf("%v %v", keyword, count))
	return count
}
func br() {
	fmt.Printf("\n--------------------------------------------------------------------------------------------------------\n")
}

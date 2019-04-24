package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"gopkg.in/AlecAivazis/survey.v1"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var debug = false

type userCreds struct {
	username string
	password string
	realm    string
}

func PromptServer() string {
	TraceEnter("PromptServer")
	toReturn := server
	if server == "unset" {
		prompt := &survey.Input{Message: "APIC Manager or Admin Endpoint"}
		survey.AskOne(prompt, &toReturn, nil)
	}
	TraceExitReturn("PromptServer", toReturn)
	return toReturn
}
func PromptCredentials(porg string) userCreds {
	TraceEnter("PromptCredentials")
	if userDetails.username == "" {
		prompt := &survey.Input{Message: "Username"}
		survey.AskOne(prompt, &userDetails.username, nil)
	}
	if userDetails.password == "" {
		promptPass := &survey.Password{Message: "Password"}
		survey.AskOne(promptPass, &userDetails.password, nil)
	}
	if userDetails.realm == "" {
		prompt := &survey.Input{Message: "Realm"}
		survey.AskOne(prompt, &userDetails.realm, nil)
	}
	TraceExit("PromptCredentials")
	return userDetails
}
func Log(s interface{}) {
	log.Printf("\t %v\n", s)
}
func Trace(s interface{}) {
	if debug {
		log.Printf("\t %v\n", s)
	}
}
func TraceExit(s string) {
	if debug {
		log.Printf("------------------------------------\n")
		log.Printf("\t Exiting - %v\n", s)
		log.Printf("------------------------------------\n")
		log.Printf("------------------------------------\n")
	}
}
func TraceExitReturn(s string, i interface{}) {
	if debug {
		log.Printf("------------------------------------\n")
		log.Printf("\t Exiting - %v\n", s)
		log.Printf("------------\n")
		log.Printf("Returning - %v\n", i)
		log.Printf("------------------------------------\n")
		log.Printf("------------------------------------\n")
	}
}
func TraceEnter(s string) {
	if debug {
		log.Printf("------------------------------------\n")
		log.Printf("\t Entering - %v\n", s)
		log.Printf("------------------------------------\n")
	}
}

type gzreadCloser struct {
	*gzip.Reader
	io.Closer
}

func (gz gzreadCloser) Close() error {
	return gz.Closer.Close()
}

func APICRequest(path string) map[string]interface{} {
	TraceEnter("APICRequest")
	Trace("path " + "https://" + server + "/api/" + path)
	req, err := http.NewRequest("GET", "https://"+server+"/api/"+path, nil)
	if err != nil {
		Trace(err.Error())
		return nil
	}
	// Trace(token)
	req.Header.Set("Authorization", "Bearer "+token)
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Connection", "keep-alive")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		Trace(err.Error())
		return nil

	}
	if resp.Status != "200 OK" {
		log.Printf("HTTP Status '%v'", resp.Status)
		return nil
	}
	defer resp.Body.Close()

	jsonObj := make(map[string]interface{})
	bodyBytes, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		Trace(err2.Error())
		return nil
	}
	err = json.Unmarshal(bodyBytes, &jsonObj)
	if err != nil {
		Trace(err.Error())
		return nil
	}
	TraceExit("APICRequest")
	return jsonObj
}

func APICRequestSize(url string) int {
	TraceEnter("APICRequest")
	Trace("path " + url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		Trace(err.Error())
		return -1
	}
	// Trace(token)
	req.Header.Set("Authorization", "Bearer "+token)
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

	Trace(len(bodyBytes))

	return len(bodyBytes)
}

func login(UC userCreds) string {
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
		Username:     UC.username,
		Password:     UC.password,
		Realm:        UC.realm,
		ClientID:     "599b7aef-8841-4ee2-88a0-84d49c4d6ff2",
		ClientSecret: "0ea28423-e73b-47d4-b40e-ddb45c48bb0c",
		GrantType:    "password",
	}

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://"+server+"/api/token", body)
	if err != nil {
		// handle err
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
		return "error 1"
	}
	if resp.Status != "200 OK" {
		log.Printf("HTTP Status %v", resp.Status)
		Trace(string(payloadBytes))
		os.Exit(2)
		return "error 2"
	}

	defer resp.Body.Close()
	jsonObj := make(map[string]interface{})

	zr, err := gzip.NewReader(resp.Body)
	if err != nil {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		Trace(bodyBytes)
		Trace(err.Error())
	} else {
		resp.Body = gzreadCloser{zr, resp.Body}
	}

	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	// Trace(string(bodyBytes))
	if err2 != nil {
		Trace(err2.Error())
		return "error 3"
	}
	err = json.Unmarshal(bodyBytes, &jsonObj)

	if err != nil {
		Trace(err.Error())
		return "error 4"
	}
	TraceExit("login")
	return jsonObj["access_token"].(string)
}

func getData(type1 string, orgId string, keyword string) int {
	TraceEnter("getData")
	path := type1 + "/" + orgId + "/" + keyword
	jsonObj := APICRequest(path)
	if jsonObj == nil {
		return -1
	}
	toReturn := int(jsonObj["total_results"].(float64))
	TraceExitReturn("getData", fmt.Sprintf("keyword %v", toReturn))
	return toReturn
}

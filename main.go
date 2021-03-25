package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	createTaskURL = "https://api.capmonster.cloud/createTask"
)

type Solver struct {
	APIKey string
}

type NoCaptchaProxylessData struct {
	Type       string `json:"type"`
	WebsiteURL string `json:"websiteURL"`
	WebsiteKey string `json:"websiteKey"`
}

type NoCaptchaProxylessTask struct {
	ClientKey string                 `json:"clientKey"`
	Task      NoCaptchaProxylessData `json:"task"`
}


func (s Solver) RecaptchaV2(siteKey, url string) (recaptchaResponse string, err error) {
	payload := NoCaptchaProxylessData{Type: "NoCaptchaTaskProxyless", WebsiteURL: url, WebsiteKey: siteKey}
	task := NoCaptchaProxylessTask{ClientKey: s.APIKey, Task: payload}
	body, err := json.Marshal(task)
	log.Println(string(body))
	if err != nil {
		log.Fatalln("Failed to Marshal NoCaptchaProxylessTask")
		return "", err
	}
	resp, err := http.Post(createTaskURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalln("Error occured while using the capmonster api")
		return "", err
	}
	defer resp.Body.Close()
	// TODO
	/*
		Check Body For Errors
		Probably Refactor A Bit
		- error types
		- just better type systems
		Environment Variables for API Keys or Maybe A Config File
	*/
	if err != nil {
		log.Fatalln("Failed to read body from capmonster api")
		return "", err
	}
	return "", nil
}

func main() {
	s := Solver{APIKey: "api key here"}
	s.RecaptchaV2("6Le-wvkSAAAAAPBMRTvw0Q4Muexq9bi0DJwx_mJ-", "https://www.google.com/recaptcha/api2/demo")
}

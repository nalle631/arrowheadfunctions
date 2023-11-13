package arrowheadfunctions

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ProviderSystem struct {
	Address    string `json:"address"`
	Port       int    `json:"port"`
	SystemName string `json:"systemName"`
}

type RegServiceRequest struct {
	Interfaces        []string       `json:"interfaces"`
	ProviderSystem    ProviderSystem `json:"providerSystem"`
	ServiceDefinition string         `json:"serviceDefinition"`
	ServiceUri        string         `json:"serviceUri"`
}

func Hello() {
	fmt.Println("Daniel-sama")
}

func GetClient(certFile string, keyFile string) *http.Client {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Panic("Certficate load error. ", err)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}
	client := &http.Client{Transport: tr}
	return client
}

func PublishService(requestBody RegServiceRequest) {

	req := GetHTTPRequest("POST", "https://localhost:8443/serviceregistry/register", requestBody, "application/json")

	client := GetClient("./usercert.pem", "./userkey.pem")
	resp, err := client.Do(req)
	fmt.Println("request sent")
	fmt.Println("requestbody: ", requestBody)
	if err != nil {
		log.Panic("Error making HTTP request using client. ", err)
	}
	fmt.Println("## Response Body:\n", resp.Body)
	fmt.Println("## Response status:\n", resp.Status, resp.StatusCode)
}

func GetHTTPRequest(method string, url string, body RegServiceRequest, contentType string) *http.Request {
	payload, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", contentType)
	return req
}

func GetRegisterServiceRequestBody(interfaces []string, address string, port int, systemName string, serviceDefinition string, serviceUri string) RegServiceRequest {
	requestBody := new(RegServiceRequest)
	requestBody.Interfaces = interfaces
	requestBody.ProviderSystem.Address = address
	requestBody.ProviderSystem.Port = port
	requestBody.ProviderSystem.SystemName = systemName
	requestBody.ServiceDefinition = serviceDefinition
	requestBody.ServiceUri = serviceUri
	return *requestBody
}

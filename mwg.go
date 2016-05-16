package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"crypto/tls"
	"flag"
	"net/http/cookiejar"
	"encoding/xml"
)

type RetVal struct {
	Entry xml.Name `xml:"entry"`
	Content string `xml:"content"`
}

func auth(host string, ignoreSSLCert bool, userName string, userPass string, commandStr string) (content string){
	url := "https://" + host + ":4712"

	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	if ignoreSSLCert {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client = &http.Client{Jar: cookieJar,Transport: tr}
	}

	req, _ := http.NewRequest("POST", url + "/Konfigurator/REST/" + commandStr, nil)
	//req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		log.Fatal(err)
	}

	var v RetVal

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	err = xml.Unmarshal([]byte(resp_body), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	//fmt.Println(resp.Status)
	return v.Content
}

func callCommand(host string, cookie string, ignoreSSLCert bool, httpMethod string, commandStr string){
	url := "https://" + host + ":4712"

	cookieJar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: cookieJar,
	}

	if ignoreSSLCert {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}

		client = &http.Client{Transport: tr}
	}

	req, _ := http.NewRequest(httpMethod, url + "/Konfigurator/REST/" + commandStr, nil)
	req.Header.Set("Cookie", "JSESSIONID="+cookie)
	//req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		log.Fatal(err)
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

}

func main() {

	// Check Flags and use them
	var ignoreSSLCertFlag = flag.Bool("ignoressl",false,"Ignore Bad SSL Certificates")
	var userNameFlag = flag.String("user","","API User Name")
	var userPassFlag = flag.String("pass", "", "API Password")
	var hostFlag = flag.String("host", "localhost", "Host IP Address or FQDN")

	flag.Parse()

	ignoreSSLCert := *ignoreSSLCertFlag
	userName := *userNameFlag
	userPass := *userPassFlag
	host := *hostFlag

	loginStr := "login?userName=" + userName + "&pass=" + userPass
	clientCookie := auth(host, ignoreSSLCert, userName,userPass, loginStr)

	callCommand(host, clientCookie, ignoreSSLCert,"GET", "appliances")

	callCommand(host, clientCookie, ignoreSSLCert,"POST", "logout")
}

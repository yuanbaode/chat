package models

import (
	"testing"
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"net"
)

func TestRandomData(t *testing.T) {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Duration(5) * time.Second)
}

func TestRandomData12(t *testing.T) {
	apiUrl := `https://api.erong28.com/api/lottery/latest/period/short`
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return
	}
	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		MaxIdleConns: 2,
	}
	client := &http.Client{
		Timeout:   time.Second * 5,
		Transport: transport,
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	if resp == nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}
func TestXXX(t *testing.T){
	apiUrl := `https://api.erong28.com/api/lottery/latest/period/short`
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(1)
		return
	}
	if resp == nil {
		fmt.Println(2)
		return
	}
	if resp.StatusCode != 200 {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	_=body
}

func TestInfoStored_TableName(t *testing.T) {
	for {
		sleepTime := getRandNum(2)
		fmt.Println(sleepTime)
		time.Sleep(time.Duration(sleepTime)*time.Second)
	}
}
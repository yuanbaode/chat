package models

import (
	"testing"
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
)

func TestRandomData(t *testing.T) {
	fmt.Println(time.Now().Unix())
	//RandomData(nil)
}


func TestXXX(t *testing.T){
	apiUrl := `http://test.lxh.wiki/api/lottery/latest/period`
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

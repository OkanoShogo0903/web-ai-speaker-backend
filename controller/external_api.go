package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string) (*string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Get err")
		return nil, err
	}
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Http read err")
		return nil, err
	}

	var wikis []Wikipedia
	err = json.Unmarshal(byteArray, &wikis)
	if err != nil {
		fmt.Println("Json unmarshal err")
		return nil, err
	}
	//fmt.Println(wikis["Body"])
	for i := range wikis {
		fmt.Println(wikis[i].Body)
	}

	if len(wikis) > 0 {
		fmt.Println("Empty err")
		return &wikis[0].Body, nil
	}
	return nil, nil
}

func request2WordApi(word *string) (*string, error) {
	//fmt.Println(*word)
	url := "http://wikipedia.simpleapi.net/api?keyword=" + *word + "&output=json"

	url = strings.Join(strings.Fields(url), "")
	res, err := Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

package controller

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Wikipedia struct {
	Body string
}

func wikipediaGet(url string) (*string, error) {
	byteArray, err := Get(url)

	var wikis []Wikipedia
	err = json.Unmarshal(byteArray, &wikis)
	if err != nil {
		fmt.Println("Json unmarshal err")
		return nil, err
	}
	for i := range wikis {
		fmt.Println(wikis[i].Body)
	}

	if len(wikis) > 0 {
		return &wikis[0].Body, nil
	}

	fmt.Println("Empty err")
	return nil, nil
}

func request2WordApi(word *string) (*string, error) {
	url := "http://wikipedia.simpleapi.net/api?keyword=" + *word + "&output=json"

	url = strings.Join(strings.Fields(url), "")
	res, err := wikipediaGet(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

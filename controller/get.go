package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Get err")
		return nil, err
	}
	byteArray, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Http read err")
		return nil, err
	}
	return byteArray, err
}

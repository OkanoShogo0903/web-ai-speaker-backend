package controller

import (
	"errors"
	"strings"
)

func trimText(input_text string, wake string) (*string, error) {
	t := strings.Replace(input_text, wake, "", 1)
	// Remove space
	t = strings.Replace(t, " ", "", -1)
	b := t
	// Remove noice
	keywords := []string{
		"を調べて",
		"で調べて",
		"を検索",
		"で検索",
		"について",
		"をついて",
		"調べて",
		"検索",
	}
	for _, k := range keywords {
		t = strings.Replace(t, k, "", -1)
	}
	if b == t {
		var err = errors.New("no keyword")
		return nil, err
	}
	return &t, nil
}

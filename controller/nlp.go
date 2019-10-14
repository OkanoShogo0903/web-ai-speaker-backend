package controller

import "strings"

func trimText(input_text string, wake string) string {
	t := strings.Replace(input_text, wake, "", 1)
	// Remove space
	t = strings.Replace(t, " ", "", -1)
	// Remove noice
	ng_list := []string{
		"を調べて",
		"で調べて",
		"を検索",
		"で検索",
		"について",
		"をついて",
		"調べて",
		"検索",
	}
	for _, ng := range ng_list {
		t = strings.Replace(t, ng, "", -1)
	}
	return t
}

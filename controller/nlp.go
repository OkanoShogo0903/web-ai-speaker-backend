package controller

import (
	"errors"
	"strings"

	"github.com/OkanoShogo0903/web-ai-speaker-backend/model"
)

var search_keywords = []string{
	"を調べて",
	"で調べて",
	"を検索",
	"で検索",
	"について",
	"をついて",
	"調べて",
	"検索",
}

var bgm_keywords = []string{
	"を再生して",
	"を流して",
	"を再生",
}

func KeywordCheck(input_text string) (*string, *model.State, error) {
	// Remove space
	t := strings.Replace(input_text, " ", "", -1)

	// Check keyword
	save := t
	for _, k := range search_keywords {
		t = strings.Replace(t, k, "", -1)
	}
	if save != t {
		m := model.Search
		return &t, &m, nil
	}

	for _, k := range bgm_keywords {
		t = strings.Replace(t, k, "", -1)
	}
	if save != t {
		m := model.Bgm
		return &t, &m, nil
	}

	// No keyword
	return nil, nil, errors.New("no keyword")
}

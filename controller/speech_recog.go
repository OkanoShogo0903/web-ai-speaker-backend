package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/OkanoShogo0903/web-ai-speaker/backend/model"

	"github.com/gin-gonic/gin"
)

type SpeechPostRequest struct {
	Text string `json:"text"`
}

type Wikipedia struct {
	Body string
}

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
	fmt.Println(url)
	res, err := Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func SpeechPost(r *model.SpeechResult) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: CORSについて、どれだけの範囲を許可するかを決める、またはHelpに投稿してみる。
		fmt.Println("SpeechPost!!!")

		if err := c.BindJSON(&r); err != nil {
			panic(err)
		}

		// WakeWordの判定
		wake := []rune("ハローワールド")
		if strings.HasPrefix(r.Text, string(wake)) == false {
			c.JSON(http.StatusNoContent, gin.H{"text": "wake word is not included"})
			return
		}

		// 入力テキストの整形
		want_search := trimText(r.Text, string(wake))
		if utf8.RuneCountInString(want_search) == 0 {
			c.JSON(http.StatusNoContent, gin.H{"text": "input ward have no body"})
			return
		}

		fmt.Println(r.Text)
		fmt.Println(want_search)

		mean, err := request2WordApi(&want_search)

		if err != nil {
			fmt.Printf("%+v", err)
			//c.JSON(http.StatusOK, gin.H{"text": "err nil"})
			c.JSON(http.StatusInternalServerError, gin.H{"text": "err"})
			return
		}
		if mean == nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusOK, gin.H{"text": "検索候補が見つかりませんでした", "question": want_search})
			return
		}

		c.JSON(http.StatusOK, gin.H{"text": mean, "question": want_search})
	}
}

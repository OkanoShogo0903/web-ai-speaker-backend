package controller

import (
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/OkanoShogo0903/web-ai-speaker-backend/model"
	"github.com/gin-gonic/gin"
)

type SpeechPostRequest struct {
	Text string `json:"text"`
}

type Wikipedia struct {
	Body string
}

func SpeechPost(r *model.SpeechResult) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.BindJSON(&r); err != nil {
			panic(err)
		}

		// WakeWordの判定
		wake := []rune("ハローワールド")
		if strings.HasPrefix(r.Text, string(wake)) == false {
			c.JSON(http.StatusNoContent, gin.H{"text": r.Text}) // Wake word is not included
			return
		}

		// 入力テキストの整形
		want_search := trimText(r.Text, string(wake))
		if utf8.RuneCountInString(want_search) == 0 {
			c.JSON(http.StatusNoContent, gin.H{"text": r.Text}) // "input ward have no body"
			return
		}

		fmt.Println(r.Text)
		fmt.Println(want_search)

		mean, err := request2WordApi(&want_search)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"text": "no text"})
			return
		}
		if mean == nil {
			fmt.Printf("%+v", err)
			c.JSON(210, gin.H{"text": "検索候補が見つかりませんでした", "question": want_search})
			return
		}

		c.JSON(http.StatusOK, gin.H{"text": mean, "question": want_search})
	}
}

package controller

import (
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/OkanoShogo0903/web-ai-speaker-backend/model"
	"github.com/gin-gonic/gin"
)

type SpeechPostRequest struct {
	Text string `json:"text"`
}

func SpeechPost(r *model.SpeechResult) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.BindJSON(&r); err != nil {
			panic(err)
		}

		// WakeWordの判定
		/*
			wake := []rune("") // ハロー
			if strings.HasPrefix(r.Text, string(wake)) == false {
				c.JSON(http.StatusNoContent, gin.H{"text": r.Text}) // Wake word is not included
				return
			}
		*/

		// 入力テキストの整形
		want_search, state, err := KeywordCheck(r.Text)
		if err != nil {
			fmt.Printf("%+v", err)
			c.JSON(http.StatusNoContent, gin.H{"text": r.Text}) // Input ward have no keyword
			return
		}
		if utf8.RuneCountInString(*want_search) == 0 {
			c.JSON(http.StatusNoContent, gin.H{"text": r.Text}) // Input ward have no body
			return
		}

		fmt.Println(r.Text)
		fmt.Println(*want_search)

		switch *state {
		case model.Search:
			mean, err := request2WordApi(want_search)
			if err != nil {
				fmt.Printf("%+v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"type": model.Search, "text": "no text"})
				return
			}
			if mean == nil {
				fmt.Printf("%+v", err)
				c.JSON(210, gin.H{"type": model.Search, "title": *want_search, "text": "検索候補が見つかりませんでした"})
				return
			}

			c.JSON(http.StatusOK, gin.H{"type": model.Search, "title": *want_search, "text": mean})
		case model.Bgm:
			title, videoid, err := request2YoutubeApi(want_search)
			if err != nil {
				fmt.Printf("%+v", err)
				c.JSON(http.StatusNoContent, gin.H{"type": model.Bgm, "text": "no video"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"type": model.Bgm, "title": *title, "videoid": *videoid})

		}
	}
}

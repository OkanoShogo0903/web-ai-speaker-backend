package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Youtube struct {
	Kind          string `json:"kind"`
	Etag          string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	RegionCode    string `json:"regionCode"`
	PageInfo      struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
	Items []struct {
		Kind string `json:"kind"`
		Etag string `json:"etag"`
		ID   struct {
			Kind    string `json:"kind"`
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`
			ChannelTitle         string `json:"channelTitle"`
			LiveBroadcastContent string `json:"liveBroadcastContent"`
		} `json:"snippet"`
	} `json:"items"`
}

func youtubeGet(url string) (*string, *string, error) {
	fmt.Println(url)
	byteArray, err := Get(url)

	var youtube Youtube
	err = json.Unmarshal(byteArray, &youtube)
	if err != nil {
		fmt.Println("Json unmarshal err")
		return nil, nil, err
	}

	if len(youtube.Items) == 0 {
		return nil, nil, errors.New("no video")
	}
	return &youtube.Items[0].ID.VideoID, &youtube.Items[0].Snippet.Title, nil
}

func request2YoutubeApi(word *string) (*string, *string, error) {
	api_key := "AIzaSyBEqWv7uVOhJmLxsC3YkFPu-efYqnmaKAA" // 一般情報しかアクセスできないので晒してもダメージ少ない
	url := "https://www.googleapis.com/youtube/v3/search?type=video&part=snippet&q=" + *word + "&key=" + api_key

	title, videoid, err := youtubeGet(url)
	if err != nil {
		return nil, nil, err
	}
	return title, videoid, nil
}

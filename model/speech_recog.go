package model

type SpeechResult struct {
	Text string `json:"text"`
}

func New() *SpeechResult {
	return &SpeechResult{}
}

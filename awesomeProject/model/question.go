package model

type Question struct {
	Qid      int    `json:"qid"`
	Question string `json:"question"`
}

type QuestionAnother struct {
	Question string `json:"question"`
}

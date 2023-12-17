package main

import (
	"encoding/json"
	"net/http"

	"github.com/rep-co/hola-quiz-backend/hola-quiz-api/remote/quizapi"
)

var client *http.Client

func GetQuizQuestions(tags, limit, difficulty string) ([]*quizapi.QuestionResult, error) {
	quizApiUrl := quizapi.NewQuizApiUrl(
		cfg.ApiKey,
		tags,
		limit,
		difficulty)
	url := quizApiUrl.GetUrl()

	questionResult := []*quizapi.QuestionResult{}
	err := GetJson(url, &questionResult)
	if err != nil {
		return nil, err
	}

	return questionResult, nil
}

func GetJson(url string, target interface{}) error {
	res, err := client.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

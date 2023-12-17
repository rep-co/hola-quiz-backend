package quizapi

type QuizApiUrl struct {
	Base       string
	ApiKey     string
	Tags       string
	Limit      string
	Difficulty string
}

func NewQuizApiUrl(apiKey, tags, limit, difficulty string) *QuizApiUrl {
	return &QuizApiUrl{
		Base:       "https://quizapi.io/api/v1/questions",
		ApiKey:     apiKey,
		Tags:       tags,
		Limit:      limit,
		Difficulty: difficulty,
	}
}

func (q *QuizApiUrl) GetUrl() string {
	return q.Base +
		"?" + "apiKey=" + q.ApiKey +
		"&" + "tags=" + q.Tags +
		"&" + "limit=" + q.Limit +
		"&" + "difficulty=" + q.Difficulty
}

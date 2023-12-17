package quizapi

type QuestionResult struct {
	ID                     int                    `json:"id"`
	Question               string                 `json:"question"`
	Description            string                 `json:"description"`
	Answers                QuestionAnswers        `json:"answers"`
	MultipleCorrectAnswers bool                   `json:"multiple_correct_answers,string"`
	CorrectAnswers         QuestionCorrectAnswers `json:"correct_answers"`
	Explanation            string                 `json:"explanation"`
}

type QuestionAnswers struct {
	AnswerA string `json:"answer_a"`
	AnswerB string `json:"answer_b"`
	AnswerC string `json:"answer_c"`
	AnswerD string `json:"answer_d"`
	AnswerE string `json:"answer_e"`
	AnswerF string `json:"answer_f"`
}

func (q *QuestionAnswers) ConvertToSlice() []string {
	answers := []string{
		q.AnswerA,
		q.AnswerB,
		q.AnswerC,
		q.AnswerD,
		q.AnswerE,
		q.AnswerF,
	}
	return answers
}

func (q *QuestionCorrectAnswers) ConvertToSlice() []bool {
	correctAnswers := []bool{
		q.AnswerA,
		q.AnswerB,
		q.AnswerC,
		q.AnswerD,
		q.AnswerE,
		q.AnswerF,
	}
	return correctAnswers
}

type QuestionCorrectAnswers struct {
	AnswerA bool `json:"answer_a_correct,string"`
	AnswerB bool `json:"answer_b_correct,string"`
	AnswerC bool `json:"answer_c_correct,string"`
	AnswerD bool `json:"answer_d_correct,string"`
	AnswerE bool `json:"answer_e_correct,string"`
	AnswerF bool `json:"answer_f_correct,string"`
}

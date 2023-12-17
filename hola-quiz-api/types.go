package main

type CreatePackRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Emojis      string `json:"emojis"`
}

type Pack struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Emojis      string `json:"emojis"`
}

func NewPack(name, description, category string, emojis string) *Pack {
	return &Pack{
		Name:        name,
		Description: description,
		Category:    category,
		Emojis:      emojis,
	}
}

type PackQuiz struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Category    string          `json:"category"`
	Questions   []*QuizQuestion `json:"questions"`
}

func NewPackQuiz(
	pack *Pack,
	questions []*QuizQuestion,
) *PackQuiz {
	return &PackQuiz{
		ID:          pack.ID,
		Name:        pack.Name,
		Description: pack.Description,
		Category:    pack.Category,
		Questions:   questions,
	}
}

type PackPreview struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

func ConvertPackToPackPreview(
	id int,
	name,
	description,
	category string,
) *PackPreview {
	return &PackPreview{
		ID:          id,
		Name:        name,
		Description: description,
		Category:    category,
	}
}

func ConvertPacksToPackPreviews(packs []*Pack) []*PackPreview {
	packPreviews := []*PackPreview{}
	for _, element := range packs {
		packPreview := ConvertPackToPackPreview(
			element.ID,
			element.Name,
			element.Description,
			element.Description)
		packPreviews = append(packPreviews, packPreview)
	}

	return packPreviews
}

func ConvertPackToPackQuiz(pack *Pack, questions []*QuizQuestion) *PackQuiz {
	return &PackQuiz{
		ID:          pack.ID,
		Name:        pack.Name,
		Description: pack.Description,
		Category:    pack.Category,
		Questions:   questions,
	}
}

type QuizQuestion struct {
	ID                     int      `json:"id"`
	Question               string   `json:"question"`
	Description            string   `json:"description"`
	Answers                []string `json:"answers"`
	AnswersCorrectness     []bool   `json:"correct_answers"`
	MultipleCorrectAnswers bool     `json:"multiple_correct_answers"`
	Explanation            string   `json:"explanation"`
}

func NewQuizQuestion(
	id int,
	question string,
	description string,
	answers []string,
	answersCorrectness []bool,
	multipleCorrectAnswers bool,
	explanation string,
) *QuizQuestion {
	return &QuizQuestion{
		ID:                     id,
		Question:               question,
		Description:            description,
		Answers:                answers,
		AnswersCorrectness:     answersCorrectness,
		MultipleCorrectAnswers: multipleCorrectAnswers,
		Explanation:            explanation,
	}
}

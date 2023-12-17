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

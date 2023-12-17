package main

type CreatePackRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type Pack struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

func NewPack(name, description, category string) *Pack {
	return &Pack{
		Name:        name,
		Description: description,
		Category:    category,
	}
}

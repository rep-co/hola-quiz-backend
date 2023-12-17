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

type PackPreview struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

func NewPackPreview(name, description, category string) *PackPreview {
	return &PackPreview{
		Name:        name,
		Description: description,
		Category:    category,
	}
}

func ConvertPacksToPackPreviews(packs []*Pack) []*PackPreview {
	packPreviews := []*PackPreview{}
	for _, element := range packs {
		packPreview := NewPackPreview(
			element.Name,
			element.Description,
			element.Description)
		packPreviews = append(packPreviews, packPreview)
	}

	return packPreviews
}

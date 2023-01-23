package models

import (
	"encoding/json"
	"os"
)

const CategoryFileName = "_category_.json"

type CategoryLink struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Category struct {
	Label    string        `json:"label"`
	Position uint8         `json:"position"`
	Link     *CategoryLink `json:"link"`
}

// WriteCategoryFile writes a category file to the given path.
func WriteCategoryFile(filePath string, label string,
	description string) error {

	category := &Category{
		Label:    label,
		Position: 0,
		Link: &CategoryLink{
			Type:        "generated-index",
			Description: description,
		},
	}
	finalBytes, err := json.MarshalIndent(category, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, finalBytes, 0o644)
	if err != nil {
		return err
	}
	return nil
}

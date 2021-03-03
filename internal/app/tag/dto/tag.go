package dto

// Tag ...
type TagDTO struct {
	Name   string `json:"name" validate:"max=50"`
	Colour string `json:"colour" validate:"max=7"`
}

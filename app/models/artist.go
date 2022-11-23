package models

type Artist struct {
	Name         string `json:"name"`
	Painting     string `json:"paint"`
	AtExhibition bool   `json:"at_exhibition"`
}

func (a *Artist) DeletePaint() string {
	return a.Painting
}

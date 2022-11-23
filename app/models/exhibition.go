package models

import (
	"fmt"
	"time"
)

func NewExhibition(id int) *Exhibition {
	return &Exhibition{
		id:        id,
		Name:      fmt.Sprintf("exhibition %s", id),
		StartDate: time.Now(),
	}
}

type Exhibition struct {
	id        int
	Name      string    `json:"name"`
	Artists   []*Artist `json:"artists"`
	StartDate time.Time `json:"startDate"`
}

func (e Exhibition) Id() int {
	return e.id
}
func (e Exhibition) AddArtist(a *Artist) {
	e.Artists = append(e.Artists, a)
}

func (e *Exhibition) DeleteArtist(a *Artist) error {
	for i, artist := range e.Artists {
		if artist.DeletePaint() == a.DeletePaint() {
			if len(e.Artists) == 1 {
				e.Artists = []*Artist{}
			} else {
				e.Artists = e.Artists[i:len(e.Artists)]

			}
			return nil
		}
	}
	return fmt.Errorf("artist %s not found", a.DeletePaint())
}

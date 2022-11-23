package controllers

import (
	"fmt"
	"mvc/app/models"
)

type ExhibitionController struct {
	exhibitions []*models.Exhibition
}

func (ec *ExhibitionController) OrganizeExhibition(e *models.Exhibition) {
	ec.exhibitions = append(ec.exhibitions, e)
}
func (ec *ExhibitionController) AddArtist(cid int, a *models.Artist) *models.Exhibition {
	for _, exhibition := range ec.exhibitions {
		if exhibition.Id() == cid {
			exhibition.AddArtist(a)
			return exhibition
		}
	}
	return nil
}
func (ec *ExhibitionController) DeleteArtist(eid int, a *models.Artist) {
	for _, exhibition := range ec.exhibitions {
		if exhibition.Id() == eid {
			if err := exhibition.DeleteArtist(a); err != nil {
				fmt.Println(err)
			}
		}
	}
}

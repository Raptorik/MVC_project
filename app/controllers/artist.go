package controllers

import "mvc/app/models"

type ArtistController struct {
	artists []*models.Artist
}

func (ac *ArtistController) InviteArtist(a *models.Artist) {
	ac.artists = append(ac.artists, a)
}

func (ac *ArtistController) FindArtist(paint string) *models.Artist {
	for _, user := range ac.artists {
		if user.DeletePaint() == paint {
			{
				return user
			}
			return user
		}
	}
	return nil
}

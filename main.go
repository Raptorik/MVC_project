// Rough brush stroke
package main

import (
	"fmt"
	"mvc/app/controllers"
	"mvc/app/models"
	"mvc/app/views"
)

func main() {
	var eid int
	ac := controllers.ArtistController{}
	ec := controllers.ExhibitionController{}

Text:
	fmt.Printf("CHOOSE WHAT TO DO: " +
		"\nInvite artist: ia" +
		"\nOrganize exhibition: oe" +
		"\nAdd artist's paintings to the gallery: ep" +
		"\nDelete artist's paintings from the gallery: dp" +
		"\nExit: e\n")
	var command string
	_, _ = fmt.Scan(&command)

	switch command {
	case `ia`:
		a := &models.Artist{`Vanya`, `black box`, false}
		ac.InviteArtist(a)
		fmt.Println(" ----- Artist invited")
		if err := views.PrintArtist(a); err != nil {
			fmt.Println(err)
		}
		goto Text
	case `oe`:
		eid++
		e := models.NewExhibition(eid)
		ec.OrganizeExhibition(e)
		fmt.Println(" ----- Exhibition opened")
		if err := views.PrintExhibition(e); err != nil {
			fmt.Println(err)
		}
		goto Text
	case `ep`:
		fmt.Printf("<artist's painting>?")
		var artistPainting string
		_, _ = fmt.Scan(&artistPainting)
		artist2add := ac.FindArtist(artistPainting)
		if artist2add == nil {
			fmt.Printf("Painting of %s not found", artistPainting)
			goto Text
		}
		fmt.Printf("<exhibition id>?")
		var exhibitionId int
		_, _ = fmt.Scan(&exhibitionId)
		exhibition := ec.AddArtist(exhibitionId, artist2add)
		fmt.Println(" ----- Artist added")
		if err := views.PrintExhibition(exhibition); err != nil {
			fmt.Println(err)
		}
		goto Text
	case `dp`:
		fmt.Printf("<Artist's painting>?")
		var artistPainting string
		_, _ = fmt.Scan(&artistPainting)
		artist2delete := ac.FindArtist(artistPainting)
		if artist2delete == nil {
			fmt.Printf("User with %s not found", artistPainting)
			goto Text
		}
		fmt.Printf("<course id>?")
		var exhibitionId int
		_, _ = fmt.Scan(&exhibitionId)
		ec.DeleteArtist(exhibitionId, artist2delete)
		fmt.Println(" ----- Painting taken off the exhibition")
		goto Text
	case `e`:
		break
	}

	fmt.Println("Closed by hoster")
}

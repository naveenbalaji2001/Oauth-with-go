package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naveenbalaji2001/Oauth-with-go/dao"
	"github.com/naveenbalaji2001/Oauth-with-go/model"
	"github.com/naveenbalaji2001/Oauth-with-go/service"
)


type TrackController struct {
	DatabaseAccessor *dao.DatabaseAccessor
	SpotifyService   *service.SpotifyService
}

// initializes a new TrackController
func NewTrackController(dbAccessor *dao.DatabaseAccessor, spotifyService *service.SpotifyService) *TrackController {
	return &TrackController{DatabaseAccessor: dbAccessor, SpotifyService: spotifyService}
}

//get track details by ISRC code
func (tc *TrackController) GetTrackDetailsByISRC(c *gin.Context) {
	isrc := c.Param("isrc")
	existingTrack, err := tc.DatabaseAccessor.GetTrackByISRC(isrc)

	if err == nil {
		// If the track is found in the database, return it
		trackDetails := model.TrackDetails{
			ISRC:         existingTrack.ISRC,
			Title:        existingTrack.Title,
			ArtistName:   existingTrack.ArtistName,
			SpotifyImage: existingTrack.SpotifyImage,
		}
		c.JSON(http.StatusOK, trackDetails)
		return
	}

	// otherwise, proceed to search in Spotify
	track, err := tc.SpotifyService.SearchTrackByISRC(isrc)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Track not found"})
		return
	}

	trackDetails := model.TrackDetails{
		ISRC:         isrc,
		Title:        track.Name,
		ArtistName:   track.Artists[0].Name,
		SpotifyImage: track.Album.Images[0].URL,
	}

	// Create a new Track record in the database
	newTrack := model.Track{
		ISRC:         isrc,
		Title:        trackDetails.Title,
		ArtistName:   trackDetails.ArtistName,
		SpotifyImage: trackDetails.SpotifyImage,
	}
	tc.DatabaseAccessor.SaveTrack(&newTrack)

	c.JSON(http.StatusOK, newTrack)
}

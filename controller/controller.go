package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naveenbalaji2001/Oauth-with-go/dao"
	"github.com/naveenbalaji2001/Oauth-with-go/model"
	"github.com/naveenbalaji2001/Oauth-with-go/service"
)

// It handles HTTP requests related to tracks
type TrackController struct {
	DatabaseAccessor *dao.DatabaseAccessor
	SpotifyService   *service.SpotifyService
}

func NewTrackController(dbAccessor *dao.DatabaseAccessor, spotifyService *service.SpotifyService) *TrackController {
	return &TrackController{DatabaseAccessor: dbAccessor, SpotifyService: spotifyService}
}


// @Summary Get track details by ISRC
// @Description Get track details from the database or Spotify by ISRC code
// @ID get-track-by-isrc
// @Produce json
// @Param isrc path string true "ISRC code of the track"
// @Success 200 {object} model.TrackDetails
// @Failure 404 {object} map[string]interface{} "Track not found"
// @Router /track/{isrc} [get]	
func (tc *TrackController) GetTrackDetailsByISRC(c *gin.Context) {
	isrc := c.Param("isrc")
	existingTrack, err := tc.DatabaseAccessor.GetTrackByISRC(isrc)

	if err == nil {
		// If track is found in the database, return it
		trackDetails := model.TrackDetails{
			ISRC:         existingTrack.ISRC,
			Title:        existingTrack.Title,
			ArtistName:   existingTrack.ArtistName,
			SpotifyImage: existingTrack.SpotifyImage,
		}
		c.JSON(http.StatusOK, trackDetails)
		return
	}

	// If Track not found in database then proceed to search in Spotify
	track, err := tc.SpotifyService.SearchTrackByISRC(isrc)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Track not found"})
		return
	}

	// Extract track details from the spotify search result
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

	// Return track details as JSON
	c.JSON(http.StatusOK, newTrack)
}

// @Summary Create a new track
// @Description Create a new track record in the database
// @ID create-track
// @Accept json
// @Produce json
// @Param trackDetails body model.TrackDetails true "Track details to create"
// @Success 201 {object} model.TrackDetails
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 409 {object} map[string]interface{} "Track with ISRC code already exists"
// @Router /track [post]
func (tc *TrackController) CreateTrack(c *gin.Context) {
	var trackDetails model.TrackDetails
	if err := c.BindJSON(&trackDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Checking if the track with given ISRC code is already exist
	if err, _ := tc.DatabaseAccessor.GetTrackByISRC(trackDetails.ISRC); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Track with ISRC code already exists"})
		return
	}

	newTrack := model.Track{
		ISRC:         trackDetails.ISRC,
		Title:        trackDetails.Title,
		ArtistName:   trackDetails.ArtistName,
		SpotifyImage: trackDetails.SpotifyImage,
	}
	tc.DatabaseAccessor.SaveTrack(&newTrack)

	c.JSON(http.StatusCreated, gin.H{"message": "Track record created successfully"})
}

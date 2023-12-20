package controller

import (
	"fmt"
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

	updatedTrack, err := tc.DatabaseAccessor.GetTrackByISRC(isrc)
	if err != nil {
		// If the track is not found in the database, proceed to search in Spotify
		track, err := tc.SpotifyService.SearchTrackByISRC(isrc)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Track not found"})
			return
		}

		// Extract track details from the Spotify search result
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
		return
	}

	// If the track is found in the database after the update, return its details
	trackDetails := model.TrackDetails{
		ISRC:         updatedTrack.ISRC,
		Title:        updatedTrack.Title,
		ArtistName:   updatedTrack.ArtistName,
		SpotifyImage: updatedTrack.SpotifyImage,
	}
	c.JSON(http.StatusOK, trackDetails)
}

// @Summary Search for tracks by artist name
// @Description Search for tracks by artist name using the Spotify API
// @ID search-track-by-artist
// @Produce json
// @Param artist_name path string true "Name of the artist"
// @Success 200 {array} model.TrackDetails
// @Failure 404 {object} map[string]interface{} "tracks not found"
// @Router /track/artist/{artist_name} [get]
func (tc *TrackController) SearchTrackByArtist(c *gin.Context) {
	artistName := c.Param("artist_name")
	tracks, err := tc.SpotifyService.SearchTrackByArtist(artistName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tracks not found"})
		return
	}
	c.JSON(http.StatusOK, tracks)
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

// @Summary Update a track by ISRC
// @Description Update an existing track record in the database by ISRC
// @Accept json
// @Produce json
// @Param isrc path string true "ISRC code of the track to be updated"
// @Param trackDetails body model.TrackDetails true "Updated track details"
// @Success 200 {object} model.TrackDetails
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /track/{isrc} [put]
func (ts *TrackController) UpdateTrack(isrc string, updatedTrackDetails *model.TrackDetails) (*model.TrackDetails, error) {
	
	existingTrack, err := ts.DatabaseAccessor.GetTrackByISRC(isrc)
	if err != nil {
		return nil, fmt.Errorf("track not found")
	}

	existingTrack.Title = updatedTrackDetails.Title
	existingTrack.ArtistName = updatedTrackDetails.ArtistName
	existingTrack.SpotifyImage = updatedTrackDetails.SpotifyImage

	if err := ts.DatabaseAccessor.DB.Save(existingTrack).Error; err != nil {
		return nil, fmt.Errorf("failed to update track in the database")
	}

	return &model.TrackDetails{
		ISRC:         existingTrack.ISRC,
		Title:        existingTrack.Title,
		ArtistName:   existingTrack.ArtistName,
		SpotifyImage: existingTrack.SpotifyImage,
	}, nil
}


func (tc *TrackController) UpdateTheTrack(c *gin.Context) {
	isrc := c.Param("isrc")
		var updatedTrackDetails model.TrackDetails
		if err := c.ShouldBindJSON(&updatedTrackDetails); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
	
		trackDetails, err := tc.UpdateTrack(isrc, &updatedTrackDetails)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
	
		c.JSON(200, trackDetails)
}

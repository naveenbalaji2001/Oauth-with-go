// service.go
package service

import (
	"fmt"
	"github.com/zmb3/spotify"
	
)

type SpotifyService struct {
	SpotifyClient *spotify.Client
}

// NewSpotifyService initializes a new SpotifyService
func NewSpotifyService(client *spotify.Client) *SpotifyService {
	return &SpotifyService{SpotifyClient: client}
}

// @Summary Search for a track by ISRC code
// @Description Search for a track by ISRC code using the Spotify API
// @ID search-track-by-isrc
// @Produce json
// @Param isrc path string true "ISRC code of the track"
// @Success 200 {object} spotify.FullTrack
// @Failure 404 {object} map[string]interface{} "track not found"
// @Router /search/track/{isrc} [get]
// SearchTrackByISRC searches for a track by ISRC code using the Spotify API

func (s *SpotifyService) SearchTrackByISRC(isrc string) (*spotify.FullTrack, error) {
	query := fmt.Sprintf("isrc:%s", isrc)
	searchResult, err := s.SpotifyClient.Search(query, spotify.SearchTypeTrack)
	if err != nil || len(searchResult.Tracks.Tracks) == 0 {
		return nil, fmt.Errorf("track not found")
	}
	return &searchResult.Tracks.Tracks[0], nil
}

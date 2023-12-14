// service.go
package service

import (
	"fmt"
	"github.com/zmb3/spotify"
	
)

type SpotifyService struct {
	SpotifyClient *spotify.Client
}

//initializes a new SpotifyService
func NewSpotifyService(client *spotify.Client) *SpotifyService {
	return &SpotifyService{SpotifyClient: client}
}

// SearchTrackByISRC using the Spotify API
func (s *SpotifyService) SearchTrackByISRC(isrc string) (*spotify.FullTrack, error) {
	query := fmt.Sprintf("isrc:%s", isrc)
	searchResult, err := s.SpotifyClient.Search(query, spotify.SearchTypeTrack)
	if err != nil || len(searchResult.Tracks.Tracks) == 0 {
		return nil, fmt.Errorf("track not found")
	}
	return &searchResult.Tracks.Tracks[0], nil
}

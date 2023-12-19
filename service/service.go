package service

import (
	"fmt"
	"github.com/naveenbalaji2001/Oauth-with-go/model"
	"github.com/zmb3/spotify"
)

type SpotifyService struct {
	SpotifyClient *spotify.Client
}


func NewSpotifyService(client *spotify.Client) *SpotifyService {
	return &SpotifyService{SpotifyClient: client}
}


func (s *SpotifyService) SearchTrackByISRC(isrc string) (*spotify.FullTrack, error) {
	query := fmt.Sprintf("isrc:%s", isrc)
	searchResult, err := s.SpotifyClient.Search(query, spotify.SearchTypeTrack)
	if err != nil || len(searchResult.Tracks.Tracks) == 0 {
		return nil, fmt.Errorf("track not found")
	}
	return &searchResult.Tracks.Tracks[0], nil
}

// @Summary Search for tracks by artist name
// @Description Search for tracks by artist name using the Spotify API
// @ID search-track-by-artist
// @Produce json
// @Param artist_name path string true "Name of the artist"
// @Success 200 {array} model.TrackDetails
// @Failure 404 {object} map[string]interface{} "tracks not found"
// @Router /track/artist/{artist_name} [get]
func (s *SpotifyService) SearchTrackByArtist(artistName string) ([]model.TrackDetails, error) {
    query := fmt.Sprintf("artist:%s", artistName)
    searchResult, err := s.SpotifyClient.Search(query, spotify.SearchTypeTrack)
    if err != nil || len(searchResult.Tracks.Tracks) == 0 {
        return nil, fmt.Errorf("tracks not found")
    }
    var trackDetails []model.TrackDetails
    for _, track := range searchResult.Tracks.Tracks {
        
        trackDetails = append(trackDetails, model.TrackDetails{
            ISRC:         track.ID.String(),
            Title:        track.Name,
            ArtistName:   track.Artists[0].Name,
			SpotifyImage: track.Album.Images[0].URL,
            
        })
    }
    return trackDetails, nil
}

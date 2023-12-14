package main

import (
	"context"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/naveenbalaji2001/Oauth-with-go/controller"
	"github.com/naveenbalaji2001/Oauth-with-go/dao"
	"github.com/naveenbalaji2001/Oauth-with-go/model"
	"github.com/naveenbalaji2001/Oauth-with-go/service"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var SpotifyCredentials = struct {
	ClientID     string
	ClientSecret string
}{
	ClientID:     "d0a672cffb8149f685b4002ec1bc9da3",
	ClientSecret: "0795acb1e47345138c00c462374d2158",
}

func main() {
	router := gin.Default()

	db, err := gorm.Open("postgres", "postgresql://postgres:Naveen1341@localhost/spotifydb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&model.Track{})

	client := authenticateSpotify()
	if client == nil {
		log.Fatal("Failed to authenticate with Spotify API")
	}

	dbAccessor := dao.NewDatabaseAccessor(db)
	spotifyService := service.NewSpotifyService(client)
	trackController := controller.NewTrackController(dbAccessor, spotifyService)

	router.GET("/track/:isrc", trackController.GetTrackDetailsByISRC)

	router.POST("/track", func(c *gin.Context) {
		var trackDetails model.TrackDetails
		if err := c.BindJSON(&trackDetails); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		if err, _ := dbAccessor.GetTrackByISRC(trackDetails.ISRC); err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Track with ISRC code already exists"})
			return
		}

		newTrack := model.Track{
			ISRC:         trackDetails.ISRC,
			Title:        trackDetails.Title,
			ArtistName:   trackDetails.ArtistName,
			SpotifyImage: trackDetails.SpotifyImage,
		}
		dbAccessor.SaveTrack(&newTrack)

		c.JSON(http.StatusCreated, gin.H{"message": "Track record created successfully"})
	})

	router.Run(":8080")
}

func authenticateSpotify() *spotify.Client {
	config := &clientcredentials.Config{
		ClientID:     SpotifyCredentials.ClientID,
		ClientSecret: SpotifyCredentials.ClientSecret,
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		log.Printf("Failed to get Spotify token: %v", err)
		return nil
	}
 //obtain access token to create new spotify API client
	client := spotify.Authenticator{}.NewClient(token)
	return &client
}

package main

import (
	"context"
	"log"
	_ "net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/naveenbalaji2001/Oauth-with-go/controller"
	"github.com/naveenbalaji2001/Oauth-with-go/dao"
	docs "github.com/naveenbalaji2001/Oauth-with-go/docs"
	"github.com/naveenbalaji2001/Oauth-with-go/model"
	"github.com/naveenbalaji2001/Oauth-with-go/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

var SpotifyCredentials = struct {
	ClientID     string
	ClientSecret string
}{
	ClientID:     "your_client_id",
	ClientSecret: "your_client_secret",
}

// @title Spotify API
// @version 1.0
// @description Golang application using Gin and GORM to interact with the Spotify API.
// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()
    docs.SwaggerInfo.BasePath = "/"
	db, err := gorm.Open("postgres", "postgresql://db_user:db_password@localhost/dbname?sslmode=disable")
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	router.GET("/track/:isrc", trackController.GetTrackDetailsByISRC)

	router.POST("/track/", trackController.CreateTrack)

	router.GET("/track/artist/:artist_name", trackController.SearchTrackByArtist)

    router.PUT("/track/:isrc", trackController.UpdateTheTrack)

	router.Run(":8080")
}

// It authenticates with Spotify and returns a Spotify client
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

	client := spotify.Authenticator{}.NewClient(token)
	return &client
}

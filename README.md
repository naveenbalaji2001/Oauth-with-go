# Spotify Track Details API

This repository contains a simple Go application that serves as an API for retrieving and storing details about music tracks using the International Standard Recording Code (ISRC). The application leverages the Gin framework for building the API and GORM for database interactions, with PostgreSQL as the chosen database. Additionally, it integrates with the Spotify API for obtaining detailed information about tracks.


## Features

Get Track Details: Retrieve detailed information about a track by providing its ISRC code.


## GET /track/:isrc

If the track details are already stored in the database, the API responds with the cached information. Otherwise, it queries the Spotify API to fetch the details, stores them in the database, and returns the information.


## POST /track

 Format for posting track details
{
  "isrc": "example_isrc",
  "title": "Example Track",
  "artist_name": "Example Artist",
  "spotify_image": "https://example.com/image.jpg"
}


## Prerequisites

Before running the application, ensure you have the following set up:

PostgreSQL database with the name spotifydb (you can modify the connection string in the code).
Spotify application credentials (client ID and client secret) to authenticate with the Spotify API.


## Installation and Usage

1. Clone the repository:
### git clone https://github.com/yourusername/spotify-track-api.git

2.Install dependencies:
### go mod tidy

3.Set up the database:

### Modify the connection string in main.go to match your PostgreSQL setup

4.Replace Spotify application credentials in main.go with your own:

	ClientID:     "your_client_id",
	ClientSecret: "your_client_secret",

5. Run the application:

### swag init
To generate docs files.

### go run main.go


### The API will be accessible at http://localhost:8080.

### You can access the Swagger UI by visiting `http://localhost:8080/swagger/index.html` in your web browser.

API should have Swagger documentation, and you can use the Swagger UI to explore and test your API endpoints. 

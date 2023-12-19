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

** PUT /track/:isrc

Update an existing track record in the database by ISRC code.


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

### You can access the Swagger UI by visiting http://localhost:8080/swagger/index.html in your web browser.

API should have Swagger documentation, and you can use the Swagger UI to explore and test your API endpoints. 

### Output Images of postman and postgresDB

![Post_trackdetails](https://github.com/naveenbalaji2001/Oauth-with-go/assets/150377130/65b24b78-f9ac-4d97-a4b3-7658424609f5)
![Get_trackdetailsBy_ISRC](https://github.com/naveenbalaji2001/Oauth-with-go/assets/150377130/8bf9cda9-0b90-444a-b700-d1442ed47aff)
![DB_details_ss](https://github.com/naveenbalaji2001/Oauth-with-go/assets/150377130/0d6e8701-f1ae-42eb-9380-3d82688722c8)

### Output Images of swagger and postgresDB

![swagger_POST_ss](https://github.com/naveenbalaji2001/Oauth-with-go/assets/150377130/7892d9a4-b2b7-4a96-874f-2259d20c55c2)
![swagger_GET_ss](https://github.com/naveenbalaji2001/Oauth-with-go/assets/150377130/1473117f-5494-4077-b555-9dfb5b5c9ec7)
![swagger_GET_ARTIST_NAME_ss](https://github.com/naveenbalaji2001/Oauth-with-go/assets/150377130/02867fad-9917-48c3-b174-4d6dba325571)
![swagger_PUT_ss](https://github.com/naveenbalaji2001/Oauth-with-go/assets/150377130/6352a00d-a9fe-47f4-b948-3a01c73ee528)
![swagger_DB_details](https://github.com/naveenbalaji2001/Oauth-with-go/assets/150377130/414cb14d-7701-4c45-b933-f388c862b15a)

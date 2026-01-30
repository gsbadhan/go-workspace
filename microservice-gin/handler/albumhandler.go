package handler

import (
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
Album request for API
*/
type AlbumRequest struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

/*
Album response for API
*/
type AlbumResponse struct {
	Id     int     `json:"Id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// in-memory store
var store = make(map[int]AlbumResponse, 0)

// Random ID generator
var seqNumber = rand.IntN(1000)

/*
define and initialise API endpoints
*/
func InitializeAlbumRoutes(privateGroupRouter *gin.RouterGroup) {
	log.Println("albums routes initialization..")
	privateGroupRouter.POST("/albums", saveAlbums)
	privateGroupRouter.GET("/albums/:id", getAlbumsById)
	privateGroupRouter.GET("/albums", getAlbums)
	log.Println("albums routes initialized")
}

/*
save new album
*/
func saveAlbums(c *gin.Context) {
	log.Println("album post request", c)
	var request AlbumRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("error in saveAlbums=", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	log.Println("album request body=", request)
	id := rand.IntN(1000)
	response := AlbumResponse{Id: id, Title: request.Title, Artist: request.Artist, Price: request.Price}
	store[id] = response
	c.JSON(http.StatusCreated, gin.H{
		"body": response,
	})
}

/*
Get existing album by ID
*/
func getAlbumsById(c *gin.Context) {
	log.Println("album get request", c)
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	response := store[id]
	c.JSON(http.StatusOK, gin.H{
		"body": response,
	})

}

/*
Get all albums
*/
func getAlbums(c *gin.Context) {
	log.Println("album get request")
	albums := []AlbumResponse{}
	for _, v := range store {
		albums = append(albums, v)
	}
	if len(albums) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"body": albums,
		})
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No album data found",
		})
	}
}

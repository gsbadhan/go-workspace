package handler

import (
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumRequest struct {
	Title  string  `json:"title" binding:"required"`
	Artist string  `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}

type AlbumRespone struct {
	Id     int     `json:"Id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var store = make(map[int]AlbumRespone, 0)
var seqNumber = rand.IntN(1000)

func InitializeRoutes(engine *gin.Engine) {
	log.Println("albums routes initialization..")
	engine.POST("/albums", saveAlbums)
	engine.GET("/albums/:id", getAlbumsById)
	engine.GET("/albums", getAlbums)
	log.Println("albums routes initialized")
}

func saveAlbums(c *gin.Context) {
	log.Println("album post request", c)
	var request AlbumRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println("error in saveAlbums=", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "content not found"})
		return
	}
	log.Println("album request body=", request)
	id := rand.IntN(1000)
	response := AlbumRespone{Id: id, Title: request.Title, Artist: request.Artist, Price: request.Price}
	store[id] = response
	c.JSON(http.StatusCreated, gin.H{
		"body": response,
	})
}

func getAlbumsById(c *gin.Context) {
	log.Println("album get request", c)
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	response := store[id]
	c.JSON(http.StatusOK, gin.H{
		"body": response,
	})

}

func getAlbums(c *gin.Context) {
	log.Println("album get request")
	albums := []AlbumRespone{}
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
			"error": "No data found",
		})
	}
}

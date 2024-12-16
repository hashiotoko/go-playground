package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var jsonIndent string = "  "

func main() {
	router := echo.New()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Start(":8081")
}

func getAlbums(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, albums, jsonIndent)
}

func postAlbums(c echo.Context) error {
	var newAlbum album

	if err := c.Bind(&newAlbum); err != nil {
		return err
	}

	albums = append(albums, newAlbum)
	return c.JSONPretty(http.StatusCreated, newAlbum, jsonIndent)
}

func getAlbumByID(c echo.Context) error {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			return c.JSONPretty(http.StatusOK, album, jsonIndent)
		}
	}

	return c.JSONPretty(http.StatusNotFound, map[string]string{
		"message": "album not found",
	}, jsonIndent)
}

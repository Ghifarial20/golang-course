package controllers

import (
	"fmt"
	dt "golang-api/datastruct"
	"net/http"

	"github.com/gin-gonic/gin"
)

var StoredAlbums []dt.Album

func SeedAlbum(ctx *gin.Context) {
	var newAlbum dt.Album

	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	StoredAlbums = append(StoredAlbums, newAlbum)

	ctx.IndentedJSON(http.StatusCreated, newAlbum)

	fmt.Println(len(StoredAlbums))
}

func ShowAll(ctx *gin.Context) {
	response := make(map[string]interface{})

	response["total"] = len(StoredAlbums)
	response["detail"] = StoredAlbums

	ctx.IndentedJSON(http.StatusOK, response)
}

func ShowSingle(ctx *gin.Context) {
	dataTitle := ctx.Param("titles") // Ambil Parameter

	var newAlbum dt.Album // Isi Response

	var isExist bool // Bila terdapat isi data ygn cocok requeset langusng berhenti

	for i, album := range StoredAlbums {
		if dataTitle == album.Title {
			isExist = true
			newAlbum = StoredAlbums[i]
			break
		}
	}

	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("album with title %v not found", dataTitle),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"album": newAlbum,
	})
}

func ShowAlbums(ctx *gin.Context) {
	var newAlbum, result dt.Album

	isExist := false
	var param interface{}

	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, album := range StoredAlbums {
		if newAlbum.Title == album.Title {
			isExist = true
			result = StoredAlbums[i]
			param = newAlbum.Title
			break
		} else if newAlbum.Artist == album.Artist {
			isExist = true
			result = StoredAlbums[i]
			param = newAlbum.Artist
			break
		}
	}

	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("album with title %v not found", param),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"album": result,
	})
}

func UpdateAlbums(ctx *gin.Context) {
	albumTitle := ctx.Param("title")
	var newAlbum dt.Album
	condition := false

	if err := ctx.ShouldBindJSON(&newAlbum); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, album := range StoredAlbums {
		if newAlbum.Title == album.Title {
			condition = true
			StoredAlbums[i] = newAlbum
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("album with title %v not found", albumTitle),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("album with title %v has been updated", albumTitle),
	})
}

func DeleteAlbums(ctx *gin.Context) {
	albumTitle := ctx.Param("title")
	var albumIndex int
	condition := false

	for i, album := range StoredAlbums {
		if albumTitle == album.Title {
			condition = true
			albumIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("album with title %v not found", albumTitle),
		})
		return
	}

	copy(StoredAlbums[albumIndex:], StoredAlbums[albumIndex+1:])
	StoredAlbums[len(StoredAlbums)-1] = dt.Album{}
	StoredAlbums = StoredAlbums[:len(StoredAlbums)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("album with title %v has been deleted", albumTitle),
	})
}

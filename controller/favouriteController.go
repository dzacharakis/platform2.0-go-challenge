package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/go-webservice/model"
	"bitbucket.org/go-webservice/service"
	"github.com/gin-gonic/gin"
)

type FavouriteController struct {
	FavouriteService service.FavouriteService
}

func (controller *FavouriteController) GetAll(c *gin.Context) {
	log.Printf("GET request at /favourites.\n")
	favourites := controller.FavouriteService.GetAll()
	c.JSON(http.StatusOK, favourites) // 200 OK
}

func (controller *FavouriteController) Get(c *gin.Context) {
	log.Printf("GET request at /favourites/%s.\n", c.Param("id"))
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New(`invalid parameter; not a number`)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) // 400 Bad Request
		return
	}

	favourite, err := controller.FavouriteService.GetByID(id)

	if favourite == (model.Favourite{}) {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()}) // 404 Not found
		return
	}
	c.JSON(http.StatusOK, favourite) // 200 OK
}

func (controller *FavouriteController) GetUserFavourites(c *gin.Context) {
	log.Printf("GET request at /user/%s/favourites.\n", c.Param("id"))
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New(`invalid parameter; not a number`)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) // 400 Bad Request
		return
	}

	user, err := controller.FavouriteService.GetByUserID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()}) // 404 Not Found
		return
	}

	c.JSON(http.StatusOK, user) // 200 OK
}

func (controller *FavouriteController) RemoveUserFavourite(c *gin.Context) {
	log.Printf("DELETE request at /users/%s/favourites/%s.\n", c.Param("userid"), c.Param("assetid"))
	userID, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		err = errors.New(`invalid parameter; not a number`)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) // 400 Bad Request
		return
	}

	assetID, err := strconv.Atoi(c.Param("assetid"))
	if err != nil {
		err = errors.New(`invalid parameter; not a number`)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) // 400 Bad Request
		return
	}

	count, err := controller.FavouriteService.DeleteByUserIDAndAssetID(userID, assetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()}) // 500 Internal Server Error
		return
	}

	if count == 0 { // 0 rows affected
		c.Writer.WriteHeader(http.StatusNotFound) // 404 Not Found
		return
	}
	c.Writer.WriteHeader(http.StatusNoContent) // 204 No Content
}

func (controller *FavouriteController) Create(c *gin.Context) {
	log.Printf("POST request at /favourites.\n")

	favourite := new(model.Favourite)

	// Call BindJSON to bind the received JSON to Favourite entity.
	if err := c.BindJSON(&favourite); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) // 400 Bad Request
		return
	}

	returningID, err := controller.FavouriteService.CreateFavourite(*favourite)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) // 400 Bad Request
		return
	}

	locationURI := fmt.Sprintf("http://%s%s/%d", c.Request.Host, c.Request.RequestURI, returningID)
	c.Writer.Header().Set("Location", locationURI)
	c.Writer.WriteHeader(http.StatusCreated) // 201 Created
}

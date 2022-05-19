package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/go-webservice/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func (controller *UserController) GetAll(c *gin.Context) {
	log.Println("GET request at /users.")
	users := controller.UserService.GetAll()
	c.JSON(http.StatusOK, users) // 200 OK
}

func (controller *UserController) Get(c *gin.Context) {
	log.Printf("GET request at /users/%s.\n", c.Param("id"))
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New(`invalid parameter; not a number`)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) // 400 Bad Request
		return
	}

	user, err := controller.UserService.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()}) // 404 Not Found
		return
	}

	c.JSON(http.StatusOK, user) // 200 OK
}

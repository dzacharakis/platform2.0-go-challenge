package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/go-webservice/service"
	"github.com/gin-gonic/gin"
)

type AudienceController struct {
	AudienceService service.AudienceService
}

func (controller *AudienceController) GetAll(c *gin.Context) {
	log.Printf("GET request at /variables.\n")
	audiences := controller.AudienceService.GetAll()
	c.JSON(http.StatusOK, audiences) // 200 OK
}

func (controller *AudienceController) Get(c *gin.Context) {
	log.Printf("GET request at /users/%s.\n", c.Param("id"))
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New(`invalid parameter; not a number`)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) // 400 Bad Request
		return
	}

	variable, err := controller.AudienceService.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()}) // 404 Not Found
		return
	}

	c.JSON(http.StatusOK, variable) // 200 OK
}

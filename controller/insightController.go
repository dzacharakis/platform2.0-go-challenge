package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/go-webservice/service"
	"github.com/gin-gonic/gin"
)

type InsightController struct {
	InsightService service.InsightService
}

func (controller *InsightController) GetAll(c *gin.Context) {
	log.Println("GET request at /insights.")
	insights := controller.InsightService.GetAll()
	c.JSON(200, insights)
}

func (controller *InsightController) Get(c *gin.Context) {
	log.Printf("GET request at /insights/%s.\n", c.Param("id"))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New(`invalid parameter; not a number`)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) // 400 Bad Request
		return
	}

	insight, err := controller.InsightService.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()}) // 404 Not Found
		return
	}

	c.JSON(http.StatusOK, insight) // 200 OK

}

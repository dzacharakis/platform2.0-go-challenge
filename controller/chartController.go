package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/go-webservice/service"
	"github.com/gin-gonic/gin"
)

type ChartController struct {
	ChartService service.ChartService
}

func (controller *ChartController) GetAll(c *gin.Context) {
	log.Println("GET request at /charts")
	charts := controller.ChartService.GetAll()
	c.JSON(200, charts)
}

func (controller *ChartController) Get(c *gin.Context) {
	log.Printf("GET request at /charts/%s.\n", c.Param("id"))
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = errors.New(`invalid parameter; not a number`)
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()}) // 400 Bad Request
		return
	}

	chart, err := controller.ChartService.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, chart)
}

package main

import (
	"database/sql"
	"log"

	"bitbucket.org/go-webservice/controller"
	"bitbucket.org/go-webservice/repository"
	"bitbucket.org/go-webservice/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=dimitris dbname=gwichallenge password=gwidbpass host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// initialize application, dependency injection
	audienceRepository := &repository.AudienceRepository{Database: db}
	audienceService := &service.AudienceService{AudienceRepository: *audienceRepository}
	audienceController := &controller.AudienceController{AudienceService: *audienceService}

	userRepository := &repository.UserRepository{Database: db}
	userService := &service.UserService{UserRepository: *userRepository}
	userController := &controller.UserController{UserService: *userService}

	insightRepository := &repository.InsightRepository{Database: db}
	insightService := &service.InsightService{InsightRepository: *insightRepository}
	insightController := &controller.InsightController{InsightService: *insightService}

	chartRepository := &repository.ChartRepository{Database: db}
	chartDataRepository := &repository.ChartDataRepository{Database: db}
	chartService := &service.ChartService{ChartRepository: *chartRepository, ChartDataRepository: *chartDataRepository}
	chartController := &controller.ChartController{ChartService: *chartService}

	favouriteRepository := &repository.FavouriteRepository{Database: db}
	favouriteService := &service.FavouriteService{FavouriteRepository: *favouriteRepository, UserRepository: *userRepository}
	favouriteController := &controller.FavouriteController{FavouriteService: *favouriteService}

	// HERE STARTS gin
	log.Printf("Starting HTTP service on %s ...", ":8080")
	r := gin.Default()
	r.GET("/users", userController.GetAll)
	r.GET("/users/:id", userController.Get)
	r.GET("/variables", audienceController.GetAll)
	r.GET("/variables/:id", audienceController.Get)
	r.GET("/insights", insightController.GetAll)
	r.GET("/insights/:id", insightController.Get)
	r.GET("/charts", chartController.GetAll)
	r.GET("/charts/:id", chartController.Get)
	r.GET("/users/:id/favourites", favouriteController.GetUserFavourites)
	r.DELETE("/users/:userid/favourites/:assetid", favouriteController.RemoveUserFavourite)
	r.GET("/favourites", favouriteController.GetAll)
	r.POST("/favourites", favouriteController.Create)
	r.GET("/favourites/:id", favouriteController.Get)

	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 to use it with dns record - demetrius.ddns.net
	// r.RunTLS("0.0.0.0:8080", `../../certs/server.crt`, `../../certs/server.key`)
}

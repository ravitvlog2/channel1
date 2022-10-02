package main

import (
	"rafit/config"
	"rafit/handler"
	"rafit/repository"
	"rafit/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.Config()
	// db.AutoMigrate([]models.IPModel{})

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	router := gin.Default()
	router.GET("/ip", handler.GetAllIp)
	router.POST("/ip", handler.CreateIp)

	router.Run()

}

package handler

import (
	"net/http"
	"rafit/helpers"
	"rafit/models"
	"rafit/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.Service
}

func NewHandler(service service.Service) *handler {
	return &handler{service}
}

func (h *handler) CreateIp(c *gin.Context) {
	request := models.ReqIpModel{}

	err := c.ShouldBindJSON(&request)
	if err != nil {
		errorMessage := gin.H{"errors": err}
		respon := helpers.ResponseApi("Failed insert", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}

	result, err := h.service.CreateIP(request)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		respon := helpers.ResponseApi("Failed insert", http.StatusBadRequest, "Failed", errorMessage)
		c.JSON(http.StatusBadRequest, respon)
		return
	}

	response := helpers.ResponseApi("success", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *handler) GetAllIp(c *gin.Context) {
	result, err := h.service.GetAllIp()
	if err != nil {
		errorMessage := gin.H{"errors": err}
		respon := helpers.ResponseApi("Failed get data", http.StatusInternalServerError, "Failed", errorMessage)
		c.JSON(http.StatusInternalServerError, respon)
		return
	}

	response := helpers.ResponseApi("success", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

package controllers

import (
	"net/http"
	"pos-desy/config"
	"pos-desy/helper"
	"pos-desy/models"

	"github.com/gin-gonic/gin"
)

type CategoryController struct{}

func (cc *CategoryController) GetCategoryData(c *gin.Context) {
	var categories []models.Category
	result := config.DB.Find(&categories)

	if result.Error != nil {
		response := helper.APIResponse("Failed Get Data Category", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success Get Category Data", http.StatusOK, "Success", categories)

	c.JSON(http.StatusOK, response)
}

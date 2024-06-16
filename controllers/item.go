package controllers

import (
	"log"
	"net/http"
	"pos-desy/config"
	"pos-desy/helper"
	"pos-desy/models"

	"github.com/gin-gonic/gin"
)

type ItemController struct{}

func (cc *ItemController) GetItemData(c *gin.Context) {
	var Items []models.Item

	result := config.DB.Find(&Items)

	if result.Error != nil {
		log.Fatal("error when fetching data", result.Error)
		response := helper.APIResponse("Error fetch data item", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Fetch Item Data", http.StatusOK, "success", Items)
	c.JSON(http.StatusBadRequest, response)
}

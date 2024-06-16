package routes

import (
	"pos-desy/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()

	categoryController := new(controllers.CategoryController)
	// category routes
	categoryRoute := r.Group("/categories")
	{
		categoryRoute.GET("/get-category-data", categoryController.GetCategoryData)
	}

	// item routes
	itemRoute := r.Group("/item")
	{
		itemRoute.GET("/get-item-data")
	}

	return r
}

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

	itemController := new(controllers.ItemController)
	// item routes
	itemRoute := r.Group("/items")
	{
		itemRoute.GET("/get-item-data", itemController.GetItemData)
	}

	reportingRoute := r.Group("/report")
	{
		// di sorting berdasarkan tanggal
		reportingRoute.GET("/get-report-data")
	}

	sellRoute := r.Group("/report")
	{
		// merupakan detail dari reporting -> satuan dari summary
		// harus ada filtering dengan tanggal
		sellRoute.GET("/get-sell-data")
	}

	return r
}

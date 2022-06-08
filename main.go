package main

import (
	"example/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// intialize gin router
	router := gin.Default()

	// index handler to view all registered routes
	router.GET("/", func(c *gin.Context) {
		type finalPath struct {
			Method string
			Path   string
		}

		data := router.Routes()
		finalPaths := make([]finalPath, 0)

		for i := 0; i < len(data); i++ {
			finalPaths = append(finalPaths, finalPath{
				Path:   data[i].Path,
				Method: data[i].Method,
			})
		}
		c.JSON(200, gin.H{
			"routes": finalPaths,
		})
	})

	//route for insert the element into the tree
	router.POST("/v1/insert", routes.InsertElement)

	//route for fetch the element from the tree
	router.GET("/v1/query", routes.QueryElement)

	//run the application
	router.Run()
}

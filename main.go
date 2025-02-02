package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var data api

func main() {
	r := gin.Default()
	r.GET("/get", getValues)
	r.POST("/post", postValues)
	r.PUT("/put", putValues)
	r.DELETE("/delete", deleteValues)
	r.Run(":9090")
}

func getValues(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func postValues(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func putValues(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "something went wrong",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func deleteValues(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "something went bad",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted",
	})
}

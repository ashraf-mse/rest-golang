package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Posts(c *gin.Context) {

}

func Show(c *gin.Context) {

}

func Store(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Something wrong with validation",
			"data":  "",
		})
		return
	}

	db.Create(&post)
	c.JSON(http.StatusCreated, gin.H{
		"error": "",
		"data":  post,
	})
}

func Update(c *gin.Context) {

}

func Delete(c *gin.Context) {

}

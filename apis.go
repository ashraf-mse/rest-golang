package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Posts(c *gin.Context) {

	// limit = c.DefaultQuery("limit","15")
	// offset := c.DefaultQuery("offset","0")
	var posts []Post
	db.Limit(15).Offset(0).Find(&posts)
	c.JSON(http.StatusAccepted, gin.H{
		"error": "",
		"data":  posts,
	})
}

func Show(c *gin.Context) {

	id := c.Param("id")
	var post Post
	db.First(&post, id)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post Not Found",
			"data":  "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": "",
		"data":  post,
	})
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
	id := c.Param("id")
	var post Post
	db.First(&post, id)
	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post Not Found",
			"data":  "",
		})
		return
	}
	db.Delete(&post)
	c.JSON(http.StatusOK, gin.H{
		"error": "",
		"data":  "Post Deleted",
	})
}

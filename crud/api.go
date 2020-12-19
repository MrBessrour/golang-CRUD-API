package main

import (
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

//select query with limit and offset
func Posts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	var posts []Post
	db.Limit(limit).Offset(offset).Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"messege": "",
		"data":    posts,
	})

}

//showing a post with it's id passed in the URL with a GET request
func Show(c *gin.Context) {
	post := getById(c)
	c.JSON(http.StatusOK, gin.H{
		"messege": "",
		"data":    post,
	})

}

//storing a mew post to the db with a POST request with
func Store(c *gin.Context) {
	var post Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messege": err.Error(),
			"data":    "",
		})
		return
	}
	post.Status = "Active"
	db.Create(&post)
	c.JSON(http.StatusOK, gin.H{
		"messege": "",
		"data":    post,
	})
}

//delete a post by it's id
func Delete(c *gin.Context) {
	post := getById(c)
	if post.ID == 0 {
		return
	}
	db.Unscoped().Delete(&post)
	c.JSON(http.StatusOK, gin.H{
		"messege": "deleted successfuly",
		"data":    "",
	})

}

//update a post with a Ptach request , the id sent in the URL
func Update(c *gin.Context) {
	oldpost := getById(c)
	var newpost Post
	if err := c.ShouldBindJSON(&newpost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messege": err.Error(),
			"data":    "",
		})
		return
	}
	oldpost.Title = newpost.Title
	oldpost.Des = newpost.Des
	if newpost.Status != "" {
		oldpost.Status = newpost.Status
	}

	db.Save(&oldpost)

	c.JSON(http.StatusOK, gin.H{
		"messege": "Post has been updated",
		"data":    oldpost,
	})

}

func getById(c *gin.Context) Post {
	id := c.Param("id")
	var post Post
	db.First(&post, id)
	if post.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"messege": "post not found",
			"data":    "",
		})
	}
	return post
}

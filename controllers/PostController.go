package controllers

import (
	"CMSGo/config"
	"CMSGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/heroku/x/hmetrics/onload"
)

func PostPost(c *gin.Context) {
	db := config.InitDb()
	defer db.Close()

	var post models.Posts
	c.Bind(&post)

	if post.Title != "" && post.Content != "" {
		// INSERT INTO "posts" (name) VALUES (post.Name);
		db.Create(&post)
		// Display error
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"code":    200,
			"message": "Success",
			"data":    post,
		})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/posts
}

func GetPosts(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	var posts []models.Posts
	// SELECT * FROM posts
	db.Find(&posts)

	// Display JSON result
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"code":    200,
		"message": "Success",
		"data":    posts,
	})
	// curl -i http://localhost:8080/api/v1/posts
}

func GetPost(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var post models.Posts
	// SELECT * FROM posts WHERE id = 1;
	db.First(&post, id)

	if post.ID != 0 {
		// Display JSON result
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  true,
			"code":    200,
			"message": "Success",
			"post":    post,
		})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Post not found"})
	}

	// curl -i http://localhost:8080/api/v1/posts/1
}

func UpdatePost(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id post
	id := c.Params.ByName("id")
	var post models.Posts
	// SELECT * FROM posts WHERE id = 1;
	db.First(&post, id)

	if post.Title != "" && post.Content != "" {

		if post.ID != 0 {
			var newPost models.Posts
			c.Bind(&newPost)

			result := models.Posts{
				Title:    newPost.Title,
				Content:  newPost.Content,
				Category: newPost.Category,
				Status:   newPost.Status,
			}

			// UPDATE posts SET firstname='newPost.Firstname', lastname='newPost.Lastname' WHERE id = post.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "Post not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/posts/1
}

func DeletePost(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id post
	id := c.Params.ByName("id")
	var post models.Posts
	// SELECT * FROM posts WHERE id = 1;
	db.First(&post, id)

	if post.ID != 0 {
		// DELETE FROM posts WHERE id = post.Id
		db.Delete(&post)
		// Display JSON result
		c.JSON(200, gin.H{"success": "Post #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Post not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/posts/1
}

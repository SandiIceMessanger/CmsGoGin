package controllers

import (
	"CMSGo/config"
	"CMSGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/heroku/x/hmetrics/onload"
)

func PostArticle(c *gin.Context) {
	db := config.InitDb()
	defer db.Close()

	var article models.Articles
	c.Bind(&article)

	if article.Title != "" && article.Content != "" {
		// INSERT INTO "articles" (name) VALUES (article.Name);
		db.Create(&article)
		// Display error
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"code":    200,
			"message": "Success",
			"data":    article,
		})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/articles
}

func GetArticles(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	var articles []models.Articles
	// SELECT * FROM articles
	db.Find(&articles)

	// Display JSON result
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"code":    200,
		"message": "Success",
		"data":    articles,
	})
	// curl -i http://localhost:8080/api/v1/articles
}

func GetArticle(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var article models.Articles
	// SELECT * FROM articles WHERE id = 1;
	db.First(&article, id)

	if article.ID != 0 {
		// Display JSON result
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  true,
			"code":    200,
			"message": "Success",
			"article": article,
		})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Article not found"})
	}

	// curl -i http://localhost:8080/api/v1/articles/1
}

func UpdateArticle(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id article
	id := c.Params.ByName("id")
	var article models.Articles
	// SELECT * FROM articles WHERE id = 1;
	db.First(&article, id)

	if article.Title != "" && article.Content != "" {

		if article.ID != 0 {
			var newArticle models.Articles
			c.Bind(&newArticle)

			result := models.Articles{
				Title:    newArticle.Title,
				Content:  newArticle.Content,
				Category: newArticle.Category,
				Status:   newArticle.Status,
			}

			// UPDATE articles SET firstname='newArticle.Firstname', lastname='newArticle.Lastname' WHERE id = article.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "Article not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/articles/1
}

func DeleteArticle(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id article
	id := c.Params.ByName("id")
	var article models.Articles
	// SELECT * FROM articles WHERE id = 1;
	db.First(&article, id)

	if article.ID != 0 {
		// DELETE FROM articles WHERE id = article.Id
		db.Delete(&article)
		// Display JSON result
		c.JSON(200, gin.H{"success": "Article #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Article not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/articles/1
}

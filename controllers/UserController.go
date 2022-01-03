package controllers

import (
	"CMSGo/config"
	"CMSGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/heroku/x/hmetrics/onload"
)

func PostUser(c *gin.Context) {
	db := config.InitDb()
	defer db.Close()

	var user models.Users
	c.Bind(&user)

	if user.Title != "" && user.Content != "" {
		// INSERT INTO "users" (name) VALUES (user.Name);
		db.Create(&user)
		// Display error
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"code":    200,
			"message": "Success",
			"data":    user,
		})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
}

func GetUsers(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	var users []models.Users
	// SELECT * FROM users
	db.Find(&users)

	// Display JSON result
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"code":    200,
		"message": "Success",
		"data":    users,
	})
	// curl -i http://localhost:8080/api/v1/users
}

func GetUser(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var user models.Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.ID != 0 {
		// Display JSON result
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":  true,
			"code":    200,
			"message": "Success",
			"user":    user,
		})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}

	// curl -i http://localhost:8080/api/v1/users/1
}

func UpdateUser(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id user
	id := c.Params.ByName("id")
	var user models.Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.Title != "" && user.Content != "" {

		if user.ID != 0 {
			var newUser models.Users
			c.Bind(&newUser)

			result := models.Users{
				Title:    newUser.Title,
				Content:  newUser.Content,
				Category: newUser.Category,
				Status:   newUser.Status,
			}

			// UPDATE users SET firstname='newUser.Firstname', lastname='newUser.Lastname' WHERE id = user.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "User not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

func DeleteUser(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id user
	id := c.Params.ByName("id")
	var user models.Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.ID != 0 {
		// DELETE FROM users WHERE id = user.Id
		db.Delete(&user)
		// Display JSON result
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}

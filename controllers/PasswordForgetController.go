package controllers

import (
	"CMSGo/config"
	"CMSGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/heroku/x/hmetrics/onload"
)

func PostPasswordForget(c *gin.Context) {
	db := config.InitDb()
	defer db.Close()

	var passwordForget models.PasswordForgets
	c.Bind(&passwordForget)

	if passwordForget.Title != "" && passwordForget.Content != "" {
		// INSERT INTO "passwordForgets" (name) VALUES (passwordForget.Name);
		db.Create(&passwordForget)
		// Display error
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"code":    200,
			"message": "Success",
			"data":    passwordForget,
		})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/passwordForgets
}

func GetPasswordForgets(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	var passwordForgets []models.PasswordForgets
	// SELECT * FROM passwordForgets
	db.Find(&passwordForgets)

	// Display JSON result
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"code":    200,
		"message": "Success",
		"data":    passwordForgets,
	})
	// curl -i http://localhost:8080/api/v1/passwordForgets
}

func GetPasswordForget(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var passwordForget models.PasswordForgets
	// SELECT * FROM passwordForgets WHERE id = 1;
	db.First(&passwordForget, id)

	if passwordForget.ID != 0 {
		// Display JSON result
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":         true,
			"code":           200,
			"message":        "Success",
			"passwordForget": passwordForget,
		})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "PasswordForget not found"})
	}

	// curl -i http://localhost:8080/api/v1/passwordForgets/1
}

func UpdatePasswordForget(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id passwordForget
	id := c.Params.ByName("id")
	var passwordForget models.PasswordForgets
	// SELECT * FROM passwordForgets WHERE id = 1;
	db.First(&passwordForget, id)

	if passwordForget.Title != "" && passwordForget.Content != "" {

		if passwordForget.ID != 0 {
			var newPasswordForget models.PasswordForgets
			c.Bind(&newPasswordForget)

			result := models.PasswordForgets{
				Title:    newPasswordForget.Title,
				Content:  newPasswordForget.Content,
				Category: newPasswordForget.Category,
				Status:   newPasswordForget.Status,
			}

			// UPDATE passwordForgets SET firstname='newPasswordForget.Firstname', lastname='newPasswordForget.Lastname' WHERE id = passwordForget.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "PasswordForget not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/passwordForgets/1
}

func DeletePasswordForget(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id passwordForget
	id := c.Params.ByName("id")
	var passwordForget models.PasswordForgets
	// SELECT * FROM passwordForgets WHERE id = 1;
	db.First(&passwordForget, id)

	if passwordForget.ID != 0 {
		// DELETE FROM passwordForgets WHERE id = passwordForget.Id
		db.Delete(&passwordForget)
		// Display JSON result
		c.JSON(200, gin.H{"success": "PasswordForget #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "PasswordForget not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/passwordForgets/1
}

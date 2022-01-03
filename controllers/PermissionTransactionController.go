package controllers

import (
	"CMSGo/config"
	"CMSGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/heroku/x/hmetrics/onload"
)

func PostPermissionTransaction(c *gin.Context) {
	db := config.InitDb()
	defer db.Close()

	var permissionTransactions models.PermissionTransactions
	c.Bind(&permissionTransactions)

	if permissionTransactions.Title != "" && permissionTransactions.Content != "" {
		// INSERT INTO "permissionTransactionss" (name) VALUES (permissionTransactions.Name);
		db.Create(&permissionTransactions)
		// Display error
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"code":    200,
			"message": "Success",
			"data":    permissionTransactions,
		})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/permissionTransactionss
}

func GetPermissionTransactions(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	var permissionTransactionss []models.PermissionTransactions
	// SELECT * FROM permissionTransactionss
	db.Find(&permissionTransactionss)

	// Display JSON result
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"code":    200,
		"message": "Success",
		"data":    permissionTransactionss,
	})
	// curl -i http://localhost:8080/api/v1/permissionTransactionss
}

func GetPermissionTransaction(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var permissionTransactions models.PermissionTransactions
	// SELECT * FROM permissionTransactionss WHERE id = 1;
	db.First(&permissionTransactions, id)

	if permissionTransactions.ID != 0 {
		// Display JSON result
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":                 true,
			"code":                   200,
			"message":                "Success",
			"permissionTransactions": permissionTransactions,
		})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "PermissionTransaction not found"})
	}

	// curl -i http://localhost:8080/api/v1/permissionTransactionss/1
}

func UpdatePermissionTransaction(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id permissionTransactions
	id := c.Params.ByName("id")
	var permissionTransactions models.PermissionTransactions
	// SELECT * FROM permissionTransactionss WHERE id = 1;
	db.First(&permissionTransactions, id)

	if permissionTransactions.Title != "" && permissionTransactions.Content != "" {

		if permissionTransactions.ID != 0 {
			var newPermissionTransaction models.PermissionTransactions
			c.Bind(&newPermissionTransaction)

			result := models.PermissionTransactions{
				Title:    newPermissionTransaction.Title,
				Content:  newPermissionTransaction.Content,
				Category: newPermissionTransaction.Category,
				Status:   newPermissionTransaction.Status,
			}

			// UPDATE permissionTransactionss SET firstname='newPermissionTransaction.Firstname', lastname='newPermissionTransaction.Lastname' WHERE id = permissionTransactions.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "PermissionTransaction not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/permissionTransactionss/1
}

func DeletePermissionTransaction(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id permissionTransactions
	id := c.Params.ByName("id")
	var permissionTransactions models.PermissionTransactions
	// SELECT * FROM permissionTransactionss WHERE id = 1;
	db.First(&permissionTransactions, id)

	if permissionTransactions.ID != 0 {
		// DELETE FROM permissionTransactionss WHERE id = permissionTransactions.Id
		db.Delete(&permissionTransactions)
		// Display JSON result
		c.JSON(200, gin.H{"success": "PermissionTransaction #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "PermissionTransaction not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/permissionTransactionss/1
}

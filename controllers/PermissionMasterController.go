package controllers

import (
	"CMSGo/config"
	"CMSGo/models"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/heroku/x/hmetrics/onload"
)

func PostPermissionMaster(c *gin.Context) {
	db := config.InitDb()
	defer db.Close()

	var permissionMaster models.PermissionMasters
	c.Bind(&permissionMaster)

	if permissionMaster.Title != "" && permissionMaster.Content != "" {
		// INSERT INTO "permissionMasters" (name) VALUES (permissionMaster.Name);
		db.Create(&permissionMaster)
		// Display error
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"code":    200,
			"message": "Success",
			"data":    permissionMaster,
		})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/permissionMasters
}

func GetPermissionMasters(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	var permissionMasters []models.PermissionMasters
	// SELECT * FROM permissionMasters
	db.Find(&permissionMasters)

	// Display JSON result
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"code":    200,
		"message": "Success",
		"data":    permissionMasters,
	})
	// curl -i http://localhost:8080/api/v1/permissionMasters
}

func GetPermissionMaster(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var permissionMaster models.PermissionMasters
	// SELECT * FROM permissionMasters WHERE id = 1;
	db.First(&permissionMaster, id)

	if permissionMaster.ID != 0 {
		// Display JSON result
		c.JSON(http.StatusOK, map[string]interface{}{
			"status":           true,
			"code":             200,
			"message":          "Success",
			"permissionMaster": permissionMaster,
		})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "PermissionMaster not found"})
	}

	// curl -i http://localhost:8080/api/v1/permissionMasters/1
}

func UpdatePermissionMaster(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id permissionMaster
	id := c.Params.ByName("id")
	var permissionMaster models.PermissionMasters
	// SELECT * FROM permissionMasters WHERE id = 1;
	db.First(&permissionMaster, id)

	if permissionMaster.Title != "" && permissionMaster.Content != "" {

		if permissionMaster.ID != 0 {
			var newPermissionMaster models.PermissionMasters
			c.Bind(&newPermissionMaster)

			result := models.PermissionMasters{
				Title:    newPermissionMaster.Title,
				Content:  newPermissionMaster.Content,
				Category: newPermissionMaster.Category,
				Status:   newPermissionMaster.Status,
			}

			// UPDATE permissionMasters SET firstname='newPermissionMaster.Firstname', lastname='newPermissionMaster.Lastname' WHERE id = permissionMaster.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "PermissionMaster not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/permissionMasters/1
}

func DeletePermissionMaster(c *gin.Context) {
	// Connection to the database
	db := config.InitDb()
	// Close connection database
	defer db.Close()

	// Get id permissionMaster
	id := c.Params.ByName("id")
	var permissionMaster models.PermissionMasters
	// SELECT * FROM permissionMasters WHERE id = 1;
	db.First(&permissionMaster, id)

	if permissionMaster.ID != 0 {
		// DELETE FROM permissionMasters WHERE id = permissionMaster.Id
		db.Delete(&permissionMaster)
		// Display JSON result
		c.JSON(200, gin.H{"success": "PermissionMaster #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "PermissionMaster not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/permissionMasters/1
}

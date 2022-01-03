package routes

import (
	"CMSGo/controllers"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/heroku/x/hmetrics/onload"
	cors "github.com/rs/cors/wrapper/gin"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           int(12 * time.Hour),
	})

	r.Use(corsConfig)

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	v1 := r.Group("api/v1")
	{
		v1.POST("/password_forgets", controllers.PostPasswordForget)
		v1.GET("/password_forgets", controllers.GetPasswordForgets)
		v1.GET("/password_forgets/:id", controllers.GetPasswordForget)
		v1.PUT("/password_forgets/:id", controllers.UpdatePasswordForget)
		v1.DELETE("/password_forgets/:id", controllers.DeletePasswordForget)

		v1.POST("/permission_masters", controllers.PostPermissionMaster)
		v1.GET("/permission_masters", controllers.GetPermissionMasters)
		v1.GET("/permission_masters/:id", controllers.GetPermissionMaster)
		v1.PUT("/permission_masters/:id", controllers.UpdatePermissionMaster)
		v1.DELETE("/permission_masters/:id", controllers.DeletePermissionMaster)

		v1.POST("/permission_transactions", controllers.PostPermissionTransaction)
		v1.GET("/permission_transactions", controllers.GetPermissionTransactions)
		v1.GET("/permission_transactions/:id", controllers.GetPermissionTransaction)
		v1.PUT("/permission_transactions/:id", controllers.UpdatePermissionTransaction)
		v1.DELETE("/permission_transactions/:id", controllers.DeletePermissionTransaction)

		v1.POST("/posts", controllers.PostPost)
		v1.GET("/posts", controllers.GetPosts)
		v1.GET("/posts/:id", controllers.GetPost)
		v1.PUT("/posts/:id", controllers.UpdatePost)
		v1.DELETE("/posts/:id", controllers.DeletePost)

		v1.POST("/users", controllers.PostUser)
		v1.GET("/users", controllers.GetUsers)
		v1.GET("/users/:id", controllers.GetUser)
		v1.PUT("/users/:id", controllers.UpdateUser)
		v1.DELETE("/users/:id", controllers.DeleteUser)
	}

	return r
}

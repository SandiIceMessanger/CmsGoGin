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
		v1.POST("/articles", controllers.PostArticle)
		v1.GET("/articles", controllers.GetArticles)
		v1.GET("/articles/:id", controllers.GetArticle)
		v1.PUT("/articles/:id", controllers.UpdateArticle)
		v1.DELETE("/articles/:id", controllers.DeleteArticle)
	}

	return r
}

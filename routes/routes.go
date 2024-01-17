package routes

import (
	"backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.Use(cors.Default())
	r.GET("/article", controllers.Findarticle)
	r.GET("/article/:id", controllers.FindarticlebyId)
	r.GET("/articles/:limit/:offset", controllers.Findbylimit)
	r.POST("/article", controllers.Addarticle)
	r.PATCH("/article/:id", controllers.Updatearticle)
	r.DELETE("/article/:id", controllers.Deletearticle)

	return r
}

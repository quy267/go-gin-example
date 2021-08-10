package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/component/appctx"
	restaurantGin "go-gin-example/modules/restaurant/transport/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	log.Println(db, err)

	db = db.Debug()

	appCtx := appctx.NewAppContext(db)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantGin.CreateRestaurant(appCtx))
			restaurants.GET("/:id", restaurantGin.GetRestaurant(appCtx))
			restaurants.GET("", restaurantGin.ListRestaurant(appCtx))
			restaurants.PUT("/:id", restaurantGin.UpdateRestaurant(appCtx))
			restaurants.DELETE("/:id", restaurantGin.DeleteRestaurant(appCtx))
		}
	}

	r.Run()
}

package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type RestaurantCreate struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"` // tag
	Address string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type Restaurant struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"` // tag
	Address string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"` // tag
	Address *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	log.Println(db, err)

	db = db.Debug()

	//newRestaurant := Restaurant{Name: "200Lab", Address: "Somewhere"}
	//
	//if err := db.Create(&newRestaurant).Error; err != nil {
	//	log.Println(err)
	//}

	//log.Println("Inserted ID:", newRestaurant.Id)

	//var oldRes Restaurant
	//
	//if err := db.Where(map[string]interface{}{"id": 3}).First(&oldRes).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(oldRes)
	//
	//var listRes []Restaurant
	//
	//if err := db.Limit(10).Find(&listRes).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(listRes)
	//
	//emptyStr := "Tani"
	//
	//dataUpdate := RestaurantUpdate{
	//	Name: &emptyStr,
	//	//Address: nil,
	//}
	//
	//if err := db.Where("id = ?", 3).Updates(&dataUpdate).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//if err := db.Table(Restaurant{}.TableName()).Where("id = ?", 3).Delete(nil).Error; err != nil {
	//	log.Println(err)
	//}

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
			restaurants.POST("", func(c *gin.Context) {
				var newData RestaurantCreate

				if err := c.ShouldBind(&newData); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Create(&newData).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": newData.Id})
			})

			restaurants.GET("/:id", func(c *gin.Context) {
				var data Restaurant

				id, err := strconv.Atoi(c.Param("id"))

				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Where("id = ?", id).First(&data).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": data})
			})

			restaurants.GET("", func(c *gin.Context) {
				var data []Restaurant

				type Paging struct {
					Page  int `json:"page" form:"page"`
					Limit int `json:"limit" form:"limit"`
				}

				var paging Paging

				if err := c.ShouldBind(&paging); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if paging.Page < 1 {
					paging.Page = 1
				}

				if paging.Limit <= 0 {
					paging.Limit = 10
				}

				offset := (paging.Page - 1) * paging.Limit

				if err := db.Offset(offset).Limit(paging.Limit).Order("id desc").Find(&data).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": data})
			})

			restaurants.PUT("/:id", func(c *gin.Context) {
				id, err := strconv.Atoi(c.Param("id"))

				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				var data RestaurantUpdate

				if err := c.ShouldBind(&data); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": true})
			})

			restaurants.DELETE("/:id", func(c *gin.Context) {
				id, err := strconv.Atoi(c.Param("id"))

				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if err := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": true})
			})
		}
	}

	r.Run()
}

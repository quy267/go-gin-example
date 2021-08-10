package restaurantgin

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/component/appctx"
	restaurantBiz "go-gin-example/modules/restaurant/biz"
	restaurantModel "go-gin-example/modules/restaurant/model"
	restaurantstorage "go-gin-example/modules/restaurant/storage"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var newData restaurantModel.RestaurantCreate

		if err := c.ShouldBind(&newData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Dependencies install
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateNewRestaurant(c.Request.Context(), &newData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": newData.Id})
	}
}

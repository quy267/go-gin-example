package restaurantgin

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/component/appctx"
	restaurantBiz "go-gin-example/modules/restaurant/biz"
	restaurantModel "go-gin-example/modules/restaurant/model"
	restaurantStorage "go-gin-example/modules/restaurant/storage"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var data restaurantModel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantStorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

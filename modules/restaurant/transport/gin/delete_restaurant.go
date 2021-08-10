package restaurantgin

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/component/appctx"
	restaurantBiz "go-gin-example/modules/restaurant/biz"
	restaurantStorage "go-gin-example/modules/restaurant/storage"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantStorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}

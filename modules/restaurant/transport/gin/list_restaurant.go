package restaurantgin

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/common"
	"go-gin-example/component/appctx"
	restaurantBiz "go-gin-example/modules/restaurant/biz"
	restaurantModel "go-gin-example/modules/restaurant/model"
	restaurantStorage "go-gin-example/modules/restaurant/storage"
	"net/http"
)

func ListRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var filter restaurantModel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := paging.Process(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantStorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewListRestaurantBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging})
	}
}

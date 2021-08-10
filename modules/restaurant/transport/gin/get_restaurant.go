package restaurantgin

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-gin-example/common"
	"go-gin-example/component/appctx"
	restaurantBiz "go-gin-example/modules/restaurant/biz"
	restaurantModel "go-gin-example/modules/restaurant/model"
	restaurantStorage "go-gin-example/modules/restaurant/storage"
	"net/http"
	"strconv"
	"time"
)

type fakeGetDataStore struct{}

func (fakeGetDataStore) FindDataWithCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*restaurantModel.Restaurant, error) {
	return &restaurantModel.Restaurant{
		SQLModel: common.SQLModel{
			Id:        1,
			Status:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:    "Fake restaurant",
		Address: "Fake address",
	}, nil
}

func GetRestaurant(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantStorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantBiz.NewGetRestaurantBiz(store)

		//fakeStore := fakeGetDataStore{}
		//biz := restaurantBiz.NewGetRestaurantBiz(fakeStore)

		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

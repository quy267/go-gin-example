package restaurantbiz

import (
	"context"
	restaurantModel "go-gin-example/modules/restaurant/model"
)

type GetRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantModel.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*restaurantModel.Restaurant, error) {
	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return result, nil
}

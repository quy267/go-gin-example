package restaurantbiz

import (
	"context"
	"go-gin-example/common"
	restaurantModel "go-gin-example/modules/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *restaurantModel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantModel.Restaurant, error)
}

type listNewRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listNewRestaurantBiz {
	return &listNewRestaurantBiz{store: store}
}

func (biz *listNewRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *restaurantModel.Filter,
	paging *common.Paging,
) ([]restaurantModel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}

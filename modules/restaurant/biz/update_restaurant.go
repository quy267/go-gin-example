package restaurantbiz

import (
	"context"
	"errors"
	"go-gin-example/common"
	restaurantModel "go-gin-example/modules/restaurant/model"
)

type UpdateRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantModel.Restaurant, error)

	Update(
		ctx context.Context,
		cond map[string]interface{},
		updateData *restaurantModel.RestaurantUpdate,
	) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantModel.RestaurantUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err == common.ErrDataNotFound {
			return errors.New("data not found")
		}

		return err
	}

	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}

	if err := biz.store.Update(ctx, map[string]interface{}{"id": id}, data); err != nil {
		return err
	}

	return nil
}

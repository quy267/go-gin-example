package restaurantstorage

import (
	"context"
	restaurantModel "go-gin-example/modules/restaurant/model"
)

func (s *sqlStore) Update(
	ctx context.Context,
	cond map[string]interface{},
	updateData *restaurantModel.RestaurantUpdate,
) error {
	db := s.db

	if err := db.Where(cond).Updates(updateData).Error; err != nil {
		return err
	}

	return nil
}

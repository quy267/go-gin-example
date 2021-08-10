package restaurantstorage

import (
	"context"
	restaurantModel "go-gin-example/modules/restaurant/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantModel.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

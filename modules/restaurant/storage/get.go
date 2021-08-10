package restaurantstorage

import (
	"context"
	"go-gin-example/common"
	restaurantModel "go-gin-example/modules/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	cond map[string]interface{},
	moreKeys ...string,
) (*restaurantModel.Restaurant, error) {
	db := s.db

	var data restaurantModel.Restaurant

	if err := db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrDataNotFound
		}

		return nil, err
	}

	return &data, nil
}

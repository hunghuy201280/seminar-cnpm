package userstorage

import (
	"auth-service/common"
	"auth-service/common/entitycommon"
	"auth-service/modules/user/usermodel"
	"context"
)

func (s *sqlStore) UpdateUser(ctx context.Context, data *usermodel.UserUpdate) (int, error) {
	updateResult := s.db.Model(&data).Updates(&data)
	if err := updateResult.Error; err != nil {
		return entitycommon.ERROR_ENTITY_ID, common.ErrDB(err)
	}

	return data.Id, nil

}

package userstorage

import (
	"auth-service/common"
	"auth-service/modules/user/usermodel"
	"context"
)

func (s *sqlStore) UpdateAvatar(ctx context.Context, data *usermodel.UpdateAvatarRequest, userId int) error {
	updateResult := s.db.Table(usermodel.User{}.TableName()).
		Where(common.JS{"id": userId}).
		Update("avatar_url", data.ImageUrl)

	if err := updateResult.Error; err != nil {
		return common.ErrDB(err)
	}

	return nil

}

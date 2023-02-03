package userbiz

import (
	"auth-service/common"
	"auth-service/modules/user/usermodel"
	"context"
)

type UpdateProfileStore interface {
	UpdateUser(ctx context.Context, data *usermodel.UserUpdate) (int, error)
}

type updateProfileBiz struct {
	updateProfileStore UpdateProfileStore
}

func NewUpdateProfileBiz(updateProfileStore UpdateProfileStore) *updateProfileBiz {
	return &updateProfileBiz{
		updateProfileStore: updateProfileStore,
	}
}

func (biz *updateProfileBiz) UpdateUser(ctx context.Context, data *usermodel.UserUpdate) (err error) {

	if _, err := biz.updateProfileStore.UpdateUser(ctx, data); err != nil {
		return common.ErrCannotUpdateEntity("User", err)
	}

	return nil

}

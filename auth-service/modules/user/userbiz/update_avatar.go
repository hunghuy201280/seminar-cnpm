package userbiz

import (
	"auth-service/common"
	"auth-service/modules/user/usermodel"
	"context"
	"errors"
	"github.com/gookit/goutil"
)

type UpdateAvatarStore interface {
	UpdateAvatar(ctx context.Context, data *usermodel.UpdateAvatarRequest, userId int) error
	FindOneUser(ctx context.Context, data *usermodel.UserFind) (*usermodel.UserLite, error)
}

type updateAvatarBiz struct {
	updateAvatarStore UpdateAvatarStore
}

func NewUpdateAvatarBiz(updateAvatarStore UpdateAvatarStore) *updateAvatarBiz {
	return &updateAvatarBiz{
		updateAvatarStore: updateAvatarStore,
	}
}

func (biz *updateAvatarBiz) UpdateAvatar(ctx context.Context, data *usermodel.UpdateAvatarRequest, userId int) (err error) {
	if goutil.IsNil(data) || goutil.IsEmpty(data.ImageUrl) {
		return common.ErrCannotUpdateEntity("User", errors.New("data is empty"))
	}
	user, err := biz.updateAvatarStore.FindOneUser(ctx, &usermodel.UserFind{
		SQLModel: common.SQLModel{Id: userId},
	})

	if err != nil || user == nil {
		return common.ErrEntityNotFound("User", err)

	}

	if err := biz.updateAvatarStore.UpdateAvatar(ctx, data, userId); err != nil {
		return common.ErrCannotUpdateEntity("User", err)
	}

	return nil

}

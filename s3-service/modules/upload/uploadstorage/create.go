package uploadstorage

import (
	"context"
	"s3-service/common"
)

func (store *sqlStore) CreateImage(ctx context.Context, data *common.File) error {
	db := store.db
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

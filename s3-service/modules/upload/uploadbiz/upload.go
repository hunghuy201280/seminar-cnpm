package uploadbiz

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"path/filepath"
	"s3-service/common"
	"s3-service/component/uploadprovider"
	"s3-service/modules/upload/uploadmodel"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.File) error
}

type uploadBiz struct {
	store    CreateImageStorage
	provider uploadprovider.UploadProvider
}

func NewUploadBiz(provider uploadprovider.UploadProvider, store CreateImageStorage) *uploadBiz {
	return &uploadBiz{store: store, provider: provider}
}

func (biz *uploadBiz) Upload(
	ctx context.Context, data []byte,
	folder, fileName string, userId int,
) (*common.File, error) {

	if strings.TrimSpace(folder) == "" {
		folder = "files"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	file, err := biz.provider.SaveFileUploaded(ctx, data,
		fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}
	file.Extension = fileExt
	file.OwnerId = userId
	file.FileName = fileName

	if err := biz.store.CreateImage(ctx, file); err != nil {
		//delete image on S3
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	return file, nil

}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		img, err = jpeg.DecodeConfig(reader)
		if err != nil {
			return 0, 0, err
		}
		return img.Width, img.Height, nil

	}
	return img.Width, img.Height, nil
}

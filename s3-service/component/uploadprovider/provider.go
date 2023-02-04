package uploadprovider

import (
	"context"
	"s3-service/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.File, error)
}

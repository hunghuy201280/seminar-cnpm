package uploadmodel

import (
	"fmt"
	"s3-service/common"
)

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot save file"),
		fmt.Sprintf("ErrCannotSaveFile"),
	)
}

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("File is not image"),
		fmt.Sprintf("ErrFileIsNotImage"),
	)
}

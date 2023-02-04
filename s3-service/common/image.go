package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type File struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	FileName  string `json:"file_name" gorm:"column:file_name;"`
	OwnerId   int    `json:"omitempty" gorm:"column:owner_id;"`
	Extension string `json:"extension" gorm:"column:extension;"`
}

func (j File) TableName() string {
	return "files"
}

func (j *File) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	var img File
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

func (j *File) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

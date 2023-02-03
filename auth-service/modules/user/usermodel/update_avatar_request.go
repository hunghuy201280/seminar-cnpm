package usermodel

type UpdateAvatarRequest struct {
	ImageUrl string `json:"image_url" form:"image_url"`
}

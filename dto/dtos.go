package dto

import "bitbucket.org/go-webservice/model"

type UserResponseBody struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserFavouriteResponseBody struct {
	UserID int           `json:"user_id"`
	Assets []model.Asset `json:"favourite_assets"`
}

func (u *UserFavouriteResponseBody) AppendChild(a model.Asset) {
	u.Assets = append(u.Assets, a)
}

// password should be omitted
func (dto *UserResponseBody) ConvertUserEntityToDTO(model model.User) *UserResponseBody {
	return &UserResponseBody{ID: model.ID, Username: model.Username, Email: model.Email}
}

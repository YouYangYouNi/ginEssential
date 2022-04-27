package dto

import "essential/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.Users) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}

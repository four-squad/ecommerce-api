package data

import (
	"ecommerce/features/user"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Avatar   string
	Name     string
	Email    string
	Address  string
	Password string
}

func ToCore(data Users) user.Core {
	return user.Core{
		ID:       data.ID,
		Avatar:   data.Avatar,
		Name:     data.Name,
		Email:    data.Email,
		Address:  data.Address,
		Password: data.Password,
	}
}

func CoreToData(data user.Core) Users {
	return Users{
		Model:    gorm.Model{ID: data.ID},
		Avatar:   data.Avatar,
		Name:     data.Name,
		Email:    data.Email,
		Address:  data.Address,
		Password: data.Password,
	}
}
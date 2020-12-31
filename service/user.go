package service

import (
	"github.com/gustavohmsilva/test-dependency-injection/datasource"
	"github.com/gustavohmsilva/test-dependency-injection/model"
)

// UserService ...
type UserService struct {
	dataSource datasource.UserDataSource
}

// NewUserService ...
func NewUserService(ds datasource.UserDataSource) UserService {
	return UserService{dataSource: ds}
}

// GetLatestUser ...
func (u UserService) GetLatestUser() (model.User, error) {
	ru, err := u.dataSource.SelectLatestUser()
	if err != nil {
		return model.User{}, err
	}
	return ru, nil
}

// InsertUser ...
func (u UserService) InsertUser(nu model.User) error {
	err := u.dataSource.InsertUser(nu)
	return err
}

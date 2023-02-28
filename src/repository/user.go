package repository

import (
	"todo-backend/database/sql"
	"todo-backend/sdk/errors"
	"todo-backend/src/entity"
	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	Get(ctx *gin.Context, params entity.UserParam) (entity.User, error)
	GetList(ctx *gin.Context, params entity.UserParam) ([]entity.User, error)
	Create(ctx *gin.Context, user entity.User) (entity.User, error)
	Update(ctx *gin.Context, userParam entity.UserParam) error
	Delete(ctx *gin.Context, userParam entity.UserParam) error
}

type user struct {
	db sql.DB
}

func InitUser(db sql.DB) UserInterface {
	return &user{db: db}
}

func (u *user) Get(ctx *gin.Context, params entity.UserParam) (entity.User, error) {
	var user entity.User

	whereClause := u.db.GetWhereClause(params)

	res := u.db.GetDB(ctx).Where(whereClause).First(&user)
	if res.RowsAffected == 0 {
		return user, errors.NotFound("User")
	} else if res.Error != nil {
		return user, res.Error
	}

	return user, nil
}

func (u *user) GetList(ctx *gin.Context, params entity.UserParam) ([]entity.User, error) {
	var users []entity.User

	whereClause := u.db.GetWhereClause(params)
	if err := u.db.GetDB(ctx).
		Limit(params.Limit).Offset(params.Offset).Where(whereClause).Find(&users).Error; err != nil {
		return users, err
	}

	if len(users) == 0 {
		return users, errors.NotFound("Users")
	}

	return users, nil
}

func (u *user) Create(ctx *gin.Context, user entity.User) (entity.User, error) {
	if err := u.db.GetDB(ctx).Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *user) Update(ctx *gin.Context, userParam entity.UserParam) error {
	whereClause := u.db.GetWhereClause(userParam)

	res := u.db.GetDB(ctx).Model(&entity.User{}).Where(whereClause).Updates(userParam)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.NotFound("User")
	}

	return nil
}

func (u *user) Delete(ctx *gin.Context, userParam entity.UserParam) error {
	whereClause := u.db.GetWhereClause(userParam)

	res := u.db.GetDB(ctx).Model(&entity.User{}).Where(whereClause).Delete(&entity.User{})
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.NotFound("User")
	}

	return nil
}

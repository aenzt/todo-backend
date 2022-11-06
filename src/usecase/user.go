package usecase

import (
	"dantal-backend/sdk/errors"
	"dantal-backend/sdk/jwt"
	"dantal-backend/sdk/password"
	"dantal-backend/src/entity"
	"dantal-backend/src/repository"
	"github.com/gin-gonic/gin"
)

type UserInterface interface {
	Login(ctx *gin.Context, userParam entity.UserParam, userInput entity.UserLoginInputParam) (entity.UserLoginResponse, error)
}

type User struct {
	userRepo repository.UserInterface
	roleRepo repository.RoleInterface
}

func InitUser(ur repository.UserInterface, rr repository.RoleInterface) UserInterface {
	return &User{
		userRepo: ur,
		roleRepo: rr,
	}
}

func (u *User) Login(ctx *gin.Context, userParam entity.UserParam, userInput entity.UserLoginInputParam) (entity.UserLoginResponse, error) {
	var userResponse entity.UserLoginResponse

	userParam.Username = userInput.Username
	user, err := u.userRepo.Get(ctx, userParam)
	if err != nil {
		return userResponse, err
	}

	if !password.Compare(user.Password, userInput.Password) {
		return userResponse, errors.NewWithCode(401, "Wrong password", "HTTPStatusUnauthorized")
	}

	role, err := u.roleRepo.Get(ctx, entity.RoleParam{ID: user.RoleID})
	if err != nil {
		return userResponse, err
	}

	user.RoleName = role.Name

	token, err := jwt.GetToken(user)
	if err != nil {
		return userResponse, errors.NewWithCode(500, "Failed to generate token", "HTTPStatusInternalServerError")
	}

	userResponse.User = user
	userResponse.Token = token

	return userResponse, nil
}

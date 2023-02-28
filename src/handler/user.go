package handler

import (
	"todo-backend/src/entity"
	"github.com/gin-gonic/gin"
)

func (r *rest) Login(ctx *gin.Context) {
	var userParam entity.UserParam
	var userInput entity.UserLoginInputParam

	if err := r.BindParam(ctx, &userParam); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	if err := r.BindBody(ctx, &userInput); err != nil {
		ErrorResponse(ctx, err)
		return
	}

	userResponse, err := r.uc.User.Login(ctx, userParam, userInput)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, "Login success", userResponse)
}

func (r *rest) GetUserProfile(ctx *gin.Context) {
	user := ctx.MustGet("user")

	SuccessResponse(ctx, "Get user profile success", user)
}

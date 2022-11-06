package repository

import (
	"dantal-backend/database/sql"
	"dantal-backend/sdk/errors"
	"dantal-backend/src/entity"
	"github.com/gin-gonic/gin"
)

type RoleInterface interface {
	Get(ctx *gin.Context, params entity.RoleParam) (entity.Role, error)
	GetList(ctx *gin.Context, params entity.RoleParam) ([]entity.Role, error)
}

type role struct {
	db sql.DB
}

func InitRole(db sql.DB) RoleInterface {
	return &role{db: db}
}

func (r *role) Get(ctx *gin.Context, params entity.RoleParam) (entity.Role, error) {
	var role entity.Role

	whereClause := r.db.GetWhereClause(params)

	if err := r.db.GetDB(ctx).Where(whereClause).First(&role).Error; err != nil {
		return role, err
	}

	return role, nil
}

func (r *role) GetList(ctx *gin.Context, params entity.RoleParam) ([]entity.Role, error) {
	var roles []entity.Role

	whereClause := r.db.GetWhereClause(params)
	if err := r.db.GetDB(ctx).
		Limit(params.Limit).Offset(params.Offset).Where(whereClause).Find(&roles).Error; err != nil {
		return roles, err
	}

	if len(roles) == 0 {
		return roles, errors.NotFound("Roles")
	}

	return roles, nil
}

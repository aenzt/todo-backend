package sql

import (
	"os"

	"dantal-backend/sdk/password"
	"dantal-backend/src/entity"
	"gorm.io/gorm"
)

type Migration struct {
	Db *gorm.DB
}

func (m *Migration) RunMigration() {
	m.Db.AutoMigrate(
		&entity.Role{},
		&entity.User{},
	)

	if !m.isUserExist() {
		m.runSeed()
	}
}

func (m *Migration) isUserExist() bool {
	var count int64
	if err := m.Db.Model(&entity.User{}).Count(&count).Error; err != nil {
		return false
	}

	return count > 0
}

func (m *Migration) runSeed() error {
	if err := m.runRoleSeed(); err != nil {
		return err
	}

	return m.runUserSeed()
}

func (m *Migration) runRoleSeed() error {
	roles := []entity.Role{
		{Name: "Super Admin"},
		{Name: "Admin"},
		{Name: "User"},
	}

	if err := m.Db.Create(&roles).Error; err != nil {
		return err
	}

	return nil
}

func (m *Migration) runUserSeed() error {
	adminPassword, err := password.Hash(os.Getenv("ADMIN_PASSWORD"))
	if err != nil {
		return err
	}

	user := entity.User{
		Name:     os.Getenv("ADMIN_NAME"),
		Username: os.Getenv("ADMIN_USERNAME"),
		Password: adminPassword,
		RoleID:   entity.SuperAdminID,
		IsMale:   true,
	}

	if err := m.Db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

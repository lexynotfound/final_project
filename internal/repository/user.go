package repository

import (
	"context"

	"github.com/lexynotfound/project_akhir/internal/infrastructure"
	"github.com/lexynotfound/project_akhir/models"
	"gorm.io/gorm"
)

type UserQuery interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUsersByID(ctx context.Context, id uint64) (models.User, error)

	DeleteUsersByID(ctx context.Context, id uint64) error
	CreateUser(ctx context.Context, user models.User) (models.User, error)
}

type UserCommand interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
}

type userQueryImpl struct {
	db infrastructure.GormPostgres
}

func NewUserQuery(db infrastructure.GormPostgres) UserQuery {
	return &userQueryImpl{db: db}
}

func (u *userQueryImpl) GetUsers(ctx context.Context) ([]models.User, error) {
	db := u.db.GetConnection()
	users := []models.User{}
	if err := db.
		WithContext(ctx).
		Table("users").
		Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userQueryImpl) GetUsersByID(ctx context.Context, id uint64) (models.User, error) {
	db := u.db.GetConnection()
	users := models.User{}
	if err := db.
		WithContext(ctx).
		Table("users").
		Where("id = ?", id).
		Find(&users).Error; err != nil {
		// if user not found, return nil error
		if err == gorm.ErrRecordNotFound {
			return models.User{}, nil
		}

		return models.User{}, err
	}
	return users, nil
}

func (u *userQueryImpl) DeleteUsersByID(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("users").
		Delete(&models.User{ID: id}).
		Error; err != nil {
		return err
	}
	return nil
}

func (u *userQueryImpl) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("users").
		Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

package repository

import (
	"context"
	// "github.com/mci-its/backend-service/dto"
	"github.com/mci-its/backend-service/entity"
	"gorm.io/gorm"
)

type (
	ViolationActionRepository interface {
		CreateViolationAction(ctx context.Context, tx *gorm.DB, violationAction entity.ViolationAction) (entity.ViolationAction, error)
	}

	violationActionRepository struct {
		db *gorm.DB
	}
)

func NewViolationActionRepository(db *gorm.DB) ViolationActionRepository {
	return &violationActionRepository{
		db: db,
	}
}

func (r *violationActionRepository) CreateViolationAction(ctx context.Context, tx *gorm.DB, violationAction entity.ViolationAction) (entity.ViolationAction, error) {
    if tx == nil {
        tx = r.db
    }

    // Create the violation action in the database
    if err := tx.WithContext(ctx).Create(&violationAction).Error; err != nil {
        return entity.ViolationAction{}, err
    }
    return violationAction, nil
}
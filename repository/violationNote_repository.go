package repository

import (
	"context"
	// "github.com/mci-its/backend-service/dto"
	"github.com/mci-its/backend-service/entity"
	"gorm.io/gorm"
)

type (
	ViolationNoteRepository interface {
		CreateViolationNotes(ctx context.Context, tx *gorm.DB, violationNotes entity.ViolationNote) (entity.ViolationNote, error)
	}

	violationNoteRepository struct {
		db *gorm.DB
	}
)

func NewViolationNoteRepository(db *gorm.DB) ViolationNoteRepository {
	return &violationNoteRepository{
		db: db,
	}
}

func (r *violationNoteRepository) CreateViolationNotes(ctx context.Context, tx *gorm.DB, violationNote entity.ViolationNote) (entity.ViolationNote, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&violationNote).Error; err != nil {
		return entity.ViolationNote{}, err
	}
	return violationNote, nil
}

package repository

import (
	"context"
	"math"
	

	"github.com/mci-its/backend-service/dto"
	"github.com/mci-its/backend-service/entity"
	"gorm.io/gorm"
)

type (
	// ViolationRepository interface for data access layer
	ViolationRepository interface {
		CreateViolation(ctx context.Context, tx *gorm.DB, violation entity.Violation) (entity.Violation, error)
		GetAllViolationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllViolationRepositoryResponse, error)
		GetViolationById(ctx context.Context, tx *gorm.DB, pkid int64) (entity.Violation, error)
		GetViolationByStatus(ctx context.Context, tx *gorm.DB, status string) (entity.Violation, error)
	}

	violationRepository struct {
		db *gorm.DB
	}
)

func NewViolationRepository(db *gorm.DB) ViolationRepository {
	return &violationRepository{
		db: db,
	}
}

func (r *violationRepository) CreateViolation(ctx context.Context, tx *gorm.DB, violation entity.Violation) (entity.Violation, error) {
	// Check if transaction is nil
	if tx == nil {
		// If nil, use the default db
		tx = r.db
	}

	// Create the violation
	if err := tx.WithContext(ctx).Create(&violation).Error; err != nil {
		return entity.Violation{}, err
	}

	return violation, nil
}

func (r *violationRepository) GetAllViolationWithPagination(ctx context.Context, tx *gorm.DB, req dto.PaginationRequest) (dto.GetAllViolationRepositoryResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var violations []entity.Violation
	var err error
	var count int64

	if req.PerPage == 0 {
		req.PerPage = 10
	}

	if req.Page == 0 {
		req.Page = 1
	}
	// Count the total of violations
	if err := tx.WithContext(ctx).Model(&entity.Violation{}).Count(&count).Error; err != nil {
		return dto.GetAllViolationRepositoryResponse{}, err
	}
	// Get all violations with pagination
	if err := tx.WithContext(ctx).Scopes(Paginate(req.Page, req.PerPage)).Find(&violations).Error; err != nil {
		return dto.GetAllViolationRepositoryResponse{}, err
	}

	totalPage := int64(math.Ceil(float64(count) / float64(req.PerPage)))

	return dto.GetAllViolationRepositoryResponse{
		Violations: violations,
		PaginationResponse: dto.PaginationResponse{
			Page:    req.Page,
			PerPage: req.PerPage,
			Count:   count,
			MaxPage: totalPage,
		},
	}, err
}

func (r *violationRepository) GetViolationById(ctx context.Context, tx *gorm.DB, pk_id int64) (entity.Violation, error) {
	if tx == nil {
		tx = r.db
	}

	var violation entity.Violation
	if err := tx.WithContext(ctx).Where("pk_id= ?", pk_id).Take(&violation).Error; err != nil {
		return entity.Violation{}, err
	}

	return violation, nil
}

func (r *violationRepository) GetViolationByStatus(ctx context.Context, tx *gorm.DB, status string) (entity.Violation, error) {
	if tx == nil {
		tx = r.db
	}

	var violation entity.Violation
	if err := tx.WithContext(ctx).Where("status = ?", status).Take(&violation).Error; err != nil {
		return entity.Violation{}, err
	}

	return violation, nil
}
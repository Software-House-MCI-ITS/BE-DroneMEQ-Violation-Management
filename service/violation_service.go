package service

import (
	// "bytes"
	"context"
	// "fmt"
	// "html/template"
	// "os"
	// "strings"
	// "sync"
	// "time"

	// "github.com/google/uuid"
	// "github.com/mci-its/backend-service/constants"
	"github.com/mci-its/backend-service/dto"
	"github.com/mci-its/backend-service/entity"

	// "github.com/mci-its/backend-service/helpers"
	"github.com/mci-its/backend-service/repository"
	// "github.com/mci-its/backend-service/utils"
)

type (
	ViolationService interface {
		CreateViolation(ctx context.Context, req dto.ViolationCreateRequest) (dto.ViolationResponse, error)
		GetAllViolationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.ViolationPaginationResponse, error)
		GetViolationById(ctx context.Context, pkid int64) (dto.ViolationResponse, error)
		GetViolationByStatus(ctx context.Context, status string) (dto.ViolationResponse, error)
	}

	violationService struct {
		violationRepo repository.ViolationRepository
	}
)

func NewViolationService(violationRepo repository.ViolationRepository) ViolationService {
	return &violationService{
		violationRepo: violationRepo,
	}
}

func (s *violationService) CreateViolation(ctx context.Context, req dto.ViolationCreateRequest) (dto.ViolationResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	violation := entity.Violation{
		ReportedByPKID: req.ReportedByPKID,
		Latitude:       req.Latitude,
		Longitude:      req.Longitude,
		ViolationType:  req.ViolationType,
		Severity:       req.Severity,
		Description:    req.Description,
	}

	violationCreate, err := s.violationRepo.CreateViolation(ctx, nil, violation)
	if err != nil {
		return dto.ViolationResponse{}, err
	}

	return dto.ViolationResponse{
		PKID:           violationCreate.PKID,
		ReportedByPKID: violationCreate.ReportedByPKID,
		Latitude:       violationCreate.Latitude,
		Longitude:      violationCreate.Longitude,
		ViolationType:  violationCreate.ViolationType,
		Severity:       violationCreate.Severity,
		Description:    violationCreate.Description,
	}, nil
}

func (s *violationService) GetAllViolationWithPagination(ctx context.Context, req dto.PaginationRequest) (dto.ViolationPaginationResponse, error) {
	dataWithPaginate, err := s.violationRepo.GetAllViolationWithPagination(ctx, nil, req)
	if err != nil {
		return dto.ViolationPaginationResponse{}, err
	}

	var datas []dto.ViolationResponse
	for _, violation := range dataWithPaginate.Violations {
		data := dto.ViolationResponse{
			PKID:           violation.PKID,
			ReportedByPKID: violation.ReportedByPKID,
			Latitude:       violation.Latitude,
			Longitude:      violation.Longitude,
			ViolationType:  violation.ViolationType,
			Severity:       violation.Severity,
			Description:    violation.Description,
		}
		datas = append(datas, data)
	}

	return dto.ViolationPaginationResponse{
		Data: datas,
		PaginationResponse: dto.PaginationResponse{
			Page:    dataWithPaginate.Page,
			PerPage: dataWithPaginate.PerPage,
			MaxPage: dataWithPaginate.MaxPage,
			Count:   dataWithPaginate.Count,
		},
	}, nil
}

func (s *violationService) GetViolationById(ctx context.Context, pkid int64) (dto.ViolationResponse, error) {
	violation, err := s.violationRepo.GetViolationById(ctx, nil, pkid)
	if err != nil {
		return dto.ViolationResponse{}, err
	}

	return dto.ViolationResponse{
		PKID:           violation.PKID,
		ReportedByPKID: violation.ReportedByPKID,
		Latitude:       violation.Latitude,
		Longitude:      violation.Longitude,
		ViolationType:  violation.ViolationType,
		Severity:       violation.Severity,
		Description:    violation.Description,
	}, nil
}

func (s *violationService) GetViolationByStatus(ctx context.Context, status string) (dto.ViolationResponse, error) {
	violation, err := s.violationRepo.GetViolationByStatus(ctx, nil, status)
	if err != nil {
		return dto.ViolationResponse{}, err
	}

	return dto.ViolationResponse{
		PKID:           violation.PKID,
		ReportedByPKID: violation.ReportedByPKID,
		Latitude:       violation.Latitude,
		Longitude:      violation.Longitude,
		ViolationType:  violation.ViolationType,
		Severity:       violation.Severity,
		Description:    violation.Description,
	}, nil
}
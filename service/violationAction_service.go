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
	ViolationActionService interface {
		CreateViolationAction(ctx context.Context, req dto.ViolationActionCreateRequest, pk_id int64) (dto.ViolationActionCreateResponse, error)
	}

	violationActionService struct {
		violationActionRepo repository.ViolationActionRepository
	}
)

func NewViolationActionService(violationActionRepo repository.ViolationActionRepository) ViolationActionService {
	return &violationActionService{
		violationActionRepo: violationActionRepo,
	}
}

func (s *violationActionService) CreateViolationAction(ctx context.Context, req dto.ViolationActionCreateRequest, ViolationID int64) (dto.ViolationActionCreateResponse, error) {
	violationAction := entity.ViolationAction{
		ViolationPKID: ViolationID,
		ActionTaken:  req.ActionTaken,
		Outcome: 	req.Outcome,
		TakenByPKID: req.TakenByPKID,
	}

	violationActionCreate, err := s.violationActionRepo.CreateViolationAction(ctx, nil, violationAction)
	if err != nil {
		return dto.ViolationActionCreateResponse{}, err
	}
	return dto.ViolationActionCreateResponse{
		PKID: violationActionCreate.PKID,
		ViolationPKID: violationActionCreate.ViolationPKID,
		ActionTaken: violationActionCreate.ActionTaken,
		Outcome: violationActionCreate.Outcome,
		TakenByPKID: violationActionCreate.TakenByPKID,
	}, nil
}
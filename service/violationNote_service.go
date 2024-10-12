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
	ViolationNoteService interface {
		CreateViolationNote(ctx context.Context, req dto.ViolationNoteCreateRequest, pk_id int64) (dto.ViolationNoteCreateResponse, error)
	}

	violationNoteService struct {
		violationNoteRepo repository.ViolationNoteRepository
	}
)

func NewViolationNoteService(violationNoteRepo repository.ViolationNoteRepository) ViolationNoteService {
	return &violationNoteService{
		violationNoteRepo: violationNoteRepo,
	}
}

func (s *violationNoteService) CreateViolationNote(ctx context.Context, req dto.ViolationNoteCreateRequest, ViolationID int64) (dto.ViolationNoteCreateResponse, error) {
	violationNote := entity.ViolationNote{
		ViolationPKID: ViolationID,
		NoteType:      req.NoteType,
		Note:          req.Note,
	}

	violationNoteCreate, err := s.violationNoteRepo.CreateViolationNotes(ctx, nil, violationNote)
	if err != nil {
		return dto.ViolationNoteCreateResponse{}, err
	}
	return dto.ViolationNoteCreateResponse{
		PKID:         violationNoteCreate.PKID,
		ViolationPKID: violationNoteCreate.ViolationPKID,
		NoteType:      violationNoteCreate.NoteType,
		Note:          violationNoteCreate.Note,
	}, nil
}
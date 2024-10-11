package dto

import (
	"time"

	"github.com/mci-its/backend-service/entity"
)

const (
    // Message Failed
    MESSAGE_FAILED_CREATE_VIOLATION = "Failed to create violation"
    MESSAGE_FAILED_GET_ALL_VIOLATION = "Failed to get all violation"
	MESSAGE_FAILED_GET_VIOLATION_BY_ID = "Failed to get violation by id"
	MESSAGE_FAILED_GET_VIOLATION_BY_STATUS = "Failed to get violation by status"
    // Message Success
    MESSAGE_SUCCESS_GET_ALL_VIOLATION = "Success to get all violation"
    MESSAGE_SUCCESS_CREATE_VIOLATION = "Success to create violation"
	MESSAGE_SUCCESS_GET_VIOLATION_BY_ID = "Success to get violation by id"
	MESSAGE_SUCCESS_GET_VIOLATION_BY_STATUS = "Success to get violation by status"
)

type GetAllViolationRepositoryResponse struct {
	Violations []entity.Violation
	PaginationResponse
}

type ViolationCreateRequest struct {
	ReportedByPKID int64   `json:"reported_by_pkid" form:"reported_by_pkid"`
	Latitude       float64 `json:"latitude" form:"latitude"`
	Longitude      float64 `json:"longitude" form:"longitude"`
	ViolationType  string  `json:"violation_type" form:"violation_type"`
	Severity       string  `json:"severity" form:"severity"`
	Description    string  `json:"description" form:"description"`
}

type ViolationResponse struct {
	PKID           int64   `json:"pkid"`
	ReportedByPKID int64   `json:"reported_by_pkid"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	ViolationType  string  `json:"violation_type"`
	Severity       string  `json:"severity"`
	Description    string  `json:"description"`
}

type ViolationPaginationResponse struct {
	Data []ViolationResponse `json:"data"`
	PaginationResponse
}

type Violation struct {
	PKID           int64     `gorm:"primary_key;auto_increment" json:"pkid"`
	ReportedByPKID int64     `gorm:"not null" json:"reported_by_pkid"`
	Latitude       float64   `gorm:"not null" json:"latitude"`
	Longitude      float64   `gorm:"not null" json:"longitude"`
	ReportedAt     time.Time `gorm:"not null" json:"reported_at"`
	ViolationType  string    `gorm:"type:varchar(255)" json:"violation_type"`
	Status         string    `gorm:"type:varchar(255)" json:"status"`
	Severity       string    `gorm:"type:varchar(255)" json:"severity"`
	Description    string    `gorm:"type:text" json:"description"`

	// Audit Columns
	OfficePKID  int64      `gorm:"not null" json:"office_pkid"`
	CreatedBy   string     `gorm:"size:255;not null" json:"created_by"`
	CreatedDate time.Time  `gorm:"not null" json:"created_date"`
	CreatedHost *string    `gorm:"size:255" json:"created_host"`
	UpdatedBy   *string    `gorm:"size:255" json:"updated_by,omitempty"`
	UpdatedDate *time.Time `json:"updated_date,omitempty"`
	UpdatedHost *string    `gorm:"size:255" json:"updated_host,omitempty"`
	IsDeleted   bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedBy   *string    `gorm:"size:255" json:"deleted_by,omitempty"`
	DeletedDate *time.Time `json:"deleted_date,omitempty"`
	DeletedHost *string    `gorm:"size:255" json:"deleted_host,omitempty"`
}

type ViolationAction struct {
	PKID          int64     `gorm:"primary_key;auto_increment" json:"pkid"`
	ViolationPKID int64     `gorm:"not null" json:"violation_pkid"`
	ActionTaken   string    `gorm:"type:varchar(255)" json:"action_taken"`
	TakenByPKID   int64     `gorm:"not null" json:"taken_by_pkid"`
	ActionTakenAt time.Time `gorm:"not null" json:"action_taken_at"`
	Outcome       string    `gorm:"type:varchar(255)" json:"outcome"`

	// Audit Columns
	OfficePKID  int64      `gorm:"not null" json:"office_pkid"`
	CreatedBy   string     `gorm:"size:255;not null" json:"created_by"`
	CreatedDate time.Time  `gorm:"not null" json:"created_date"`
	CreatedHost *string    `gorm:"size:255" json:"created_host"`
	UpdatedBy   *string    `gorm:"size:255" json:"updated_by,omitempty"`
	UpdatedDate *time.Time `json:"updated_date,omitempty"`
	UpdatedHost *string    `gorm:"size:255" json:"updated_host,omitempty"`
	IsDeleted   bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedBy   *string    `gorm:"size:255" json:"deleted_by,omitempty"`
	DeletedDate *time.Time `json:"deleted_date,omitempty"`
	DeletedHost *string    `gorm:"size:255" json:"deleted_host,omitempty"`
}

type ViolationNote struct {
	PKID          int64  `gorm:"primary_key;auto_increment" json:"pkid"`
	ViolationPKID int64  `gorm:"not null" json:"violation_pkid"`
	Note          string `gorm:"type:text" json:"note"`
	NoteType      string `gorm:"type:varchar(255)" json:"note_type"`

	// Audit Columns
	OfficePKID  int64      `gorm:"not null" json:"office_pkid"`
	CreatedBy   string     `gorm:"size:255;not null" json:"created_by"`
	CreatedDate time.Time  `gorm:"not null" json:"created_date"`
	CreatedHost *string    `gorm:"size:255" json:"created_host"`
	UpdatedBy   *string    `gorm:"size:255" json:"updated_by,omitempty"`
	UpdatedDate *time.Time `json:"updated_date,omitempty"`
	UpdatedHost *string    `gorm:"size:255" json:"updated_host,omitempty"`
	IsDeleted   bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedBy   *string    `gorm:"size:255" json:"deleted_by,omitempty"`
	DeletedDate *time.Time `json:"deleted_date,omitempty"`
	DeletedHost *string    `gorm:"size:255" json:"deleted_host,omitempty"`
}

type ViolationMedia struct {
	PKID           int64     `gorm:"primary_key;auto_increment" json:"pkid"`
	ViolationPKID  int64     `gorm:"not null" json:"violation_pkid"`
	MediaURL       string    `gorm:"type:varchar(255)" json:"media_url"`
	MediaType      string    `gorm:"type:varchar(255)" json:"media_type"`
	UploadedByPKID int64     `gorm:"not null" json:"uploaded_by_pkid"`
	UploadedAt     time.Time `gorm:"not null" json:"uploaded_at"`

	// Audit Columns
	OfficePKID  int64      `gorm:"not null" json:"office_pkid"`
	CreatedBy   string     `gorm:"size:255;not null" json:"created_by"`
	CreatedDate time.Time  `gorm:"not null" json:"created_date"`
	CreatedHost *string    `gorm:"size:255" json:"created_host"`
	UpdatedBy   *string    `gorm:"size:255" json:"updated_by,omitempty"`
	UpdatedDate *time.Time `json:"updated_date,omitempty"`
	UpdatedHost *string    `gorm:"size:255" json:"updated_host,omitempty"`
	IsDeleted   bool       `gorm:"not null;default:false" json:"is_deleted"`
	DeletedBy   *string    `gorm:"size:255" json:"deleted_by,omitempty"`
	DeletedDate *time.Time `json:"deleted_date,omitempty"`
	DeletedHost *string    `gorm:"size:255" json:"deleted_host,omitempty"`
}

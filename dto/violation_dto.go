package dto

import "time"

type Violation struct {
	PKID           int64     `gorm:"primary_key;auto_increment"`
	ReportedByPKID int64     `gorm:"not null"`
	Latitude       float64   `gorm:"not null"`
	Longitude      float64   `gorm:"not null"`
	ReportedAt     time.Time `gorm:"not null"`
	ViolationType  string    `gorm:"type:enum('ILLEGAL_FISHING', 'POLLUTION', 'HABITAT_DESTRUCTION', 'ASSET_MISSING', 'OTHER')"`
	Status         string    `gorm:"type:enum('REPORTED', 'UNDER_REVIEW', 'RESOLVED', 'DISMISSED')"`
	Severity       string    `gorm:"type:enum('LOW', 'MEDIUM', 'HIGH')"`
	Description    string    `gorm:"type:text"`

	// Audit Columns
	OfficePKID  int64     `gorm:"not null"`
	CreatedBy   string    `gorm:"size:255;not null"`
	CreatedDate time.Time `gorm:"not null"`
	CreatedHost *string   `gorm:"size:255"`
	UpdatedBy   *string   `gorm:"size:255"`
	UpdatedDate *time.Time
	UpdatedHost *string `gorm:"size:255"`
	IsDeleted   bool    `gorm:"not null;default:false"`
	DeletedBy   *string `gorm:"size:255"`
	DeletedDate *time.Time
	DeletedHost *string `gorm:"size:255"`
}

type ViolationAction struct {
	PKID          int64     `gorm:"primary_key;auto_increment"`
	ViolationPKID int64     `gorm:"not null"`
	ActionTaken   string    `gorm:"type:varchar(255)"`
	TakenByPKID   int64     `gorm:"not null"`
	ActionTakenAt time.Time `gorm:"not null"`
	Outcome       string    `gorm:"type:enum('SUCCESSFUL', 'FAILED', 'PARTIAL')"`

	// Audit Columns
	OfficePKID  int64     `gorm:"not null"`
	CreatedBy   string    `gorm:"size:255;not null"`
	CreatedDate time.Time `gorm:"not null"`
	CreatedHost *string   `gorm:"size:255"`
	UpdatedBy   *string   `gorm:"size:255"`
	UpdatedDate *time.Time
	UpdatedHost *string `gorm:"size:255"`
	IsDeleted   bool    `gorm:"not null;default:false"`
	DeletedBy   *string `gorm:"size:255"`
	DeletedDate *time.Time
	DeletedHost *string `gorm:"size:255"`
}

type ViolationNote struct {
	PKID          int64     `gorm:"primary_key;auto_increment"`
	ViolationPKID int64     `gorm:"not null"`
	Note          string    `gorm:"type:text"`
	NoteType      string    `gorm:"type:enum('OBSERVATION', 'FOLLOWUP', 'RESOLUTION', 'OTHER')"`

	// Audit Columns
	OfficePKID  int64     `gorm:"not null"`
	CreatedBy   string    `gorm:"size:255;not null"`
	CreatedDate time.Time `gorm:"not null"`
	CreatedHost *string   `gorm:"size:255"`
	UpdatedBy   *string   `gorm:"size:255"`
	UpdatedDate *time.Time
	UpdatedHost *string `gorm:"size:255"`
	IsDeleted   bool    `gorm:"not null;default:false"`
	DeletedBy   *string `gorm:"size:255"`
	DeletedDate *time.Time
	DeletedHost *string `gorm:"size:255"`
}

type ViolationMedia struct {
	PKID          int64     `gorm:"primary_key;auto_increment"`
	ViolationPKID int64     `gorm:"not null"`
	MediaURL      string    `gorm:"type:varchar(255)"`
	MediaType     string    `gorm:"type:enum('IMAGE', 'VIDEO', 'DOCUMENT')"`
	UploadedByPKID int64    `gorm:"not null"`
	UploadedAt    time.Time `gorm:"not null"`

	// Audit Columns
	OfficePKID  int64     `gorm:"not null"`
	CreatedBy   string    `gorm:"size:255;not null"`
	CreatedDate time.Time `gorm:"not null"`
	CreatedHost *string   `gorm:"size:255"`
	UpdatedBy   *string   `gorm:"size:255"`
	UpdatedDate *time.Time
	UpdatedHost *string `gorm:"size:255"`
	IsDeleted   bool    `gorm:"not null;default:false"`
	DeletedBy   *string `gorm:"size:255"`
	DeletedDate *time.Time
	DeletedHost *string `gorm:"size:255"`
}

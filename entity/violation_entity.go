package entity

import "time"

type violationType string
type status string
type severity string
type outcome string
type noteType string
type mediaType string

const (
    // Violation Type
    IllegalFishing      violationType = "ILLEGAL_FISHING"
    Pollution           violationType = "POLLUTION"
    HabitatDestruction  violationType = "HABITAT_DESTRUCTION"
    AssetMissing        violationType = "ASSET_MISSING"
    OtherViolation      violationType = "OTHER"
)

const (
    // Status
    Reported    status = "REPORTED"
    UnderReview status = "UNDER_REVIEW"
    Resolved    status = "RESOLVED"
    Dismissed   status = "DISMISSED"
)

const (
    // Severity
    Low    severity = "LOW"
    Medium severity = "MEDIUM"
    High   severity = "HIGH"
)

const (
    // Outcome
    Successful outcome = "SUCCESSFUL"
    Failed     outcome = "FAILED"
    Partial    outcome = "PARTIAL"
)

const (
    // Note Type
    Observation noteType = "OBSERVATION"
    FollowUp    noteType = "FOLLOWUP"
    Resolution  noteType = "RESOLUTION"
    OtherNote   noteType = "OTHER"
)

const (
    // Media Type
    Image    mediaType = "IMAGE"
    Video    mediaType = "VIDEO"
    Document mediaType = "DOCUMENT"
)

type Violation struct {
    PKID           int64         `gorm:"primary_key;auto_increment" json:"pkid"`
    ReportedByPKID int64         `gorm:"not null" json:"reported_by_pkid"` // Field is now exported (starts with uppercase)
    Latitude       float64       `gorm:"not null" json:"latitude"`         // Fixed missing closing quote for json tag
    Longitude      float64       `gorm:"not null" json:"longitude"`        // Fixed missing closing quote for json tag
    ReportedAt     time.Time     `gorm:"not null" json:"reported_at"`      // Field is now exported (starts with uppercase)
    ViolationType  violationType `gorm:"not null" json:"violation_type"`
    Status         status        `gorm:"not null" json:"status"`
    Severity       severity      `gorm:"not null" json:"severity"`
    Description    string        `gorm:"type:text" json:"description"`
    
    // Audit Columns
    OfficePKID  int64      `gorm:"not null" json:"office_pkid"`
    CreatedBy   string     `json:"created_by"`
    CreatedDate time.Time  `json:"created_date"`
    CreatedHost string     `json:"created_host"`
    UpdatedBy   *string    `gorm:"size:255" json:"updated_by,omitempty"`
    UpdatedDate *time.Time `json:"updated_date,omitempty"`
    UpdatedHost *string    `gorm:"size:255" json:"updated_host,omitempty"`
    IsDeleted   bool       `gorm:"default:false" json:"is_deleted"`
    DeletedBy   *string    `json:"deleted_by,omitempty"`
    DeletedDate *time.Time `json:"deleted_date,omitempty"`
    DeletedHost *string    `json:"deleted_host,omitempty"`
}

type ViolationAction struct {
    PKID          int64      `gorm:"primary_key;auto_increment" json:"pkid"`
    ViolationPKID int64      `gorm:"not null" json:"violation_pkid"` // Field is now exported (starts with uppercase)
    ActionTaken   string     `gorm:"type:varchar(255)" json:"action_taken"` // Field is now exported (starts with uppercase)
    TakenByPKID   int64      `gorm:"not null" json:"taken_by_pkid"`
    ActionTakenAt time.Time  `json:"action_taken_at"`
    Outcome       outcome    `json:"outcome"`
    // Audit Column
    OfficePKID  int64      `gorm:"not null" json:"office_pkid"`
    CreatedBy   string     `json:"created_by"`
    CreatedDate time.Time  `json:"created_date"`
    CreatedHost string     `json:"created_host"`
    UpdatedBy   *string    `gorm:"size:255" json:"updated_by,omitempty"`
    UpdatedDate *time.Time `json:"updated_date,omitempty"`
    UpdatedHost *string    `gorm:"size:255" json:"updated_host,omitempty"`
    IsDeleted   bool       `gorm:"default:false" json:"is_deleted"`
    DeletedBy   *string    `json:"deleted_by,omitempty"`
    DeletedDate *time.Time `json:"deleted_date,omitempty"`
    DeletedHost *string    `json:"deleted_host,omitempty"`
}

type ViolationNote struct {
    PKID          int64      `gorm:"primary_key;auto_increment" json:"pkid"`
    ViolationPKID int64      `gorm:"not null" json:"violation_pkid"` // Field is now exported (starts with uppercase)
    Note          string     `gorm:"type:text" json:"note"` // Field is now exported (starts with uppercase)
    NoteType      noteType   `json:"note_type"`
    // Audit Column
    OfficePKID  int64      `gorm:"not null" json:"office_pkid"`
    CreatedBy   string     `json:"created_by"`
    CreatedDate time.Time  `json:"created_date"`
    CreatedHost string     `json:"created_host"`
    UpdatedBy   *string    `gorm:"size:255" json:"updated_by,omitempty"`
    UpdatedDate *time.Time `json:"updated_date,omitempty"`
    UpdatedHost *string    `gorm:"size:255" json:"updated_host,omitempty"`
    IsDeleted   bool       `gorm:"default:false" json:"is_deleted"`
    DeletedBy   *string    `json:"deleted_by,omitempty"`
    DeletedDate *time.Time `json:"deleted_date,omitempty"`
    DeletedHost *string    `json:"deleted_host,omitempty"`
}

type ViolationMedia struct {
    PKID           int64      `gorm:"primary_key;auto_increment" json:"pkid"`
    ViolationPKID  int64      `gorm:"not null" json:"violation_pkid"` // Field is now exported (starts with uppercase)
    MediaURL       string     `json:"media_url"` // Field is now exported (starts with uppercase)
    MediaType      mediaType  `json:"media_type"`
    UploadedByPKID int64      `gorm:"not null" json:"uploaded_by_pkid"`
    UploadedAt     time.Time  `json:"uploaded_at"`
    // Audit Column
    OfficePKID  int64      `gorm:"not null" json:"office_pkid"`
    CreatedBy   string     `json:"created_by"`
    CreatedDate time.Time  `json:"created_date"`
    CreatedHost string     `json:"created_host"`
    UpdatedBy   *string    `gorm:"size:255" json:"updated_by,omitempty"`
    UpdatedDate *time.Time `json:"updated_date,omitempty"`
    UpdatedHost *string    `gorm:"size:255" json:"updated_host,omitempty"`
    IsDeleted   bool       `gorm:"default:false" json:"is_deleted"`
    DeletedBy   *string    `json:"deleted_by,omitempty"`
    DeletedDate *time.Time `json:"deleted_date,omitempty"`
    DeletedHost *string    `json:"deleted_host,omitempty"`
}
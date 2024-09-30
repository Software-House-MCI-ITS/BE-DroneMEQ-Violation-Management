package seeds

import (
    "encoding/json"
    "errors"
    "io"
    "os"

    "github.com/mci-its/backend-service/entity"
    "gorm.io/gorm"
)

func ListViolationNoteSeeder(db *gorm.DB) error {
    jsonFile, err := os.Open("./migrations/json/violation_note.json")
    if err != nil {
        return err
    }
    defer jsonFile.Close()

    jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
        return err
    }

    var listViolationNote []entity.ViolationNote
    if err := json.Unmarshal(jsonData, &listViolationNote); err != nil {
        return err
    }

    hasTable := db.Migrator().HasTable(&entity.ViolationNote{})
    if !hasTable {
        if err := db.Migrator().CreateTable(&entity.ViolationNote{}); err != nil {
            return err
        }
    }

    for _, data := range listViolationNote {
        var violationNote entity.ViolationNote
        err := db.Where(&entity.ViolationNote{PKID: data.PKID}).First(&violationNote).Error
        if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
            return err
        }

        isData := db.Find(&violationNote, "pk_id = ?", data.PKID).RowsAffected
        if isData == 0 {
            if err := db.Create(&data).Error; err != nil {
                return err
            }
        }
    }

    return nil
}
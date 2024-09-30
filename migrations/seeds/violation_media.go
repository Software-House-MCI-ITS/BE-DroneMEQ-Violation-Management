package seeds

import (
    "encoding/json"
    "errors"
    "io"
    "os"

    "github.com/mci-its/backend-service/entity"
    "gorm.io/gorm"
)

func ListViolationMediaSeeder(db *gorm.DB) error {
    jsonFile, err := os.Open("./migrations/json/violation_media.json")
    if err != nil {
        return err
    }
    defer jsonFile.Close()

    jsonData, err := io.ReadAll(jsonFile)
    if err != nil {
        return err
    }

    var listViolationMedia []entity.ViolationMedia
    if err := json.Unmarshal(jsonData, &listViolationMedia); err != nil {
        return err
    }

    hasTable := db.Migrator().HasTable(&entity.ViolationMedia{})
    if !hasTable {
        if err := db.Migrator().CreateTable(&entity.ViolationMedia{}); err != nil {
            return err
        }
    }

    for _, data := range listViolationMedia {
        var violationMedia entity.ViolationMedia
        err := db.Where(&entity.ViolationMedia{PKID: data.PKID}).First(&violationMedia).Error
        if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
            return err
        }

        isData := db.Find(&violationMedia, "pk_id = ?", data.PKID).RowsAffected
        if isData == 0 {
            if err := db.Create(&data).Error; err != nil {
                return err
            }
        }
    }

    return nil
}
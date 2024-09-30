package seeds

import (
    "encoding/json"
    "errors"
    "io"
    "os"

    "github.com/mci-its/backend-service/entity"
    "gorm.io/gorm"
)

func ListViolationActionSeeder(db *gorm.DB) error {
    jsonFile, err := os.Open("./migrations/json/violation_action.json")
    if err != nil {
        return err
    }
    defer jsonFile.Close()

    jsonData, err := io.ReadAll(jsonFile)
    if err != nil {
        return err
    }

    var listViolationAction []entity.ViolationAction
    if err := json.Unmarshal(jsonData, &listViolationAction); err != nil {
        return err
    }

    hasTable := db.Migrator().HasTable(&entity.ViolationAction{})
    if !hasTable {
        if err := db.Migrator().CreateTable(&entity.ViolationAction{}); err != nil {
            return err
        }
    }

    for _, data := range listViolationAction {
        var violationAction entity.ViolationAction
        err := db.Where(&entity.ViolationAction{PKID: data.PKID}).First(&violationAction).Error
        if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
            return err
        }

        isData := db.Find(&violationAction, "pk_id = ?", data.PKID).RowsAffected
        if isData == 0 {
            if err := db.Create(&data).Error; err != nil {
                return err  
            }
        }
    }

    return nil
}
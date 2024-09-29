package seeds

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/mci-its/backend-service/entity"
	"gorm.io/gorm"
)

func ListViolationSeeder(db *gorm.DB) error {
	jsonFile, err := os.Open("./migrations/json/violation.json")
	if err != nil {
		return err
	}

	jsonData, _ := io.ReadAll(jsonFile)

	var listViolation []entity.Violation
	if err := json.Unmarshal(jsonData, &listViolation); err != nil {
		return err
	}

	hasTable := db.Migrator().HasTable(&entity.Violation{})
	if !hasTable {
		if err := db.Migrator().CreateTable(&entity.Violation{}); err != nil {
			return err
		}
	}

	for _, data := range listViolation {
		var violation entity.Violation
		err := db.Where(&entity.Violation{PKID: data.PKID}).First(&violation).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	
		isData := db.Find(&violation, "PKID = ?", data.PKID).RowsAffected
		if isData == 0 {
			if err := db.Create(&data).Error; err != nil {
				return err
			}
		}
	}
	
	return nil

}

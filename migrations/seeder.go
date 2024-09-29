package migrations

import (
	"github.com/mci-its/backend-service/migrations/seeds"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) error {
	if err := seeds.ListUserSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListViolationSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListViolationActionSeeder(db); err != nil {
		return err
	}
	
	if err := seeds.ListViolationMediaSeeder(db); err != nil {
		return err
	}

	if err := seeds.ListViolationNoteSeeder(db); err != nil {
		return err
	}


	return nil
}

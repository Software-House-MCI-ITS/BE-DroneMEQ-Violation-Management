package migrations

import (
    "log"

    "github.com/mci-its/backend-service/entity"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
    if err := db.AutoMigrate(
        &entity.Violation{},
        &entity.ViolationAction{},
        &entity.ViolationNote{},
        &entity.ViolationMedia{},
		&entity.User{},
    ); err != nil {
        return err
    }

    log.Println("migration completed successfully")

    return nil
}

func dropAllTables(db *gorm.DB) error {
    tables := []interface{}{
        &entity.Violation{},
        &entity.ViolationAction{},
        &entity.ViolationNote{},
        &entity.ViolationMedia{},
		&entity.User{},
    }

    for _, table := range tables {
        if db.Migrator().HasTable(table) {
            if err := db.Migrator().DropTable(table); err != nil {
                return err
            }
            log.Printf("Dropped table for %T successfully\n", table)
        } else {
            log.Printf("Table for %T does not exist\n", table)
        }
    }

    log.Println("drop all tables completed successfully")
    return nil
}

func Fresh(db *gorm.DB) error {
	if err := dropAllTables(db); err != nil {
		log.Printf("Error dropping tables: %v", err)
		return err
	}

	if err := Migrate(db); err != nil {
		log.Printf("Error during migration: %v", err)
		return err
	}

	log.Println("Fresh migration completed successfully.")
	return nil
}
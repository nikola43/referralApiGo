package db

import (
	"fmt"
	"log"

	"github.com/nikola43/pdexrefapi/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GormDB *gorm.DB

func Migrate() {
	// DROP
	GormDB.Migrator().DropTable(&models.User{})
	GormDB.Migrator().DropTable(&models.Referral{})

	// CREATE
	GormDB.AutoMigrate(&models.User{})
	GormDB.AutoMigrate(&models.Referral{})
}

func InitializeDatabase(user, password, dbName, host, port string, migrate bool) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	GormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal(err)
	}

	if migrate {
		Migrate()
	}
}

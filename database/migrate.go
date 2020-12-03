package database

// import "gorm.io/gorm"
import (
	"go-crawler/models/dbmodels"

	"gorm.io/gorm"
	// 	_ "github.com/go-sql-driver/mysql"
	// 	"github.com/jinzhu/gorm"
	// 	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func MigrateUser(db *gorm.DB) {

	db.Debug().AutoMigrate(&dbmodels.User{}) //Model or Struct

	//Drops table if already exists
	// db.Debug().DropTableIfExists(&dbmodels.User{})
	//Auto create table based on Model
	db.Debug().AutoMigrate(&dbmodels.User{})
}

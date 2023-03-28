package bootstrap

import (
	"fmt"
	"fpt/global"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDB() {
	//根据驱动配置进行初始化
	switch global.Ftpserver.Config.Database.Database {
	case "sqlite":
		global.Ftpserver.DB = initSqliteGorm()
	default:
		return
	}
}

func initSqliteGorm() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(global.Ftpserver.Config.Database.Database), &gorm.Config{})
	if err != nil {
		fmt.Println("Faild to connect database: ", err)
	}

	db.AutoMigrate(&global.FileInfo{})
	return db
}

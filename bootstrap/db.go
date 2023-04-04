package bootstrap

import (
	"fmt"
	"fpt/global"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//初始化数据库
func InitializeDB() {
	//根据驱动配置进行初始化
	switch global.Ftpserver.Config.Database.Driver {
	case "sqlite":
		global.Ftpserver.DB = initSqliteGorm()
	default:
		return
	}
}

//初始化sqlite数据库
func initSqliteGorm() *gorm.DB {
	os.Mkdir(global.Ftpserver.Config.Database.DatabaseFile, os.ModePerm)
	os.Mkdir(global.Ftpserver.Config.Ftp.Storage, os.ModePerm)

	databasePath := global.Ftpserver.Config.Database.DatabaseFile + "/" + global.Ftpserver.Config.Database.DatabaseName
	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	fmt.Println("The database path is :", databasePath)
	if err != nil {
		fmt.Println("Faild to connect database: ", err)
	}

	db.AutoMigrate(&global.FileInfo{})

	fmt.Println("Database initialize success!")
	return db
}

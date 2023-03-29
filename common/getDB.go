package common

import (
	"database/sql"
	"fpt/global"

	"gorm.io/gorm"
)

func GetGromDB() *gorm.DB {
	return global.Ftpserver.DB
}

func GetSqlDB() (*sql.DB, error) {
	db, err := GetGromDB().DB()
	if err != nil {
		return nil, err
	}

	return db, nil
}

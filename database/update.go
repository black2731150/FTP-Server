package database

import (
	"fmt"
	"fpt/common"
	"fpt/global"
)

func UpdateText(name, text string) bool {

	Lock.Lock()
	db, err := common.GetSqlDB()
	if err != nil {
		fmt.Println("Faild to Get DB: ", err)
	}

	rs, err := db.Exec("update file_infos set text=? where name = ?;", text, name)
	if err != nil {
		fmt.Println("update faild:", err)
		return false
	}

	id, err := rs.LastInsertId()
	if err != nil {
		fmt.Println("rs is error: ", id, "  ", err)
		return false
	}
	Lock.Unlock()

	return true
}

func UpdateText1(name, text string) bool {
	DB := global.Ftpserver.DB
	Lock.Lock()
	DB.Table("file_infos").Where("name = ?", name).Update("text", text)
	Lock.Unlock()
	return true
}

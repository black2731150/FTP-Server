package database

import (
	"fmt"
	"fpt/common"
)

func UpdateText(name, text string) bool {

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

	return true
}

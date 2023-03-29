package database

import (
	"fmt"
	"fpt/common"
	"fpt/global"
)

func GetFileInfoFromDB(page, size int, search string) (int, int, []global.FileInfo) {
	//页数规范
	if size < 1 {
		size = 10
	}
	if page < 1 {
		page = 1
	}

	//获取数据库
	db, err := common.GetSqlDB()
	if err != nil {
		fmt.Println("Faild to connect database: ", err)
	}

	rows, err := db.Query("select * from file_infos where (name like ? or branch like ?) order by create_time desc limit ?,?;", "%"+search+"%", "%"+search+"%", (page-1)*size, size)
	if err != nil {
		fmt.Println("Faild to Get rows: ", err)
		panic("")
	}
	defer rows.Close()

	var fileinfos []global.FileInfo
	for rows.Next() {
		fileinfo := global.FileInfo{}
		err := rows.Scan(&fileinfo.Name, &fileinfo.Size, &fileinfo.CreateTime, &fileinfo.DownLandURL, &fileinfo.Md5, &fileinfo.Branch, &fileinfo.Text)
		if err != nil {
			fmt.Println("Faild to Get fileinfo: ", err)
			panic("")
		}
		fileinfos = append(fileinfos, fileinfo)
	}

	return getSize(search), getTotal(), fileinfos
}

func getTotal() int {
	var Total int
	common.GetGromDB().Raw("select count() from file_infos;").Scan(&Total)
	return Total
}

func getSize(search string) int {
	var Size int
	common.GetGromDB().Raw("select count() from file_infos where (name like ? or branch like ?)", "%"+search+"%", "%"+search+"%").Scan(&Size)
	return Size
}

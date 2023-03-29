package database

import (
	"fmt"
	"fpt/common"
	"fpt/global"
)

//把上传文件信息写进数据库里
func WriteToDB(filepath string, text string, branch string, url string) error {
	filestat, err := common.GetFileStat(filepath)
	if err != nil {
		fmt.Println("Faild to Get filInfo:", err)
		return err
	}

	fileinfo := new(global.FileInfo)
	fileinfo.Name = filestat.Name()
	fileinfo.Size = filestat.Size()
	fileinfo.CreateTime = filestat.ModTime().Format("2006-01-02 15:04:05")
	fileinfo.DownLandURL = url + "/files/" + filestat.Name()
	fileinfo.Md5 = common.HashMD5(filepath)
	fileinfo.Branch = branch
	fileinfo.Text = text

	global.Ftpserver.DB.Create(fileinfo)

	return nil
}

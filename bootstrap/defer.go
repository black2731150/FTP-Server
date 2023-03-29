package bootstrap

import "fpt/global"

//程序结束后需要做的事情，比如释放数据库连接
func DeferThings() {
	if global.Ftpserver.DB != nil {
		db, _ := global.Ftpserver.DB.DB()
		db.Close()
	}
}

package bootstrap

import "fpt/global"

func DeferThings() {
	if global.Ftpserver.DB != nil {
		db, _ := global.Ftpserver.DB.DB()
		db.Close()
	}
}

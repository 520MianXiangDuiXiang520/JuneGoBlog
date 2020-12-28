package main

import (
	"JuneGoBlog/src"
	juneDao "github.com/520MianXiangDuiXiang520/GinTools/dao"
	juneEmail "github.com/520MianXiangDuiXiang520/GinTools/email"
	utils "github.com/520MianXiangDuiXiang520/GinTools/log"
)

// 进行一些初始化操作
func doInit() {
	// 加载配置文件
	src.InitSetting("./setting.json")
	setting := src.GetSetting()
	// 初始化数据库连接
	err := juneDao.InitDBSetting(setting.MySQLSetting)
	if err != nil {
		utils.ExceptionLog(err, "Unable to connect to the database")
		panic("Unable to connect to the database")
	}
	if src.GetSetting().Others.Redis {
		// 初始化 Redis Pool
		err = juneDao.InitRedisPool(setting.RedisSetting)
		if err != nil {
			utils.ExceptionLog(err, "Unable to connect to the redis server")
			panic("Unable to connect to the redis server")
		}
	}
	// 初始化 SMTP
	juneEmail.InitSMTPDialer(setting.SMTPSetting.Host, setting.SMTPSetting.Username,
		setting.SMTPSetting.Password, setting.SMTPSetting.Port)
}

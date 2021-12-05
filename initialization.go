package main

import (
	"JuneGoBlog/src"
	juneDao "github.com/520MianXiangDuiXiang520/GinTools/dao"
	juneEmail "github.com/520MianXiangDuiXiang520/GinTools/email"
	utils "github.com/520MianXiangDuiXiang520/GinTools/log"
	"github.com/520MianXiangDuiXiang520/GoTools/dao"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

// 进行一些初始化操作
func doInit() {
	logFileName := "/blog/log/api.log"
	logF, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	logger := log.New(logF, "[gorm]", log.Ldate|log.Lshortfile)
	if err != nil {
		logF = os.Stdout
		log.Printf("can not open %s, err: %v", logFileName, err)
	}
	log.SetOutput(logF)
	gin.DefaultWriter = io.MultiWriter(logF)
	// 加载配置文件
	src.InitSetting("./setting.json")
	setting := src.GetSetting()
	// 初始化数据库连接
	err = juneDao.InitDBSetting(setting.MySQLSetting)
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
	dao.GetDB().SetLogger(logger)
	// 初始化 SMTP
	juneEmail.InitSMTPDialer(setting.SMTPSetting.Host, setting.SMTPSetting.Username,
		setting.SMTPSetting.Password, setting.SMTPSetting.Port)
}

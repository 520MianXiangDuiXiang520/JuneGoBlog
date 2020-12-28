package src

import (
	"github.com/520MianXiangDuiXiang520/GinTools/json"
	junePath "github.com/520MianXiangDuiXiang520/GinTools/path"
	"path"
	"runtime"
	"sync"
	"time"
)

type Setting struct {
	Others         Others    `json:"others"`
	MySQLSetting   MySQLConn `json:"mysql_setting"`
	RedisSetting   RedisConn `json:"redis_setting"`
	SMTPSetting    SMTPConn  `json:"smtp_setting"`
	CorsAccessList []string  `json:"cors_access_list"`
}

type Others struct {
	AbstractLen int    `json:"abstractLen"` // 默认的摘要长度
	Redis       bool   `json:"redis"`       // 是否使用redis
	MyEmail     string `json:"my_email"`
	DetailLink  string `json:"detail_link"`
	SiteLink    string `json:"site_link"`
}

// mysql 配置
type MySQLConn struct {
	Engine    string        `json:"engine"`
	DBName    string        `json:"db_name"`
	User      string        `json:"user"`
	Password  string        `json:"password"`
	Host      string        `json:"host"`
	Port      int           `json:"port"`
	MIdleConn int           `json:"max_idle_conn"` // 最大空闲连接数
	MOpenConn int           `json:"max_open_conn"` // 最大打开连接数
	MLifetime time.Duration `json:"max_lifetime"`  // 连接超时时间
	LogMode   bool          `json:"log_mode"`
}

// redis 配置
type RedisConn struct {
	Host      string        `json:"host"`
	Password  string        `json:"password"`
	Port      int           `json:"port"`
	MIdleConn int           `json:"max_idle_conn"` // 最大空闲连接数
	MOpenConn int           `json:"max_open_conn"` // 最大打开连接数
	MLifetime time.Duration `json:"max_lifetime"`  // 连接超时时间
}

// email 配置
type SMTPConn struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

var setting *Setting
var settingLock sync.Mutex

func InitSetting(filePath string) {
	defer func() {
		if e := recover(); e != nil {
			settingLock.Unlock()
		}
	}()
	filename := filePath
	if !junePath.IsAbs(filePath) {
		_, currently, _, _ := runtime.Caller(1)
		filename = path.Join(path.Dir(currently), filePath)
	}
	if setting == nil {
		settingLock.Lock()
		if setting == nil {
			json.FromFileLoadToObj(&setting, filename)
		}
		settingLock.Unlock()
	}
}

func GetSetting() *Setting {
	if setting == nil {
		panic("setting Uninitialized！")
	}
	return setting
}

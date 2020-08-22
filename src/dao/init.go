package dao

import (
	"JuneGoBlog/src/util"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"strings"
)

var DB *gorm.DB
var RedisPool *redis.Pool

type mysqlStruct struct {
	DbName   string `ini:"dbname"`
	User     string `ini:"user"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Ip       string `ini:"ip"`
}

type redisConn struct {
	Host     string `ini:"host"`
	Password string `ini:"password"`
	Port     string `ini:"port"`
}

func init() {
	var err error
	err = InitDB()
	err = InitRedis()
	if err != nil {
		return
	}
}

func InitRedis() error {
	rc := new(redisConn)
	util.Load("../../secret.ini", "redis", rc)
	var err error

	RedisPool = &redis.Pool{ //实例化一个连接池
		MaxIdle:     16,  //最初的连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp",
				rc.Host+":"+rc.Port,
				redis.DialPassword(rc.Password),
			)
		},
	}

	log.Println("============== Redis Connect Success ==============")
	return err
}

func InitGoRedis() {

}

func InitDB() error {
	var err error
	var ms mysqlStruct
	util.Load("../../secret.ini", "mysql", &ms)
	s := []string{ms.User, ":", ms.Password, "@tcp(", ms.Ip, ":", strconv.Itoa(ms.Port), ")/",
		ms.DbName, "?charset=utf8&parseTime=True&loc=Local"}
	connStr := strings.Join(s, "")
	DB, err = gorm.Open("mysql", connStr)
	if err != nil {
		log.Println("Open DB Error!!!")
		return err
	}
	log.Println("============== MySQL Connect Success ==============")
	return nil
}

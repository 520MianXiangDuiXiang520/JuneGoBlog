package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"strings"
)

var DB *gorm.DB

type mysqlStruct struct {
	DbName   string `ini:"dbname"`
	User     string `ini:"user"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Ip       string `ini:"ip"`
}

func init() {
	err := InitDB()
	if err != nil {
		return
	}
}

func InitDB() error {
	var err error
	var ms mysqlStruct
	Load("../../secret.ini", "mysql", &ms)
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

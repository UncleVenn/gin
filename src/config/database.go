package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var Db *sql.DB
var DbError error

const (
	USERNAME = "root"
	PASSWORD = "root"
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DATABASE = "golangtest"
	CHARSET  = "UTF8"
)

func init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USERNAME, PASSWORD, HOST, PORT, DATABASE, CHARSET)
	// 打开连接失败
	Db, DbError = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if DbError != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + DbError.Error())
	}
	// 最大连接数
	Db.SetMaxOpenConns(100)
	// 闲置连接数
	Db.SetMaxIdleConns(20)
	// 最大连接周期
	Db.SetConnMaxLifetime(100 * time.Second)

	if DbError = Db.Ping(); nil != DbError {
		panic("数据库链接失败: " + DbError.Error())
	}
}

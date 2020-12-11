package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Mydb struct {
	Dsn string
	Con *gorm.DB
}

func (mydb *Mydb) SetDbDsn() {
	mydb.Dsn = "root:@tcp(127.0.0.1:3306)/test_tamu_com?charset=utf8mb4&parseTime=True&loc=Local"
}

func (mydb *Mydb) CreateDb() {
	mydb.SetDbDsn()
	db, _ := gorm.Open(mysql.Open(mydb.Dsn), &gorm.Config{})
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	mydb.Con = db

}

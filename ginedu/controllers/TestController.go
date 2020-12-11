package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type IdolSrTable struct {
	ID      int `gorm:"primarykey"`
	Table   string
	IsClear int8
}

type Game struct {
	ID          int `gorm:"primarykey"`
	Pid         int
	Name        string
	Icon        string
	MinPrice    int
	Unit        string
	Desc        string
	Sort        int8
	Status      int8
	IsRecommend int8
	CreateAt    int64
	UpdateAt    int64
	DeleteAt    int64
}

func (v Game) TableName() string {

	return "idol_game"
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "test",
	})
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "index",
	})

}

func GoRsync(c *gin.Context) {
	cCp := c.Copy()
	go func() {

		time.Sleep(5 * time.Second)

		log.Println("Done! in path " + cCp.Request.URL.Path)
	}()

	c.JSON(http.StatusOK, gin.H{
		"title": "goroutine",
	})

}

func DbList(c *gin.Context) {

	dsn := "root:@tcp(127.0.0.1:3306)/test_tamu_com?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	test := IdolSrTable{}

	db.First(&test)
	log.Printf("result:%+v", test)

}

func DbGames(c *gin.Context) {

	println(c.Param)
	dsn := "root:@tcp(127.0.0.1:3306)/test_tamu_com?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	test := Game{}

	db.Last(&test)

	type Re struct {
		CreateAt string
	}

	re := Re{}

	re.CreateAt = time.Unix(test.CreateAt, 0).Format("2006-01-02 15:04:05")
	c.JSON(http.StatusOK, re)

}

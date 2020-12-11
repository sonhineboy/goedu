package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"reflect"
	"suiyidian.cn/sonhineboy/gintwo/core/database"
	"suiyidian.cn/sonhineboy/gintwo/models"
	"time"
)

var (
	g     errgroup.Group
	myDb  database.Mydb
	dbCon *gorm.DB
)

type GameParam struct {
	Status int    `form:"status" json:"status"`
	Name   string `form:"name" binding:"required,CustomValidationErrors" json:"name" label:"名字"`
}

func main() {
	r := gin.Default()
	r.Use(gin.Recovery())
	myDb.CreateDb()
	dbCon = myDb.Con

	// 读取env
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("APP_NAME"))

	//自定义验证消息
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")

	//自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		zhErr := zhtranslations.RegisterDefaultTranslations(v, trans)

		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")

			log.Println(name)

			if name == "-" {
				return ""
			}
			return name
		})

		if zhErr != nil {
			fmt.Println(zhErr)
		}

		err := v.RegisterValidation("CustomValidationErrors", CustomValidationErrors)
		if err != nil {
			fmt.Println("error")
		}
	}

	r.GET("/", func(context *gin.Context) {
		var game []models.Game
		dbCon.Find(&game)

		context.XML(http.StatusOK, gin.H{
			"test": "my name is kevin",
			"you":  game,
		})
	})

	r.GET("/game/lists", func(context *gin.Context) {
		var request GameParam

		if err := context.ShouldBind(&request); err == nil {
			context.JSON(http.StatusOK, request)
		} else {
			errs, _ := err.(validator.ValidationErrors)
			context.JSON(http.StatusOK, gin.H{
				"errors": errs.Translate(trans),
			})
		}

	})

	server02 := &http.Server{
		Addr:         ":9999",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}

func CustomValidationErrors(fl validator.FieldLevel) bool {
	return fl.Field().String() != "admin"
}

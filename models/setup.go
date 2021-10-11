package models

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"restapiGin/environment"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Client *redis.Client

func init() {
	//Initializing redis
	Client = redis.NewClient(&redis.Options{
		Addr: environment.ViperEnvVariable("REDIS_DSN"), //redis port
	})
	_, err := Client.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis connected")
}

func SetupDB() *gorm.DB {
	USER := environment.ViperEnvVariable("DB_USER")
	PASS := environment.ViperEnvVariable("DB_PASS")
	HOST := environment.ViperEnvVariable("DB_HOST")
	PORT := environment.ViperEnvVariable("DB_PORT")
	DBNAME := environment.ViperEnvVariable("DB_NAME")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}

	return db
}



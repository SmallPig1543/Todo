package dao

import (
	"Todo/conf"
	"Todo/repository/db/model"
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitMySQL() {

	dsn := conf.MySqlUser + ":" + conf.MySqlPassword + "@tcp(" + conf.MysqlIP + ":" + conf.MysqlPort + ")/" + conf.MySqlDataBase + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(model.User{})
	_ = db.AutoMigrate(model.Task{})
	DB = db
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}

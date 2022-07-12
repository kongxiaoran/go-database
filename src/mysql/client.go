package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	id   int
	name string
	age  int
}

func getMysqlClient() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("连接数据库失败", err)
	}

	person := new(Person)
	db.First(&person)
	if person != nil {
		fmt.Println(person)
	}

}

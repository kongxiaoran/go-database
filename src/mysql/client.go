package mysql

import (
	"encoding/json"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	ID   int64
	Name string
	Age  int64
}

type MysqlConfig struct {
	Ip       string
	User     string
	Password string
}

func GetMysqlClient() {

	dsn := "root:298444kxrtxy@tcp(150.158.165.30:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("连接数据库失败", err)
	}

	var person Person
	db.Raw("select id,name,age from person where id = ?", 1).Scan(&person)
	fmt.Println(person)

	// db.Rows("select id,name,age from person")

}

func GetGORMClient() {

	// dsn := "root:298444kxrtxy@tcp(150.158.165.30:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	fmt.Println("连接数据库失败", err)
	// }

	// var person Person
	// db.Raw("select id,name,age from person where id = ?", 1).Scan(&person)
	// fmt.Println(person)
	file, _ := os.Open("init.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	conf := MysqlConfig{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Ip)
}

/*
	GORM 官方文档：
	https://gorm.io/zh_CN/docs/connecting_to_the_database.html


*/

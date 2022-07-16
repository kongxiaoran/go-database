package mysql

import (
	"fmt"

	util "go-database/src/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	ID   int64
	Name string
	Age  int64
}

type MysqlConfig struct {
	Work     MysqlWork
	Ip       string
	User     string
	Password string
	Base     string
}

type MysqlWork struct {
	PrepareSQL     string
	SQL            string
	PrimaryKey     string
	StartTimeStamp string
	EndTimeStamp   string
	ThreadNumber   int
	Model          string
}

var Config *MysqlConfig
var Work *MysqlWork
var DB *gorm.DB
var err error
var dsn string

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

// 加载配置文件，初始化 mysql客户端
func init() {
	Config = util.GetObjectFromJSON("./mysql/init.json", new(MysqlConfig)).(*MysqlConfig)
	Work = &Config.Work

	dsn = Config.User + ":" + Config.Password + "@tcp(" + Config.Ip + ")/" + Config.Base + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败", err)
	}
}

func GetGORMClient() {

	var person Person
	DB.Raw("select id,name,age from person where id = ?", 1).Scan(&person)
	fmt.Println(person)
}

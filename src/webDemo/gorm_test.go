package webDemo

import (
	"fmt"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db, _ = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:000000@tcp(127.0.0.1:3306)/small_web?charset=utf8mb4&parseTime=True&loc=Local", // 账号密码地址端口
	}), &gorm.Config{})
}

type TestTable struct {
	Id   int
	Name string
	Code string
	Age  int
}

func TestGormConnect(t *testing.T) {
	CreateTable(t)
	// CreateOneRow()
	SelectOneRow()
}

func CreateTable(t *testing.T) {
	fmt.Printf("db.Migrator().HasTable(\"test_table\"): %v\n", db.Migrator().HasTable("test_table"))
	// db.Migrator().DropTable("test_table")
	if ok := db.Migrator().HasTable("test_table"); !ok {
		// t.Logf("ok: %v\n", ok)
		t.Log("还未建立表格..")
		t.Log("建立表中..")
		db.Table("test_table").AutoMigrate(&TestTable{})
		t.Log("新的表格建立完成")
	} else {
		t.Log("已经建立表格..")
	}
}

func CreateOneRow() {
	db.Table("test_table").Create(&TestTable{Name: "lcy", Code: "200", Age: 20})
}

func SelectOneRow() {
	var list []TestTable
	db.Table("test_table").Find(&list)
	for i, val := range list {
		fmt.Printf("i: %v\n", i)
		fmt.Printf("val: %v\n", val)
	}
}

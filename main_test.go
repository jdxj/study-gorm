package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsnTemp = "%s:%s@tcp(%s:3306)/test_study_gorm?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	GormDB *gorm.DB
	SqlDB  *sql.DB
)

func TestMain(t *testing.M) {
	dsn := fmt.Sprintf(dsnTemp, user, pass, host)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	GormDB = db
	SqlDB = sqlDB
	os.Exit(t.Run())
}

func TestCreate(t *testing.T) {
	user := &User{
		Name:   "jdxj",
		Gender: 1,
		Phone:  "abc",
	}
	err := GormDB.Create(user).Error
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestTimeRounding(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format(time.RFC3339Nano))
	fmt.Println(now.Round(time.Second).Format(time.RFC3339Nano))
	// output:
	//     2022-04-24T13:49:34.891059+08:00
	//     2022-04-24T13:49:35+08:00
}

func TestFind(t *testing.T) {
	user := &User{}
	db := GormDB.Where("id = 1")

	var c int64
	err := db.Count(&c).Error
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	err = db.Find(user).Error
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	fmt.Printf("%+v\n", user)
}

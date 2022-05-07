package gorm

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/jdxj/study-gorm/er"
	"github.com/jdxj/study-gorm/key"
)

var (
	GormDB *gorm.DB
	SqlDB  *sql.DB
)

func TestMain(t *testing.M) {
	db, err := gorm.Open(mysql.Open(key.DSN()))
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
	user := &er.User{
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

func TestTake(t *testing.T) {
	user := er.User{}
	err := GormDB.WithContext(context.Background()).
		Where("id = ?", 3).
		Take(&user).
		Error
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", user)
}

func TestFind(t *testing.T) {
	user := &er.User{}
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

func TestMySQL_Dialector_quoteTo(t *testing.T) {
	d := mysql.Dialector{}

	buffer := bytes.NewBuffer(nil)
	d.QuoteTo(buffer, "a.b")
	//d.QuoteTo(buffer, "`a.b")
	//d.QuoteTo(buffer, "a")
	//d.QuoteTo(buffer, "a`b")

	fmt.Println(buffer.String())
}

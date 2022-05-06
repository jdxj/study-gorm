package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"github.com/jdxj/study-gorm/er"
	"github.com/jdxj/study-gorm/key"
)

var (
	db *sql.DB
)

func TestMain(t *testing.M) {
	var err error
	db, err = sql.Open("mysql", key.DSN())
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	os.Exit(t.Run())
}

func TestSelect(t *testing.T) {
	row := db.QueryRowContext(context.Background(), "select * from users where id = ?", 1)

	user := &er.User{}
	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt, &user.Name, &user.Gender, &user.Phone)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", user)
}

func TestInsert(t *testing.T) {
	user := &er.User{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Name:   "123",
		Gender: 4,
		Phone:  "567",
	}

	_, err := db.ExecContext(
		context.Background(),
		"insert into users (created_at,updated_at,name,gender,phone) values (?,?,?,?,?)",
		user.CreatedAt, user.UpdatedAt, user.Name, user.Gender, user.Phone)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

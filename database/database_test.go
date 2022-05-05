package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"

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

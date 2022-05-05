package gorm

import (
	"fmt"
	"testing"
)

func TestSQLDB(t *testing.T) {

	rows, err := SqlDB.Query("select * from users")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			t.Fatalf("%s\n", err)
		}

		for i, c := range columns {
			fmt.Printf("%d: %s; ", i, c)
		}
		fmt.Println()
		rows.Scan()
	}
}

func TestSendNilChan(t *testing.T) {
	var c chan int
	c <- 1
}

package key

import "fmt"

const (
	dsnTemp = "%s:%s@tcp(%s:3306)/test_study_gorm?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	User string
	Pass string
	Host string
)

func DSN() string {
	return fmt.Sprintf(dsnTemp, User, Pass, Host)
}

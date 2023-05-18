package database

import (
	"fmt"

	"github.com/dev-hyunsang/daily-todo/config"
	"github.com/dev-hyunsang/daily-todo/ent"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() (*ent.Client, error) {
	client, err := ent.Open("mysql",
		fmt.Sprintf(
			"root:%s@tcp(%s)/%s?parseTime=True", config.GetEnv("MYSQL_PASSWORD"), config.GetEnv("MYSQL_HOST"), config.GetEnv("MYSQL_DATABASE")),
	)

	return client, err
}

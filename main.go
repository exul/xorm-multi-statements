package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func main() {
	var err error
	engine, err := xorm.NewEngine("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		panic(err)
	}

	session := engine.NewSession()
	defer session.Close()
	db := session.DB()

	if _, err := db.Exec("DROP TABLE if exists `test`"); err != nil {
		panic(err)
	}

	if _, err := db.Exec("CREATE TABLE `test` (`id` int(11) NOT NULL, `value` int(11) NOT NULL) "); err != nil {
		panic(err)
	}

	if _, err := db.Exec("INSERT INTO test VALUES (1, 1)"); err != nil {
		panic(err)
	}

	interpolatedQuery := `
		set @id = 1;
	 	select @id
	`

	ctx := context.TODO()
	_, err = db.QueryContext(ctx, interpolatedQuery)
	if err != nil {
		panic(err)
	}

	fmt.Println("Yay")
}

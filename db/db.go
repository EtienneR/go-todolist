package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
	"log"
)

type Messages struct {
	Id    int64  `db:"id"`
	Title string `db:"title"`
}

func InitDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:@/todolist")
	//db, err := sql.Open("mysql", "root:@tcp(192.168.59.103:3306)/todolist")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	dbmap.AddTableWithName(Messages{}, "messages").SetKeys(true, "Id")
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create table failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

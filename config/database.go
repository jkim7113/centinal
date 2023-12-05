package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jkim7113/centinal/util"
)

func CreateConnection() *sql.DB {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	pw := os.Getenv("DB_PW")
	name := os.Getenv("DB_NAME")

	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pw, host, port, name)

	db, err := sql.Open("mysql", DSN)
	util.PanicIfError(err)

	err = db.Ping()
	util.PanicIfError(err)

	fmt.Println("DB Connection OK")
	return db
}

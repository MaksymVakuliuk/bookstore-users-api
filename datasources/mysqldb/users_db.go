package mysqldb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mySql_Username = "MYSQL_USERNAME"
	mySql_Password = "MYSQL_PASSWORD"
	mySql_Address  = "MYSQL_ADDRESS"
	mySql_DBName   = "MYSQL_DB_NAME"
)

var (
	UsersDB *sql.DB

	mySqlUsername = os.Getenv(mySql_Username)
	mySqlPassword = os.Getenv(mySql_Password)
	mySqlAddress  = os.Getenv(mySql_Address)
	mySqlDBName   = os.Getenv(mySql_DBName)
)

func init() {
	datasourseName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", mySqlUsername, mySqlPassword, mySqlAddress, mySqlDBName)
	var err error
	UsersDB, err = sql.Open("mysql", datasourseName)
	if err != nil {
		panic(err)
	}
	if err = UsersDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}

package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/AnggaArdhinata/k-stylehub/src/configs"
)

var db *sql.DB

var err error

func Init()  {
	conf := configs.GetConfig()

	dbConnect := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", dbConnect)
	if err != nil {
		panic("Connection DB Error!")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}

}

func CreateCon() *sql.DB  {
	return db
}

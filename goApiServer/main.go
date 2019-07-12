package main

import (
	"database/sql"
	"fmt"
	_ "myGoWeb/goApiServer/routers"
	"github.com/astaxie/beego"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// init to connect DB
	db, _ := connectDB()
	_, err := db.Exec("insert into users (username, password) values(?, ?)", "will", "zte")
	if err != nil {
		fmt.Println("insert into database failed with errMsg ==> " + err.Error())
	}

	// init provider

	// init model


	startBeegoWebServer()
}

func connectDB() (*sql.DB, error) {
	var db *sql.DB
	var err error
	connectString := "willjiang:willjiang@tcp(localhost:3306)/angular?charset=utf8"
	for {
		db, err = sql.Open("mysql", connectString)
		if err != nil {
			fmt.Println("Connect DB mysql failed , trying after 5 second. errMsg ==> " + err.Error())
			time.Sleep( 5 * time.Second)
			continue
		}
		break
	}
	return db, nil
}

func startBeegoWebServer() {
	beego.Run()
}


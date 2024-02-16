package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	id   int
	name string
}

func main() {

	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:33060)/%v", DbUser, DbPasswprd, DBName)

	db, err := sql.Open("mysql", connectionString)
	checkError(err)
	defer db.Close()

	rows, err := db.Query("SELECT * from data")
	checkError(err)

	for rows.Next() {
		var data Data
		err := rows.Scan(&data.id, &data.name)
		checkError(err)
		fmt.Println(data)
	}

}

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}

}

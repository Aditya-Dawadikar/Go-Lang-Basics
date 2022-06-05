package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/temp_db")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("db is connected")
		defer db.Close()
		// make sure connection is available
		err = db.Ping()
		fmt.Println(err)
		if err != nil {
			fmt.Println("MySQL db is not connected")
			fmt.Println(err.Error())
		}
	}
	return db
}

type Person struct {
	name string
	id   string
}

func readData(db *sql.DB) {
	result, err := db.Query("select * from people_record")
	defer result.Close()

	if err != nil {
		log.Fatal(err)
	}

	for result.Next() {
		var p Person
		err := result.Scan(&p.name, &p.id)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(p)
	}
}

func writeData(db *sql.DB) {
	sql := "Insert into people_record values('sherloc','105')"
	res, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(res.RowsAffected())
}

func main() {
	db := CreateCon()
	defer db.Close()
	readData(db)
	writeData(db)
}

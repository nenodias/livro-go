package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3308)/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("update usuarios set nome = ? where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec("Xubiru", 1)
	stmt.Exec("Xrabas", 2)

	stmt2, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	stmt2.Exec(3)
}

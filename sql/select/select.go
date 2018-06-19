package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3308)/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select id, nome from usuarios where id >= ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var u Usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}

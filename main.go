package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "nakama"
)

type character struct {
	name        string
	age         int
	devilFruit  string
	affiliation string
}

func main() {

	//connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//open db
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	//close db
	defer db.Close()

	//check db
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	luffy := character{name: "Monkey D. Luffy", devilFruit: "Gomu Gomu no Mi", affiliation: "Straw Hat Pirates"}
	zoro := character{name: "Roronoa Zoro", age: 17, affiliation: "Straw Hat Pirates"}

	//insert
	query := `insert into "characters" ("name", "affiliation") values ($1, $2), ($3, $4)`
	_, e := db.Exec(query, zoro.name, zoro.affiliation, luffy.name, luffy.affiliation)
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println("Connected to Nakama!")

}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/jeremyholbrecht/nakama/internal/routes"
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

type devilFruit struct {
	name     string
	category string
	owner    string
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

	gomuGomuNoMi := devilFruit{name: "Gomu Gomu no Mi", category: "Paramecia", owner: luffy.name}
	//insert
	insertInToCharacters := `insert into "characters" ("name", "affiliation") values ($1, $2), ($3, $4)`
	_, insertInToCharactersErr := db.Exec(insertInToCharacters, zoro.name, zoro.affiliation, luffy.name, luffy.affiliation)
	if insertInToCharactersErr != nil {
		log.Fatal(insertInToCharactersErr)
	}

	insertInToDevilFruits := `insert into "devilfruits" ("name", "category", "owner") values ($1,$2,$3)`
	_, insertInToDevilFruitsErr := db.Exec(insertInToDevilFruits, gomuGomuNoMi.name, gomuGomuNoMi.category, gomuGomuNoMi.owner)

	if insertInToDevilFruitsErr != nil {
		log.Fatal(insertInToDevilFruitsErr)
	}
	fmt.Println("Connected to Nakama!")

	router := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf("%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	err2 := http.ListenAndServe(addr, router)
	if err2 != nil {
		panic(err2)
	}

}

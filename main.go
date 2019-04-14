package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var tpl *template.Template
var db *sql.DB

func init() {
	tpl = template.Must(template.ParseGlob("./templates/*.html"))
	dbInit()
}

func dbInit() {
	user := os.Getenv("DB_USER")
	pw := os.Getenv("DB_PASSWORD")
	connection := fmt.Sprintf("%s:%s@unix(/var/run/mysqld/mysqld.sock)/practiceDB", user, pw)
	var err error
	db, err = sql.Open("mysql", connection)
	if err != nil {
		log.Fatal("Database init fatal: ", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Database pinging fatal: ", err)
	}
}

func main() {
	dbInit()
	defer db.Close()
	log.Println(db)
	log.Println("Website Launch on http://localhost:8080")
	http.Handle("/", http.FileServer(http.Dir("www")))
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/join", join)
	http.ListenAndServe(":8080", nil)
}

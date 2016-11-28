package main

import (
	"database/sql"
	"log"

	"github.com/kataras/iris"
	_ "github.com/mattn/go-sqlite3"
)

const DBpath = "./db/glip.db"

func initDB(c *iris.Context) {
	db := openDB(DBpath)
	defer db.Close()

	createTables(db)

	c.Next()
}

func openDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Panic(err)
	}
	return db
}

func createTables(db *sql.DB) {
	tables := `
	CREATE TABLE IF NOT EXISTS glip(
		Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		Url TEXT NOT NULL,
		Msg TEXT NOT NULL,
		Idate DATETIME NOT NULL,
		Ip TEXT NOT NULL
	);
	`
	_, err := db.Exec(tables)
	if err != nil {
		log.Fatal(err)
	}
}

func checkURL(db *sql.DB, url string) ([]string, error) {
	infos := make([]string, 5)
	query := `
	SELECT * FROM glip WHERE Url = ?;
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(url).Scan(&infos[0], &infos[1], &infos[2], &infos[3], &infos[4])
	return infos, err
}

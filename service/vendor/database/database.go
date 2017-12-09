package database

import (
	"database/sql"
	"log"
)

const createUserTableSQL = `CREATE TABLE IF NOT EXISTS User (
	uid INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	password TEXT NOT NULL,
	phone TEXT NOT NULL,
	email TEXT NOT NULL,
	CONSTRAINT uu UNIQUE (username)
);`

const createLoginTableSQL = `CREATE TABLE IF NOT EXISTS Login (
	token TEXT NOT NULL PRIMARY KEY,
	username TEXT NOT NULL,
	FOREIGN KEY(username) REFERENCES User(username),
	CONSTRAINT uu UNIQUE (username)
);`

var theDB *sql.DB

func prepareDB(dbfile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbfile)
	// open database fail
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return nil, err
	}
	_, err = db.Exec(createUserTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = db.Exec(createLoginTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("database initialized")
	return db, nil
}

// InitializeDB with specific file
func InitializeDB(dbfile string) error {
	if theDB != nil {
		theDB.Close()
	}
	db, err := prepareDB(dbfile)
	if err != nil {
		return err
	}
	theDB = db
	return nil
}

func pQuery(db *sql.DB, s string, args ...interface{}) (chan *sql.Rows, error) {
	stmt, err := db.Prepare(s)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	chanRow := make(chan *sql.Rows)
	go func() {
		for rows.Next() {
			chanRow <- rows
		}
		stmt.Close()
	}()
	return chanRow, nil

}
func pExec(db *sql.DB, sql string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(args...)
	stmt.Close()
	return result, err
}

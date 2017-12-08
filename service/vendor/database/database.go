package database

import (
	"args"
	"database/sql"
	"entity"
	"log"
)

const createTableSQL = `CREATE TABLE IF NOT EXISTS User (
	id INTEGER NOT NULL PRIMARY KEY,
	username TEXT NOT NULL,
	password TEXT NOT NULL,
	phone TEXT NOT NULL,
	email TEXT NOT NULL
);`

type agendaDB struct {
	db *sql.DB
}

var theDB *agendaDB

func prepareDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", *args.DB)
	// open database fail
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		return nil, err
	}
	_, err = db.Exec(createTableSQL)
	// create table fail
	if err != nil {
		return nil, err
	}
	log.Print("database initialized")
	return db, nil
}

func getUserDAO() (entity.UserDAO, error) {
	if theDB == nil {
		db, err := prepareDB()
		if err != nil {
			return nil, err
		}
		theDB = &agendaDB{db}
	}
	return theDB, nil
}

func (d *agendaDB) GetAllUsers() []*entity.User {
	return nil
}

func (d *agendaDB) StoreUsers(user *entity.User) {

}

func (d *agendaDB) GetUser(username string) *entity.User {
	return nil
}

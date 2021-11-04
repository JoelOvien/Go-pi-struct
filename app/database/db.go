package database

import (
	"log"

	"github.com/jmoiron/sqlx"
)

//Our interface called PostDB, the interface behaves like a contract(methods to implement)
//we respect this if we want our DB to be a PostDB database
type PostDB interface{
	Open() error
	Close() error
}

//create a DB struct for our app with a single field db of type sqlxDB which is a superset of database/sql from the standard library

type DB struct{
	db *sqlx.DB
}

//Function to open our db
func(d *DB) Open()error{
	//Open a connection to our database
	pg, err := sqlx.Open("postgres",pgConnStr)

	if err != nil{
		return err
	}
	log.Println("Connected to Database!")
	pg.MustExec(createSchema)
	d.db = pg

	return nil
}

//Function to close our db
func(d *DB) Close()error{
	return d.db.Close()
}
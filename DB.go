package main

import (
	"log"
	"os"
)

type DB struct {
	data map[string]map[string]string
}

func initDB() DB {
	db := DB{make(map[string]map[string]string)}

	mkerr := os.Mkdir("./data", os.ModePerm)

	err := Load("./data", &db.data)
	if err != nil && mkerr != nil {
		log.Printf("Load Data Error: %v\n", err)
	}

	return db
}

func (db *DB) add(collection string, key string, value string) {
	if len(db.data[collection]) == 0 {
		db.data[collection] = make(map[string]string)
	}

	db.data[collection][key] = value

	err := Save("./data", collection, db.data[collection])
	if err != nil {
		log.Printf("Save Data Error: %v\n", err)
	}
}

func (db *DB) get(collection string, key string) string {
	return db.data[collection][key]
}

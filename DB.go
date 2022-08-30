package main

import (
	"log"
	"os"
)

type DB struct {
	data map[string]string
}

func initDB() DB {
	db := DB{make(map[string]string)}

	mkerr := os.Mkdir("./data", os.ModePerm)

	err := Load(&db.data)
	if err != nil && mkerr != nil {
		log.Printf("Load Data Error: %v\n", err)
	}

	return db
}

func (db *DB) add(key string, value string) {
	db.data[key] = value

	err := Save(db.data)
	if err != nil {
		log.Printf("Save Data Error: %v\n", err)
	}
}

func (db *DB) get(key string) string {
	return db.data[key]
}

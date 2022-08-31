package main

import (
	"fmt"
	"log"
	"os"
)

type Database struct {
	name        string
	data        map[string]map[string]string
	collections []string
}
type DB struct {
	databases []Database
	names     map[string]int
}

func initDB() DB {
	db := DB{
		[]Database{},
		make(map[string]int),
	}

	mkerr := os.Mkdir("./data", os.ModePerm)

	sub, err := os.ReadDir("./data")
	if err != nil {
		log.Printf("Load Data Error: %v\n", err)
	} else {
		for _, ifile := range sub {
			if ifile.IsDir() {
				db.databases = append(db.databases, Database{
					ifile.Name(),
					make(map[string]map[string]string),
					[]string{},
				})
				db.names[ifile.Name()] = len(db.names)

				err := Load("./data", &db.databases[len(db.databases)-1])
				if err != nil && mkerr != nil {
					log.Printf("Load Data Error: %v\n", err)
				}
			}
		}
	}

	return db
}

func (db *Database) add(collection string, key string, value string) {
	if len(db.data[collection]) == 0 {
		db.data[collection] = make(map[string]string)
	}

	db.data[collection][key] = value

	err := Save("./data", db, collection)
	if err != nil {
		log.Printf("Save Data Error: %v\n", err)
	}
}

func (db *Database) get(collection string, key string) string {
	return db.data[collection][key]
}

func (db *DB) createDB(database string, collection string) {
	newdb := Database{
		database,
		make(map[string]map[string]string),
		[]string{collection},
	}
	newdb.data[collection] = make(map[string]string)

	db.databases = append(db.databases, newdb)
	db.names[database] = len(db.databases) - 1

	err := os.Mkdir(fmt.Sprintf("./data/%s", database), os.ModePerm)
	if err != nil {
		log.Printf("Save Mkdir Error: %s\n", err)
	}
	err = Save("./data", &db.databases[db.names[database]], collection)
	if err != nil {
		log.Printf("Save Error: %s\n", err)
	}
}

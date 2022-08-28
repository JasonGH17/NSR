package main

type DB struct {
	data map[string] string
}

func (db *DB) add(key string, value string) {
	db.data[key] = value
}

func (db *DB) get(key string) string {
	return db.data[key]
}
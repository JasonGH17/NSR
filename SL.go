package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"sync"
)

var lock sync.Mutex

func Save(path string, database *Database, collection string) error {
	lock.Lock()
	defer lock.Unlock()

	file, err := os.Create(fmt.Sprintf("%s/%s/%s.bin", path, database.name, collection))
	if err != nil {
		return err
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	err = enc.Encode(database.data[collection])
	if err != nil {
		return err
	}

	return nil
}

func Load(path string, database *Database) error {
	lock.Lock()
	defer lock.Unlock()

	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, ifile := range files {
		if !ifile.IsDir() {
			ifile.Info()
			name := ifile.Name()[:len(ifile.Name())-4]

			fmt.Printf("Loading Collection: %s/%s/%s.bin\n", path, database.name, name)
			file, err := os.Open(fmt.Sprintf("%s/%s/%s.bin", path, database.name, name))
			if err != nil {
				return err
			}
			defer file.Close()

			buff := make(map[string]string)

			dec := gob.NewDecoder(file)
			err = dec.Decode(&buff)
			if err != nil {
				return err
			}

			(*database).data[name] = buff
			(*database).collections = append((*database).collections, name)
		} else {
			continue
		}
	}

	return nil
}

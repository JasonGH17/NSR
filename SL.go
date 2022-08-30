package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"sync"
)

var lock sync.Mutex

func Save(path string, collection string, data map[string]string) error {
	lock.Lock()
	defer lock.Unlock()

	file, err := os.Create(fmt.Sprintf("%s/%s.bin", path, collection))
	if err != nil {
		return err
	}
	defer file.Close()

	enc := gob.NewEncoder(file)
	err = enc.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func Load(path string, data *map[string]map[string]string) error {
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

			fmt.Printf("Loading Collection: %s/%s.bin\n", path, name)
			file, err := os.Open(fmt.Sprintf("%s/%s.bin", path, name))
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

			(*data)[name] = buff
		} else {
			continue
		}
	}

	return nil
}

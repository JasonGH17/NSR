package main

import (
	"encoding/gob"
	"os"
	"sync"
)

var lock sync.Mutex

func Save(data map[string]string) error {
	lock.Lock()
	defer lock.Unlock()

	file, err := os.Create("./data/data.bin")
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

func Load(data *map[string]string) error {
	lock.Lock()
	defer lock.Unlock()

	file, err := os.Open("./data/data.bin")
	if err != nil {
		return err
	}
	defer file.Close()
	
	dec := gob.NewDecoder(file)
	return dec.Decode(data)
}

/*
func unmarshal(reader io.Reader, data *interface{}) error {
	return json.NewDecoder(reader).Decode(data)
}

func Save(path string, data interface{}) error {
	lock.Lock()
	defer lock.Unlock()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader, err := marshal(data)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, reader)
	return err
}

func Load(path string, data *interface{}) error {
	lock.Lock()
	defer lock.Unlock()

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return unmarshal(file, data)
}
*/

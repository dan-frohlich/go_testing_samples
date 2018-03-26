package main

import (
	"github.com/tidwall/buntdb"
)

var db *buntdb.DB

func initializeDB() error {
	var err error
	db, err = buntdb.Open("data.db")
	return err
}

func closeDB() {
	db.Close()
}

func put(key, value string) error {
	return db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})
}

func get(key string) (string, error) {
	var cv string
	err := db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil {
			return err
		}
		cv = val
		return nil
	})
	return cv, err
}

func list() ([][]string, error) {
	results := [][]string{}
	err := db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("", func(key, value string) bool {
			results = append(results, []string{key, value})
			return true
		})
		return err
	})
	return results, err
}

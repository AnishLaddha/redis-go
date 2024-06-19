package main

import (
	"sync"
	"time"
)

type data struct {
	val    string
	expiry time.Time
}

type Database struct {
	keyval map[string]data
	mux    sync.Mutex
}

func newDatabase() *Database {
	return &Database{
		keyval: make(map[string]data),
	}
}

func (db *Database) Set(key string, item string, exp time.Time) {
	db.mux.Lock()
	defer db.mux.Unlock()
	db.keyval[key] = data{val: item, expiry: exp}
}

func (db *Database) Get(key string) (data, bool) {
	db.mux.Lock()
	defer db.mux.Unlock()
	d, exists := db.keyval[key]
	if exists && !d.expiry.IsZero() && time.Now().After(d.expiry) {
		delete(db.keyval, key)
		exists = false
	}
	return d, exists
}

func (db *Database) Del(key string) bool {
	db.mux.Lock()
	defer db.mux.Unlock()
	_, exists := db.keyval[key]
	if exists {
		delete(db.keyval, key)
		return true
	}
	return false
}

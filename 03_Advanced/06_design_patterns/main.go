package main

import "fmt"

type Database struct {
	URL     string
	Timeout int
	Cache   bool
}

type DBOption func(*Database)

func WithTimeout(t int) DBOption {
	// A Closure! We return an anonymous function that remembers `t`.
	// When someone runs this anonymous function later, it will set Database.Timeout to `t`!
	return func(d *Database) {
		d.Timeout = t
	}
}

func WithCache(c bool) DBOption {
	return func(d *Database) {
		d.Cache = c
	}
}

// Bug Fix: You wrote `*DBOption` as the return type!
// We are building and returning a Database pointer, not an Option pointer!
func NewDatabase(url string, opts ...DBOption) *Database {
	// 1. We build the struct with safe Default Values
	db := &Database{
		URL:     url,
		Timeout: 30,
		Cache:   false,
	}
	// 2. The loop grabs every function the user passed in (e.g. `WithCache(true)`)
	// and executes it on our brand new `db` variable, mutating the struct!
	for _, optFunc := range opts {
		optFunc(db)
	}

	return db
}
func main() {
	db := NewDatabase("postgres://localhost:5432", WithCache(true))
	fmt.Printf("%+v\n", db)
}

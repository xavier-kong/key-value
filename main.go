package main

import (
	"fmt"
	"net/http"
)

type Store struct {
	Store map[string]string
	// Mu    mutex
}

var store Store

type Status enum {
	Failure
	Success
}

type OpResult struct{
	status: Status
	res: string
}

func (store Store) get(key string) {
	val, err := store.Store[key]

	if err {
	}

	return val
}

func (store Store) add(key string, value string) {
	val, err := store.Store[key]

	if val != "" {
	}

	store.Store[key] = value

	return { success }
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func (store Store) init() {
	if store.Store == nil {
		store = Store{
			Store: make(map[string]string),
		}
	}
}

func storeDispatch(w http.ResponseWriter, req *http.Request) {
	store.init()

	if req.Method == "GET" {
	}
}

func main() {
	store.init()

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/store", storeDispatch)

	http.ListenAndServe(":8090", nil)
}

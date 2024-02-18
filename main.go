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

type ResBody struct {
	key   string
	value string
}

type OpResult struct {
	success bool
	res     ResBody
}

func (store Store) get(key string) OpResult {
	val, exists := store.Store[key]

	res := OpResult{
		success: false,
		res: ResBody{
			key:   key,
			value: "",
		},
	}

	if !exists {
		res.res.value = "Key doesn't exist in store"
		return res
	}

	if val == "" {
		res.res.value = "Value is empty"
		return res
	}

	res.success = true
	res.res.value = val

	return res
}

func (store Store) add(key string, value string) OpResult {
	_, exists := store.Store[key]

	res := OpResult{
		success: false,
		res: ResBody{
			key:   key,
			value: "",
		},
	}

	if exists {
		res.res.value = "Value exists"
		return res
	}

	store.Store[key] = value

	return res
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

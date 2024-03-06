package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Store struct {
	Store map[string]string
	Mu    sync.Mutex
}

var store Store

type ResBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OpResult struct {
	Success bool    `json:"success"`
	Res     ResBody `json:"res"`
}

func (store *Store) get(key string) OpResult {
	store.Mu.Lock()
	defer store.Mu.Unlock()
	val, exists := store.Store[key]

	res := OpResult{
		Success: false,
		Res: ResBody{
			Key:   key,
			Value: "",
		},
	}

	if !exists {
		res.Res.Value = "Key doesn't exist in store"
		return res
	}

	if val == "" {
		res.Res.Value = "Value is empty"
		return res
	}

	res.Success = true
	res.Res.Value = val

	return res
}

func (store *Store) delete(key string) OpResult {
	store.Mu.Lock()
	defer store.Mu.Unlock()

	_, exists := store.Store[key]

	res := OpResult{
		Success: false,
		Res: ResBody{
			Key:   key,
			Value: "",
		},
	}

	if !exists {
		res.Res.Value = "Key doesn't exist in store"
		return res
	}

	delete(store.Store, key)

	res.Success = true

	return res
}

func (store *Store) update(key string, value string) OpResult {
	store.Mu.Lock()
	defer store.Mu.Unlock()

	_, exists := store.Store[key]

	res := OpResult{
		Success: false,
		Res: ResBody{
			Key:   key,
			Value: value,
		},
	}

	if !exists {
		res.Res.Value = "Key doesn't exist in store"
		return res
	}

	store.Store[key] = value

	res.Success = true

	return res
}

type RequestBody struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func parseRequestBody(req *http.Request) (string, string) {
	var body RequestBody
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		return "", err.Error()
	}
	return body.Key, body.Value
}

func (store *Store) add(key string, value string) OpResult {
	store.Mu.Lock()
	defer store.Mu.Unlock()

	_, exists := store.Store[key]

	res := OpResult{
		Success: false,
		Res: ResBody{
			Key:   key,
			Value: "",
		},
	}

	if exists {
		res.Res.Value = "Value exists"
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

func (store *Store) init() {
	if store.Store == nil {
		store.Store = make(map[string]string)
	}
}

func storeDispatch(w http.ResponseWriter, req *http.Request) {
	store.init()

	var res OpResult

	w.Header().Set("Content-Type", "application/json")
	key, value := parseRequestBody(req)

	if req.Method == "GET" {
		res = store.get(key)
	}

	if req.Method == "POST" {
		res = store.add(key, value)
	}

	if req.Method == "DELETE" {
		res = store.delete(key)
	}

	if req.Method == "PUT" {
		res = store.update(key, value)
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	store.init()

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/store", storeDispatch)

	http.ListenAndServe(":8090", nil)
}

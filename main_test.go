package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/hello", nil)
	responseRecorder := httptest.NewRecorder()

	hello(responseRecorder, request)

	res := responseRecorder.Result()

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading response")
	}

	if string(data) != "hello\n" {
		t.Errorf("Server not live")
	}
}

func TestGet(t *testing.T) {
	getBody := make(map[string]string)
	getBody["key"] = "test"

	body, _ := json.Marshal(getBody)

	request := httptest.NewRequest(http.MethodGet, "/store", bytes.NewReader(body))

	responseRecorder := httptest.NewRecorder()

	storeDispatch(responseRecorder, request)

	res := responseRecorder.Result()

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading response")
	}

	getRes := OpResult{}

	json.Unmarshal([]byte(string(data)), &getRes)

	fmt.Printf("test %v\n", getRes)
	fmt.Println(getRes.res.value)

	// postRequest := httptest.NewRequest(http.MethodPost, "/store", bytes.NewReader(body))

	// postResponseRecorded := httptest.NewRecorder()
}

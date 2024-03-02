package main

import (
	"bytes"
	"encoding/json"
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

	if getRes.Success != false {
		t.Errorf("expected success to be false")
	}

	if getRes.Res.Value != "Key doesn't exist in store" {
		t.Errorf("expected res value to be 'Key doesn't exist in store' found %s", getRes.Res.Value)
	}

	postRequest := httptest.NewRequest(http.MethodPost, "/store", bytes.NewReader(body))

	postResponseRecorder := httptest.NewRecorder()

	storeDispatch(postResponseRecorder, postRequest)

	// dont test the post request here, assume that it is valid

	storeDispatch(responseRecorder, request)

	newRes := responseRecorder.Result()

	defer newRes.Body.Close()

	data, err = ioutil.ReadAll(newRes.Body)

	json.Unmarshal([]byte(string(data)), &getRes)

	if getRes.Success != true {
		t.Errorf("expected success to be true")
	}

	if getRes.Res.Value != "test" {
		t.Errorf("expected value to be 'test' got '%s' instead", getRes.Res.Value)
	}
}

package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// TestJSONMarshal test the json module marshal function
func TestJSONMarshal(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/json/marshal", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "{\"author\":\"Raggaer\",\"package\":\"bison\"}" {
		t.Fatalf("Wrong json.marshal. Expected '{\"author\":\"Raggaer\",\"package\":\"bison\"}' but got '%s'", string(bodyContent))
	}
}

// TestJSONUnmarshal test the json module unmarshal function
func TestJSONUnmarshal(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/json/unmarshal", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "Author is Raggaer and package is bison" {
		t.Fatalf("Wrong json.marshal. Expected 'Author is Raggaer and package is bison' but got '%s'", string(bodyContent))
	}
}

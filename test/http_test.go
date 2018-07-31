package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// TestHTTPRedirect test the http module redirect function
func TestHTTPRedirect(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/http/redirect", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.Request.URL.String() != "https://github.com/Raggaer" {
		t.Fatalf("Wrong http.redirect URL. Expected 'https://github.com/Raggaer' but got '%s'", resp.Request.URL.String())
	}
}

func TestHTTPWrite(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/http/write", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "Testing the HTTP module" {
		t.Fatalf("Wrong http.write content. Expected 'Testing the HTTP module' but got '%s'", string(bodyContent))
	}
}

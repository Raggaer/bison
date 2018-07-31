package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

// TestHTTPWrite test the http module write function
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

func TestHTTPGetRequestMethod(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/http/request_method", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.ToLower(string(bodyContent)) != "get" {
		t.Fatalf("Wrong http.method value. Expected 'get' but got '%s'", strings.ToLower(string(bodyContent)))
	}
}

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

// TestHTTPGetRequestMethod test the http module method function
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

// TestHTTPGetRelativeURL test the http module uri function
func TestHTTPGetRelativeURL(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/http/uri", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "/http/uri" {
		t.Fatalf("Wrong http.method value. Expected '/http/uri' but got '%s'", string(bodyContent))
	}
}

// TestHTTPParam test the http module param function
func TestHTTPParam(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/http/param/raggaer", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "raggaer" {
		t.Fatalf("Wrong http.param value. Expected 'raggaer' but got '%s'", string(bodyContent))
	}
}

// TestHTTPServeFile test the http module serveFile function
func TestHTTPServeFile(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/http/serve_file", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "Serving a file from lua" {
		t.Fatalf("Wrong http.param value. Expected 'Serving a file from lua' but got '%s'", string(bodyContent))
	}
}

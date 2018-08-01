package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// TestURLQueryEscape test the url module queryEscape function
func TestURLQueryEscape(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/url/query_escape", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "Testing+bison+URL+module" {
		t.Fatalf("Wrong url.queryEscape content. Expected 'Testing+bison+URL+module' but got '%s'", string(bodyContent))
	}
}

// TestURLQueryUnescape test the url module queryUnescape function
func TestURLQueryUnescape(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/url/query_unescape", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "Testing bison URL module" {
		t.Fatalf("Wrong url.queryUnescape content. Expected 'esting bison URL module' but got '%s'", string(bodyContent))
	}
}

// TestURLPathEscape test the url module pathEscape function
func TestURLPathEscape(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/url/path_escape", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "Testing%20bison%20URL%20module" {
		t.Fatalf("Wrong url.pathEscape content. Expected '%s' but got '%s'", "Testing%20bison%20URL%20module", string(bodyContent))
	}
}

// TestURLPathUnescape test the url module pathUnescape function
func TestURLPathUnescape(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/url/path_unescape", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "Testing bison URL module" {
		t.Fatalf("Wrong url.pathUnescape content. Expected 'Testing bison URL module' but got '%s'", string(bodyContent))
	}
}

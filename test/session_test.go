package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"testing"
)

// TestSessionSet test the session module set function

func TestSessionSet(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/session/set", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("Wrong session.set status code. Expected '200' but got '%d'", resp.StatusCode)
	}
}

// TestSessionGet test the session module get function
func TestSessionGet(t *testing.T) {
	// Create a http client with a cookiejar
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{
		Jar: jar,
	}

	port := make(chan int, 1)
	defer createTestServer(port, t).Close()

	// First set a cache value
	addr := <-port
	resps, err := client.Get(fmt.Sprintf("http://localhost:%d/session/set", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resps.Body.Close()
	if resps.StatusCode != 200 {
		t.Fatalf("Wrong session.set status code. Expected '200' but got '%d'", resps.StatusCode)
	}

	// Get a cache value
	respg, err := client.Get(fmt.Sprintf("http://localhost:%d/session/get", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer respg.Body.Close()
	bodyContent, err := ioutil.ReadAll(respg.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "Testing session.get function" {
		t.Fatalf("Wrong cache.get value. Expected 'Testing session.get function' but got '%s'", string(bodyContent))
	}
}

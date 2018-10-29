package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// TestEnvironmentSet test the env set function
func TestEnvironmentSet(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/env/set", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("Wrong env.set status code. Expected '200' but got '%d'", resp.StatusCode)
	}
}

// TestEnvironmentGet test the env get function
func TestEnvironmentGet(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/env/get", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "Testing env.set function" {
		t.Fatalf("Wrong env.get value. Expected 'Testing env.set function' but got '%s'", string(bodyContent))
	}
}

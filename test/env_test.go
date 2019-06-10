package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	// First set the env variable
	respSet, err := http.Get(fmt.Sprintf("http://localhost:%d/env/set", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer respSet.Body.Close()
	if respSet.StatusCode != 200 {
		t.Fatalf("Wrong env.set status code. Expected '200' but got '%d'", respSet.StatusCode)
	}

	// Get the env variable
	respGet, err := http.Get(fmt.Sprintf("http://localhost:%d/env/get", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer respGet.Body.Close()
	bodyContent, err := ioutil.ReadAll(respGet.Body)
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(bodyContent)) != "Testing env.set function" {
		t.Fatalf("Wrong env.get value. Expected 'Testing env.set function' but got '%s'", string(bodyContent))
	}
}

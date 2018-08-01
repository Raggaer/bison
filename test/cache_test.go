package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// TestCacheSet test the cache module set function

func TestCacheSet(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/cache/set", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("Wrong cache.set status code. Expected '200' but got '%d'", resp.StatusCode)
	}
}

// TestCacheGet test the cache module get function
func TestCacheGet(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()

	// First set a cache value
	addr := <-port
	resps, err := http.Get(fmt.Sprintf("http://localhost:%d/cache/set", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resps.Body.Close()
	if resps.StatusCode != 200 {
		t.Fatalf("Wrong cache.set status code. Expected '200' but got '%d'", resps.StatusCode)
	}

	// Get a cache value
	respg, err := http.Get(fmt.Sprintf("http://localhost:%d/cache/get", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer respg.Body.Close()
	bodyContent, err := ioutil.ReadAll(respg.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "Testing cache.set function" {
		t.Fatalf("Wrong cache.get value. Expected 'Testing cache.set function' but got '%s'", string(bodyContent))
	}
}

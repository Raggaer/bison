package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// TestConfigGet test the config module get function
func TestConfigGet(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/config/get", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "testing-bison" {
		t.Fatalf("Wrong config.get value. Expected 'testing-bison' but got '%s'", string(bodyContent))
	}
}

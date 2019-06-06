package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

// TestTemplateRender test the template module render function
func TestTemplateRender(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/template/render", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "<p>Author: Raggaer</p>" {
		t.Fatalf("Wrong template.render content. Expected '<p>Author: Raggaer</p>' but got '%s'", string(bodyContent))
	}
}

// TestTemplateExecute test the template execute function
func TestTemplateExecute(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/template/execute", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	bodyStr := strings.TrimSpace(string(bodyContent))
	if bodyStr != "<p>Framework: Bison</p><p>Author: Raggaer</p>" {
		t.Fatalf("Wrong template.execute content. Expected '<p>Framework: Bison</p><p>Author: Raggaer</p>' but got '%s'", bodyStr)
	}
}

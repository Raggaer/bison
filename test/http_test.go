package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
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

// TestHTTPSetCookie test the http module setCookie function
func TestHTTPSetCookie(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/http/set_cookie", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("Wrong http.setCookie status code. Expected '200' but got %d", resp.StatusCode)
	}
}

// TestHTTPGetCookie test the http module getCookie function
func TestHTTPGetCookie(t *testing.T) {
	// Create a http client with a cookiejar
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{
		Jar: jar,
	}

	// First set a cookie
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resps, err := client.Get(fmt.Sprintf("http://localhost:%d/http/set_cookie", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resps.Body.Close()
	if resps.StatusCode != 200 {
		t.Fatalf("Wrong http.setCookie status code. Expected '200' but got %d", resps.StatusCode)
	}

	// Finally retrieve a cookie
	respg, err := client.Get(fmt.Sprintf("http://localhost:%d/http/get_cookie", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer respg.Body.Close()
	bodyContent, err := ioutil.ReadAll(respg.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "Testing http.setCookie function" {
		t.Fatalf("Wrong http.getCookie value. Expected 'Testing http.setCookie function' but got '%s'", string(bodyContent))
	}
}

// TestHTTPRemoteAddress test the http module remoteAddress function
func TestHTTPRemoteAddress(t *testing.T) {
	port := make(chan int, 1)
	defer createTestServer(port, t).Close()
	addr := <-port
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/http/remote_address", addr))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(bodyContent) != "::1" && string(bodyContent) != "127.0.0.1" {
		t.Fatalf("Wrong http.param value. Expected '::1' or '127.0.0.1' but got '%s'", string(bodyContent))
	}
}

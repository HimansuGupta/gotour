package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDevMvOnly(t *testing.T) {
	handler := http.HandlerFunc(panicDemo)
	executeRequestDevMv("Get", "/panic", devMw(handler))
}

func executeRequestDevMv(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()
	rr.Result()
	handler.ServeHTTP(rr, req)
	return rr, err
}

func TestMainFunc(t *testing.T) {
	go main()
	time.Sleep(2 * time.Second)
	//os.Exit(m.Run())

}

func TestFuncThatPanics(t *testing.T) {
	//funcThatPanics()
	assert.Panics(t, funcThatPanics, " Oh no..! Some thing went wrong..Please have look Source code..")
}

func TestMakeLink(t *testing.T) {

}

func TestHello(t *testing.T) {
	fmt.Println("=================TestHello ============")
	req, error := http.NewRequest("GET", "/hello", nil)
	if error != nil {
		t.Fatalf("Error while request   %v", error)
	}
	response := executeRequest(req, http.HandlerFunc(hello))
	val := strings.Contains(response.Body.String(), "Hello...!")
	fmt.Println(val)
	checkResponseCode(t, http.StatusOK, response.Code)
	assert.Equal(t, val, true)
	if body := response.Body.String(); body == "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestPanicAfterDemo(t *testing.T) {
	req, err := http.NewRequest("GET", "/panic-after", nil)
	if err != nil {
		t.Fatalf("Error while request   %v", err)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("checking errror", err)
		}
		fmt.Println("checking errror")
	}()
	panicAfterDemo(rec, req)
	res := rec.Result()
	checkResponseCode(t, res.StatusCode, http.StatusInternalServerError)
}

func TestDevMv(t *testing.T) {

	//handler := http.HandlerFunc(panicDemo)
	req, error := http.NewRequest("GET", "/panicDemo", nil)
	if error != nil {
		t.Fatalf("Error while request   %v", error)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("checking errror", err)
		}
		fmt.Println("checking errror")
	}()

	panicAfterDemo(rec, req)
	res := rec.Result()
	checkResponseCode(t, res.StatusCode, http.StatusInternalServerError)
	//response := executeRequest(req, handler)
	fmt.Println(rec.Body.String())

	//assert.Panics(t, response.Body.String(), " Oh no..! Some thing went wrong..Please have look Source code..")

	//executeRequest("Get", "/panic", devMw(handler))
}

func TestSourceCodeHandler_Error(t *testing.T) {
	req, err := http.NewRequest("GET", "/debug", nil)
	if err != nil {
		t.Fatalf("Error while request   %v", err)
	}
	rec := httptest.NewRecorder()

	sourceCodeHandler(rec, req)
	checkResponseCode(t, http.StatusInternalServerError, rec.Code)
}
func TestSourceCodeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/debug?line=69&path=/home/himansu/go/src/github.com/himansu/Gophercises/recover_chroma/main.go", nil)
	if err != nil {
		t.Fatalf("Error while request   %v", err)
	}
	rec := httptest.NewRecorder()

	sourceCodeHandler(rec, req)

	checkResponseCode(t, http.StatusOK, rec.Code)
}
func TestSourceCodeHandler_Failed(t *testing.T) {

	req, err := http.NewRequest("GET", "/debug?line=69&path=/home/himansu/go/src/github.com/himansu/Gophercises/recover_/main.go", nil)
	if err != nil {
		t.Fatalf("Error while request   %v", err)
	}
	rec := httptest.NewRecorder()

	sourceCodeHandler(rec, req)

	checkResponseCode(t, http.StatusInternalServerError, rec.Code)
}

func executeRequest(req *http.Request, handler http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	assert.Equal(t, expected, actual)
}

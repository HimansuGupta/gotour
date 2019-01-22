package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/himansu/Gophercises/transform/primitive"
	homedir "github.com/mitchellh/go-homedir"
)

func TestHome(t *testing.T) {
	fmt.Println("=================Test	Home============")
	req, error := http.NewRequest("GET", "/", nil)
	if error != nil {
		t.Fatalf("Error while request   %v", error)
	}
	response := executeRequest(req, http.HandlerFunc(Home))
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpload(t *testing.T) {
	fmt.Println("=================Test Upload============")
	hmdir, _ := homedir.Dir()
	imgPath := filepath.Join(hmdir, "img/gslogo.png")
	file, err := os.Open(imgPath)
	if err != nil {
		t.Error("error in opening file")
	}
	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("img", file.Name())
	if err != nil {
		t.Error("error in copy")
	}
	_, copyerr := io.Copy(part, file)
	if copyerr != nil {
		t.Error("error in copy")
	}
	err = writer.Close()
	if err != nil {
		t.Error("error in close writer")
	}

	req, error := http.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if error != nil {
		t.Fatalf("Error while request   %v", error)
	}
	response := executeRequest(req, http.HandlerFunc(Home))

	if err != nil {
		t.Errorf("=== Expected status %v", err)
	}

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUploadContentError(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img/gslogo.png")
	fmt.Println("======>", imgPath)
	file, err := os.Open(imgPath)
	if err != nil {
		t.Error("error in opening file")
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("img", file.Name())
	if err != nil {
		t.Error("error in copy")
	}
	_, err = io.Copy(part, file)
	if err != nil {
		t.Error("error in copy")
	}
	err = writer.Close()
	if err != nil {
		t.Error("error in close writer")
	}
	req, _ := http.NewRequest("POST", srv.URL+"/upload", body)
	//req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := http.DefaultClient.Do(req)
	checkResponseCode(t, http.StatusInternalServerError, res.StatusCode)
}

func executeRequest(req *http.Request, handler http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestTempFile(t *testing.T) {
	_, err := tempfile("/nofile/", "txt")
	if err == nil {
		t.Error("Expected error but got no error")
	}
}

func TestModifyMode(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/gslogo.png?mode=2", nil)
	res, _ := http.DefaultClient.Do(req)
	checkResponseCode(t, http.StatusOK, res.StatusCode)
}

func TestModifyModeWithError(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/gslogo.png?mode=ji", nil)
	res, _ := http.DefaultClient.Do(req)
	checkResponseCode(t, http.StatusBadRequest, res.StatusCode)
}

func TestModifyWithWrongFile(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/test.txt?mode=2", nil)
	res, _ := http.DefaultClient.Do(req)
	checkResponseCode(t, http.StatusBadRequest, res.StatusCode)
}

func TestModifyModeShapeNotExist(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/birds.png?mode=0&n=5", nil)
	res, _ := http.DefaultClient.Do(req)
	checkResponseCode(t, http.StatusBadRequest, res.StatusCode)
}

func TestModifyModeerrorMOde(t *testing.T) {
	srv := httptest.NewServer(GetMux())
	defer srv.Close()
	req, _ := http.NewRequest("GET", srv.URL+"/modify/birds.png?mode=3&n=a", nil)
	res, _ := http.DefaultClient.Do(req)
	checkResponseCode(t, http.StatusBadRequest, res.StatusCode)
}

func TestGenImage(t *testing.T) {
	rs := bytes.NewReader(nil)
	mode := primitive.ModeCombo
	_, err := genImage(rs, "txt", -1, mode)
	if err == nil {
		t.Error("Expected error but no error")
	}
}

func TestGenImgError(t *testing.T) {
	rs := bytes.NewReader(nil)
	opts := []genOpts{
		{N: -1, M: primitive.ModeCombo},
	}

	_, err := genImages(rs, "txt", opts...)
	if err == nil {
		t.Error("Expected error but no error")
	}
}

func TestRenderwrongchoiseError(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8888", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rs := bytes.NewReader(nil)
	rec := httptest.NewRecorder()
	renderModeChoices(rec, req, rs, "txt")
	res := rec.Result()
	checkResponseCode(t, http.StatusInternalServerError, res.StatusCode)
}

func TestRenderNoShapes(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8888", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rs := bytes.NewReader(nil)
	rec := httptest.NewRecorder()
	renderNoShapesChoices(rec, req, rs, "txt", primitive.ModeCircle)
	res := rec.Result()
	checkResponseCode(t, http.StatusInternalServerError, res.StatusCode)
}

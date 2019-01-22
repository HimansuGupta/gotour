package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func testGetEmployees(t *testing.T) {
	fmt.Println("=================TestGetEmployees============")
	req, error := http.NewRequest("GET", "/employees", nil)
	if error != nil {
		t.Fatalf("Error while request   %v", error)
	}
	response := executeRequest(req, http.HandlerFunc(GetEmployees))
	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body == "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func testGetEmployeesNoRecord(t *testing.T) {
	fmt.Println("=================TestGetEmployees NoRecord============")
	req, error := http.NewRequest("GET", "/employees", nil)
	if error != nil {
		t.Fatalf("Error while request   %v", error)
	}
	response := executeRequest(req, http.HandlerFunc(GetEmployees))
	checkResponseCode(t, http.StatusNoContent, response.Code)
	if body := response.Body.String(); body == "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func testGetEmployeeByID(t *testing.T) {
	fmt.Println("=================TestGetEmployee ByID ============")
	srv := httptest.NewServer(getMux())
	req, err := http.NewRequest("GET", srv.URL+"/employee/GS-2080", nil)
	if err != nil {
		fmt.Printf("Error in request %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("=== Expected status %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but got different status %v", res.StatusCode)
	}
}

func testGetEmployeeByIDIvalidID(t *testing.T) {
	fmt.Println("=================TestGetEmployee ByID ============")
	srv := httptest.NewServer(getMux())
	req, err := http.NewRequest("GET", srv.URL+"/employee/200", nil)
	if err != nil {
		fmt.Printf("Error in request %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("=== Expected status %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status ok but got different status %v", res.StatusCode)
	}
}

func testAddEmployee(t *testing.T) {
	fmt.Println("================= TestAddEmployee ============")

	srv := httptest.NewServer(getMux())
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	data := []byte(`{
		"ID": "GS-2080",
		"Name": "Himansu Gupta",
		"Address": "Flat No 201 Aund",
		"Practice": "IBM",
		"Team": "CB_CICD_Monitoring",
		"manager": "Neha Adkar"
	  }`)

	req, err := http.NewRequest("POST", srv.URL+"/employee", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error in request %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error in res %v", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected status ok but got different status %v", res.StatusCode)
	}
}

func testAddOneeMOreEmployee(t *testing.T) {
	fmt.Println("================= TestAddEmployee ============")

	srv := httptest.NewServer(getMux())
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	data := []byte(`{
		"ID": "GS-2082",
		"Name": "Himansu Gupta",
		"Address": "Flat No 201 Aund",
		"Practice": "IBM",
		"Team": "CB_CICD_Monitoring",
		"manager": "Neha Adkar"
	  }`)

	req, err := http.NewRequest("POST", srv.URL+"/employee", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error in request %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error in res %v", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected status ok but got different status %v", res.StatusCode)
	}
}

func testAddInvalidEmployee(t *testing.T) {
	fmt.Println("================= TestAddInvalidEmployee ============")

	srv := httptest.NewServer(getMux())
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	data := []byte(`{
		"ID": "GS-2081",
		"Name": "",
		"Address": "Flat No 201 Aund",
		"Practice": "IBM",
		"Team": "CB_CICD_Monitoring",
		"manager": "Neha Adkar"
	  }`)

	req, err := http.NewRequest("POST", srv.URL+"/employee", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error in request %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error in res %v", err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status ok but got different status %v", res.StatusCode)
	}
}

func testAddDuplicateEmployee(t *testing.T) {
	fmt.Println("================= TestAddInvalidEmployee ============")

	srv := httptest.NewServer(getMux())
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	data := []byte(`{
		"ID": "GS-2080",
		"Name": "Himansu",
		"Address": "Flat No 201 Aund",
		"Practice": "IBM",
		"Team": "CB_CICD_Monitoring",
		"manager": "Neha Adkar"
	  }`)

	req, err := http.NewRequest("POST", srv.URL+"/employee", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error in request %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error in res %v", err)
	}
	if res.StatusCode != http.StatusConflict {
		t.Errorf("Expected status ok but got different status %v", res.StatusCode)
	}
}

func testUpdateEmployee(t *testing.T) {
	fmt.Println("================= TestAddEmployee ============")

	srv := httptest.NewServer(getMux())
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	data := []byte(`{
		"ID": "GS-2080",
		"Name": "Himansu Gupta",
		"Address": "Flat No 201 Aund",
		"Practice": "IBM CICD Monitoring ",
		"Team": "CB_CICD_Monitoring",
		"manager": "Neha Adkar"
	  }`)

	req, err := http.NewRequest("PUT", srv.URL+"/updateemployee", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error in request %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error in res %v", err)
	}
	if res.StatusCode != http.StatusAccepted {
		t.Errorf("Expected status ok but got different status %v", res.StatusCode)
	}
}

func testUpdateInvalidEmployee(t *testing.T) {
	fmt.Println("================= TestAddEmployee ============")

	srv := httptest.NewServer(getMux())
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	data := []byte(`{
		"ID": "GS-2981",
		"Name": "Himansu Gupta",
		"Address": "Flat No 201 Aund",
		"Practice": "IBM",
		"Team": "CB_CICD_Monitoring",
		"manager": "Neha Adkar"
	  }`)

	req, err := http.NewRequest("PUT", srv.URL+"/updateemployee", bytes.NewBuffer(data))
	if err != nil {
		fmt.Printf("Error in request %v", err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error in res %v", err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status ok but got different status %v", res.StatusCode)
	}
}

func testDeleteEmployeeByID(t *testing.T) {
	fmt.Println("=================TestGetEmployee ByID ============")
	srv := httptest.NewServer(getMux())
	req, err := http.NewRequest("DELETE", srv.URL+"/employee/GS-2080", nil)
	if err != nil {
		fmt.Printf("Error in request %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("=== Expected status %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status ok but got different status %v", res.StatusCode)
	}
}

func testDeleteEmployeeByIDIvalidID(t *testing.T) {
	fmt.Println("=================TestGetEmployee ByID ============")
	srv := httptest.NewServer(getMux())
	req, err := http.NewRequest("DELETE", srv.URL+"/employee/200", nil)
	if err != nil {
		fmt.Printf("Error in request %v", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("=== Expected status %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status ok but got different status %v", res.StatusCode)
	}
}

func TestMaster(t *testing.T) {
	testGetEmployeesNoRecord(t)
	testAddEmployee(t)
	testAddInvalidEmployee(t)
	testAddDuplicateEmployee(t)
	testAddOneeMOreEmployee(t)
	testGetEmployees(t)
	testGetEmployeeByID(t)
	testGetEmployeeByIDIvalidID(t)
	testUpdateEmployee(t)
	testUpdateInvalidEmployee(t)
	testDeleteEmployeeByID(t)
	testDeleteEmployeeByIDIvalidID(t)
}

func executeRequest(req *http.Request, handler http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	assert.Equal(t, expected, actual)
}

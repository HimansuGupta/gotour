package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var str = `{
	"ID": "GS-2080",
	"Name": "Himansu Gupta",
	"Address": "Flat No 201 Aund",
	"Practice": "IBM",
	"Team": "CB_CICD_Monitoring",
	"manager": "Neha Adkar"
  }`

//TestJSONToStructPass convert JSOn to Struct
func TestJSONToStructPass(t *testing.T) {

	employee, error := JSONToStruct(str)
	if error != nil {
		t.Error("Expected data in Employee Struct:- Error ", error)
	} else {
		fmt.Println("employee ", employee)
	}
}

//TestJSONToStruct test for Fail
func TestJSONToStruct(t *testing.T) {
	str := `{
		"ID": "GS-2080",
		"Name": "Himansu Gupta",
 		"Address": "Flat No 201 Aund,
 		"Practice": "IBM",
		"Team": "CB_CICD_Monitoring",
 		"manager": "Neha Adkar"
	  }`
	employee, error := JSONToStruct(str)
	if error != nil {
		fmt.Println("error ", error)
	} else {
		fmt.Println("employee ", employee)
	}

}

// TestCalculate Pre difine function  test
func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}

}

// TestCalculate Pre difine function  test with Assert Test
func TestAssertCalculate(t *testing.T) {
	assert.Equal(t, Calculate(2), 4)
}

// TestCalculate Pre difine function  test with struct
func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{-5, -3},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

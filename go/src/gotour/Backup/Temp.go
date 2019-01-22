package Temptest

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var str = `{
// 	"ID": "GS-2080",
// 	"Name": "Himansu Gupta",
// 	"Address": "Flat No 201 Aund",
// 	"Practice": "IBM",
// 	"Team": "CB_CICD_Monitoring",
// 	"manager": "Neha Adkar"
//   }`

// func TestJSONToStructPass(t *testing.T) {

// 	employee, error := JSONToStruct(str)
// 	if error != nil {
// 		t.Error("Expected data in Employee Struct:- Error ", error)
// 	} else {
// 		fmt.Println("employee ", employee)
// 	}
// }

// func TestJSONToStruct(t *testing.T) {
//'","source": "go","startLineNumber": 28,"startColumn": 4,"endLineNumber": 28,"endColumn": 4} 	str := `{
//'","source": "go","startLineNumber": 28,"startColumn": 4,"endLineNumber": 28,"endColumn": 4} 		"ID": "GS-2080",
//'","source": "go","startLineNumber": 28,"startColumn": 4,"endLineNumber": 28,"endColumn": 4} 		"Name": "Himansu Gupta",
//'","source": "go","startLineNumber": 28,"startColumn": 4,"endLineNumber": 28,"endColumn": 4} 		"Address": "Flat No 201 Aund,
//'","source": "go","startLineNumber": 28,"startColumn": 4,"endLineNumber": 28,"endColumn": 4} 		"Practice": "IBM",
//'","source": "go","startLineNumber": 28,"startColumn": 4,"endLineNumber": 28,"endColumn": 4} 		"Team": "CB_CICD_Monitoring",
//'","source": "go","startLineNumber": 28,"startColumn": 4,"endLineNumber": 28,"endColumn": 4} 		"manager": "Neha Adkar"
// 	  }`
// 	employee, error := JSONToStruct(str)
// 	if error != nil {
// 		fmt.Println("error ", error)
// 	} else {
// 	'","source": "go","startLineNumber": 28,"startColumn": 4,"endLineNumber": 28,"endColumn": 4}	fmt.Println("employee ", employee)
// 	'","source": "go","startLineNumber": 28,"startColumn": 4,"endLineNumber": 28,"endColumn": 4}}
// }'","source": "go","startLineNumber": 28,"startColumn": 4,"endLineNumber": 28,"endColumn": 4}
//'","source": "go","startLineNumber": 28,"startColumn": 4,"endLineNumber": 28,"endColumn": 4}
// f'","source": "go","startLineNumber": 28,"startColumn": 4,"endLineNumber": 28,"endColumn": 4}unc TestCalculate(t *testing.T) {
// 	'","source": "go","startLineNumber": 28,"startColumn": 4,"endLineNumber": 28,"endColumn": 4}assert.Equal(t, Calculate(2), 4)
// }

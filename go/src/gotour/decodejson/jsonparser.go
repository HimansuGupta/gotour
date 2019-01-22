package decodejson

import "fmt"
import "encoding/json"

//Employee Details Structure
type Employee struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Practice string `json:"practice"`
	Team     string `json:"team"`
	Manager  string `json:"manager"`
}

//Response custom response
type Response struct {
	Status  string     `json:"status"`
	Code    string     `json:"code"`
	Message string     `json:"message"`
	Data    []Employee `json:"data"`
}

// JSONTOStruct to Predefine Structure
func JSONTOStruct(jsonString string) (*Employee, error) {
	var employee Employee
	error := json.Unmarshal([]byte(jsonString), &employee)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%+v\n", employee)
	}
	return &employee, error
}

//IsValidRequest have mandatry data or not
func IsValidRequest(employee *Employee) bool {
	flag := false
	if employee.ID != "" && employee.Name != "" {
		return true
	}
	return flag
}

// fmt.Printf("%+v\n", employee.ID)
// fmt.Printf("%+v\n", employee.Name)
// fmt.Printf("%+v\n", employee.Address)
// fmt.Printf("%+v\n", employee.Practice)
// fmt.Printf("%+v\n", employee.Team)
// fmt.Printf("%+v\n", employee.Manager)

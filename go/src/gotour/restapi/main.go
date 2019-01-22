package main

// Package controller Dash Testrail API
//
// The purpose of this application is to store and retrieve Employee records
//
//
//
//     BasePath: /api/testrail/v1
//     Version: 0.0.1
//     License: Copy Right © 2019 GG Lab Private Limited

//		All rights reserved.
//		This material is confidential and proprietary to GS Lab Private Limited.
//		No part of this material should be reproduced, published in any form by any means,
//		electronic or mechanical including photocopy or any information storage or
//		retrieval system nor shougo:generate swagger generate spec -m -o ./swagger.json
//ld the material be disclosed to third parties without
//		the express written authorization of GS Lab GS LabPrivate Limited.Licensed Materials -

//
//     Contact: Himansu Gupta <himansu.gupta@gslab.com>
//
//     Consumes:
//       - application/json
//
//     Produces:
//       - application/json
//
//     Security:
//       - token:
//
//     SecurityDefinitions:
//       token:
//         type: apiKey
//         in: header
//         name: Authorization
//
//
// 		swagger:meta

//		go:generate swagger generate spec -m -o ./swagger.json

import (
	"encoding/json"
	"fmt"
	"gotour/databaseconnection"
	"gotour/decodejson"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetEmployees from db
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Employees  Implement")
	//var employees []decodejson.Employee
	var response decodejson.Response
	err := databaseconnection.GetConnection()
	if err != nil {
		response.Status = "Time out"
		response.Code = "504"
		response.Message = "The server is unavailable to handle this request right now..!!"
		w.WriteHeader(http.StatusGatewayTimeout)
		json.NewEncoder(w).Encode(&response)
	} else {
		w.Header().Add("Content-Type", "application/json")
		employees, error := databaseconnection.FindAllEmployees()
		if error != nil {
			fmt.Println("if DB Error Come")
			response.Status = "Time out"
			response.Code = "504"
			response.Message = "The server is unavailable to handle this request right now..!!"
			w.WriteHeader(http.StatusGatewayTimeout)
			json.NewEncoder(w).Encode(&response)
		} else {
			if len(employees) > 0 {
				//fmt.Println("if DB Data Come")
				response.Status = "SUCCESS"
				response.Code = "200"
				response.Message = "Get Record successfuly..!!"
				response.Data = employees
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(&response)
			} else {
				//fmt.Println("if DB No Record ")
				response.Status = "No Content"
				response.Code = "204"
				response.Message = "DB is Empty,Record Not Found in DB..!!"
				response.Data = employees
				w.WriteHeader(http.StatusNoContent)
				json.NewEncoder(w).Encode(&response)
			}
		}
	}
}

// GetEmployeeByID from db
func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET Employee By ID Implement")
	var employees []decodejson.Employee
	//var employe decodejson.Employee
	var response decodejson.Response
	//databaseconnection.GetConnection()
	params := mux.Vars(r)
	//fmt.Println(len(params))
	//fmt.Println("Searchin ID ", params["id"])
	w.Header().Add("Content-Type", "application/json")

	err := databaseconnection.GetConnection()
	if err != nil {
		response.Status = "Time out"
		response.Code = "504"
		response.Message = "The server is unavailable to handle this request right now..!!"
		w.WriteHeader(http.StatusGatewayTimeout)
		json.NewEncoder(w).Encode(&response)
	} else {

		flag, error := databaseconnection.IsRecordExists(params["id"])
		if error != nil {
			//	fmt.Println("In Error DB")
			response.Status = "Time out"
			response.Code = "504"
			response.Message = "The server is unavailable to handle this request right now..!!"
			w.WriteHeader(http.StatusGatewayTimeout)
			json.NewEncoder(w).Encode(&response)
		}
		if flag != false {
			employee, error := databaseconnection.FindByID(params["id"])

			if error != nil {
				response.Status = "Time out"
				response.Code = "504"
				response.Message = "The server is unavailable to handle this request right now..!!"
				w.WriteHeader(http.StatusGatewayTimeout)
				json.NewEncoder(w).Encode(&response)
			} else {
				employees = append(employees, employee)
				response.Status = "SUCCESS"
				response.Code = "200"
				response.Message = "Get Record successfuly..!!"
				response.Data = employees
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(&response)
			}
		} else {
			response.Status = "Failed"
			response.Code = "404"
			response.Message = "Invalid ID,Record Not Found in DB..!!"
			w.WriteHeader(http.StatusNotFound)
			//response.Data = employees
			json.NewEncoder(w).Encode(&response)
		}
	}
}

// AddNewEmployee from db
func AddNewEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POSt -- Add Employee Implement")
	var employees []decodejson.Employee
	var employee decodejson.Employee
	var response decodejson.Response
	var decodeerror error
	var dberror error
	var flag bool
	decodeerror = json.NewDecoder(r.Body).Decode(&employee)
	//employees[0] = employee
	if decodeerror != nil {
		response.Status = "Bad Request"
		response.Code = "400"
		response.Message = "Please have look API RequestBody..!!"
		response.Data = employees
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	} else {
		err := databaseconnection.GetConnection()
		if err != nil {
			response.Status = "Time out"
			response.Code = "504"
			response.Message = "The server is unavailable to handle this request right now..!!"
			w.WriteHeader(http.StatusGatewayTimeout)
			json.NewEncoder(w).Encode(&response)
		} else {
			if decodejson.IsValidRequest(&employee) {
				flag, dberror = databaseconnection.IsRecordExists(employee.ID)

				if flag {
					response.Status = "Conflict"
					response.Code = "409"
					response.Message = "Record alredy exists in DB..!!"
					response.Data = employees
					w.WriteHeader(http.StatusConflict)
					json.NewEncoder(w).Encode(response)
				} else {
					if dberror != nil {
						response.Status = "Time out"
						response.Code = "504"
						response.Message = "The server is unavailable to handle this request right now..!!"
						w.WriteHeader(http.StatusGatewayTimeout)
						json.NewEncoder(w).Encode(&response)
					} else {
						employee, error := databaseconnection.AddEmployee(employee)
						//		fmt.Println(employee)
						employees = append(employees, employee)
						if error != nil {
							response.Status = "Time out"
							response.Code = "504"
							response.Message = "The server is unavailable to handle this request right now..!!"
							w.WriteHeader(http.StatusGatewayTimeout)
							json.NewEncoder(w).Encode(&response)
						} else {
							response.Status = "SUCCESS"
							response.Code = "200"
							response.Message = "Record added Successfuly..!!"
							response.Data = employees
							w.WriteHeader(http.StatusCreated)
							json.NewEncoder(w).Encode(response)
						}
					}
				}
			} else {
				response.Status = "Bad Request"
				response.Code = "400"
				response.Message = "Request Missing ID or Name!!"
				response.Data = employees
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
			}
		}
	}
}

//UpdateEmployeeByID update Employe
func UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POSt -- Update Employees Implement")
	var employees []decodejson.Employee
	var employee decodejson.Employee
	var response decodejson.Response
	var decodeerror error
	var dberror error
	var flag bool
	decodeerror = json.NewDecoder(r.Body).Decode(&employee)
	//employees[0] = employee
	if decodeerror != nil {
		response.Status = "Bad Request"
		response.Code = "400"
		response.Message = "Please have look API RequestBody..!!"
		response.Data = employees
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	} else {
		err := databaseconnection.GetConnection()
		if err != nil {
			response.Status = "Time out"
			response.Code = "504"
			response.Message = "The server is unavailable to handle this request right now..!!"
			w.WriteHeader(http.StatusGatewayTimeout)
			json.NewEncoder(w).Encode(&response)
		} else {
			if decodejson.IsValidRequest(&employee) {
				flag, dberror = databaseconnection.IsRecordExists(employee.ID)

				if !flag {
					response.Status = "Bad Request"
					response.Code = "404"
					response.Message = "Invalid ID, Record Not found in DB..!!"
					response.Data = employees
					w.WriteHeader(http.StatusNotFound)
					json.NewEncoder(w).Encode(response)
				} else {
					if dberror != nil {
						response.Status = "Time out"
						response.Code = "504"
						response.Message = "The server is unavailable to handle this request right now..!!"
						w.WriteHeader(http.StatusGatewayTimeout)
						json.NewEncoder(w).Encode(&response)
					} else {
						employee, error := databaseconnection.UpdateByID(employee)
						//fmt.Println(employee)
						employees = append(employees, employee)
						if error != nil {
							response.Status = "Time out"
							response.Code = "504"
							response.Message = "The server is unavailable to handle this request right now..!!"
							w.WriteHeader(http.StatusGatewayTimeout)
							json.NewEncoder(w).Encode(&response)
						} else {
							response.Status = "SUCCESS"
							response.Code = "202"
							response.Message = "Record Updated Successfuly..!!"
							response.Data = employees
							w.WriteHeader(http.StatusAccepted)
							json.NewEncoder(w).Encode(response)
						}
					}
				}
			} else {
				response.Status = "Bad Request"
				response.Code = "400"
				response.Message = "Request Missing ID or Name!!"
				response.Data = employees
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
			}
		}
	}

	// var employee decodejson.Employee
	// error := json.NewDecoder(r.Body).Decode(&employee)
	// if error == nil {
	// 	databaseconnection.GetConnection()
	// 	employee, error := databaseconnection.UpdateByID(employee)
	// 	if error == nil {
	// 		fmt.Println("Employee Updated successfuly ", employee)
	// 	} else {
	// 		fmt.Println("Error :- ", error)
	// 	}
	// 	json.NewEncoder(w).Encode(employee)
	// }
}

// DeleteEmployee based on ID
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Employees  Implement")
	//var employees []decodejson.Employee
	//var employe decodejson.Employee
	var response decodejson.Response
	//databaseconnection.GetConnection()
	params := mux.Vars(r)
	fmt.Println(len(params))
	fmt.Println(params["id"])
	w.Header().Add("Content-Type", "application/json")
	err := databaseconnection.GetConnection()
	if err != nil {
		response.Status = "Time out"
		response.Code = "504"
		response.Message = "The server is unavailable to handle this request right now..!!"
		w.WriteHeader(http.StatusGatewayTimeout)
		json.NewEncoder(w).Encode(&response)
	} else {
		flag, error := databaseconnection.IsRecordExists(params["id"])
		if error != nil {
			//	fmt.Println("In Error DB")
			response.Status = "Time out"
			response.Code = "504"
			response.Message = "The server is unavailable to handle this request right now..!!"
			w.WriteHeader(http.StatusGatewayTimeout)
			json.NewEncoder(w).Encode(&response)
		}
		if flag != false {
			error := databaseconnection.DeleteByID(params["id"])
			if error != nil {
				response.Status = "Time out"
				response.Code = "504"
				response.Message = "The server is unavailable to handle this request right now..!!"
				w.WriteHeader(http.StatusGatewayTimeout)
				json.NewEncoder(w).Encode(&response)
			} else {
				response.Status = "SUCCESS"
				response.Code = "200"
				response.Message = "Record Deleted Successfuly..!!"
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(&response)
			}
		} else {
			response.Status = "Failed"
			response.Code = "404"
			response.Message = "Invalid ID,Record Not Found in DB..!!"
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(&response)
		}
	}
}

func getMux() *mux.Router {
	log.Println("In getMux / Handler")
	router := mux.NewRouter()
	router.HandleFunc("/employees", GetEmployees).Methods("GET")
	router.HandleFunc("/employee", AddNewEmployee).Methods("POST")
	router.HandleFunc("/employee/{id}", GetEmployeeByID).Methods("GET")
	router.HandleFunc("/updateemployee", UpdateEmployeeByID).Methods("PUT")
	router.HandleFunc("/employee/{id}", DeleteEmployee).Methods("DELETE")
	return router
}

//main entry define all api uri and methods
func main() {
	log.Println("In Main/Controller")
	log.Fatal(http.ListenAndServe(":8080", getMux()))
}

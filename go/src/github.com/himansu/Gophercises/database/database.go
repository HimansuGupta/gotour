package database

import (
	"fmt"

	decodejson "github.com/himansu/restapi/apidecodejson"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//EmployeeDB for contain Server,DB and Collection name
type EmployeeDB struct {
	Server     string
	Database   string
	Collection string
}

//Mydb mgo.Connect.Database
type Mydb struct {
	db *mgo.Database
}

//var Session *mgo.Database

// SetConnectionCredential database and collection
func (emp *EmployeeDB) SetConnectionCredential(host, dbName, collectionName string) {
	emp.Server = host
	emp.Database = dbName
	emp.Collection = collectionName
}

//var db *Database

// GetConnection to database using EmployeeDB Stuct
func (emp *EmployeeDB) GetConnection() (*mgo.Database, error) {
	session, err := mgo.Dial(emp.Server)
	if err != nil {
		return nil, err
	}
	database := session.DB(emp.Database)
	return database, err
}

// FindAllEmployees details
func (mydb *Mydb) FindAllEmployees(empdb *EmployeeDB) ([]decodejson.Employee, error) {
	fmt.Println("Start Find All")
	var employees []decodejson.Employee
	err := mydb.db.C(empdb.Collection).Find(bson.M{}).All(&employees)
	return employees, err
}

// FindByID details
func (mydb *Mydb) FindByID(empdb *EmployeeDB, id string) (decodejson.Employee, error) {
	var employee decodejson.Employee
	err := mydb.db.C(empdb.Collection).Find(bson.M{"id": id}).One(&employee)
	return employee, err
}

// IsRecordExists find record exists
func (mydb *Mydb) IsRecordExists(empdb *EmployeeDB, id string) (bool, error) {
	count, err := mydb.db.C(empdb.Collection).Find(bson.M{"id": id}).Count()
	if count > 0 {
		return true, err
	}
	return false, err
}

//AddEmployee new Employee in employeedb
func (mydb *Mydb) AddEmployee(empdb *EmployeeDB, employee decodejson.Employee) (decodejson.Employee, error) {
	err := mydb.db.C(empdb.Collection).Insert(&employee)
	return employee, err
}

//UpdateByID Employee can edit/update his detail
func (mydb *Mydb) UpdateByID(empdb *EmployeeDB, employee decodejson.Employee) (decodejson.Employee, error) {
	err := mydb.db.C(empdb.Collection).Update(bson.M{"id": employee.ID}, &employee)
	return employee, err
}

// DeleteByID employee details
func (mydb *Mydb) DeleteByID(empdb *EmployeeDB, id string) error {
	err := mydb.db.C(empdb.Collection).Remove(bson.M{"id": id})
	return err
}

func mainTemp() {
	var empdbcred EmployeeDB
	empdbcred.SetConnectionCredential("localhost", "employeedb", "employeedetails")
	var tempdbb Mydb
	Session, err := empdbcred.GetConnection()
	tempdbb.db = Session
	if err != nil {
		fmt.Println(err)
	}
	emp, err := tempdbb.FindAllEmployees(&empdbcred)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(emp)

}

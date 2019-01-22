package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/himansu/Gophercises/task/db"
	homedir "github.com/mitchellh/go-homedir"

	"github.com/stretchr/testify/assert"
)

func setDB() string {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "test.db")
	return dbPath
}
func removeDBFile() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "test.db")
	os.Remove(dbPath)
}

func TestAddCommand(t *testing.T) {
	dbPath := setDB()
	db.Init(dbPath)

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)

	oldStdout := os.Stdout
	os.Stdout = record
	taskName := []string{"TaskTwo"}
	addCmd.Run(addCmd, taskName)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)

	if err != nil {
		t.Error("error occured while test case run :: ", err)
	}
	output := string(content)
	fmt.Println("Himansu   ", output)
	//val := strings.Contains(output, "Added "+taskName[0]+" to your task list.")
	val := strings.Contains(output, taskName[0])
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Dba.Close()
}

func TestListCommand(t *testing.T) {
	dbPath := setDB()
	db.Init(dbPath)

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record

	listCmd.Run(listCmd, nil)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run ::  ", err)
	}

	output := string(content)
	val := strings.Contains(output, "TaskTwo")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Dba.Close()
}
func TestDoCommandParseError(t *testing.T) {
	dbPath := setDB()
	db.Init(dbPath)

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record
	a := []string{"a"}

	doCmd.Run(doCmd, a)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run ::  ", err)
	}

	output := string(content)
	val := strings.Contains(output, "Failed to parse")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Dba.Close()
}

func TestDoCommandInvalidTaskError(t *testing.T) {
	dbPath := setDB()
	db.Init(dbPath)
	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record
	a := []string{"1000"}

	doCmd.Run(doCmd, a)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run ::  ", err)
	}

	output := string(content)
	val := strings.Contains(output, "Invalid task number")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Dba.Close()
}
func TestDoCommand(t *testing.T) {
	dbPath := setDB()
	db.Init(dbPath)

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record
	taskNum := []string{"1"}

	doCmd.Run(doCmd, taskNum)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run ::  ", err)
	}

	output := string(content)
	val := strings.Contains(output, "Marked")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Dba.Close()
}

func TestListCommand_Notask(t *testing.T) {
	removeDBFile()
	dbPath := setDB()
	db.Init(dbPath)

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record
	//a := []string{"Task123"}
	listCmd.Run(listCmd, nil)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run :: ", err)
	}

	output := string(content)
	val := strings.Contains(output, "You have no tasks to complete! Why not take a vacation? 🏖")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()
	db.Dba.Close()
}

func TestListCommand_DbError(t *testing.T) {
	removeDBFile()
	dbPath := setDB()
	db.Init(dbPath)
	db.Dba.Close()

	record, _ := os.OpenFile("test_stdout.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = record

	listCmd.Run(listCmd, nil)
	record.Seek(0, 0)
	content, err := ioutil.ReadAll(record)
	if err != nil {
		t.Error("error occured while test case run :: ", err)
	}

	output := string(content)
	val := strings.Contains(output, "Something went wrong with DataBase:")
	assert.Equalf(t, true, val, "they should be equal")
	record.Truncate(0)
	record.Seek(0, 0)
	os.Stdout = oldStdout
	record.Close()

}

// func TestMain(m *testing.M) {
// 	dashtest.ControlCoverage(m)
// }

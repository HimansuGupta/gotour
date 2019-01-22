package main

import (
	"path/filepath"
	"testing"

	"github.com/himansu/Gophercises/task/db"
	homedir "github.com/mitchellh/go-homedir"
)

func TestMain(m *testing.M) {
	main()

}
func TestMainError(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	db.Init(dbPath)
	main()
}

// func TestMainHomeDirError(t *testing.T) {
// 	tempDir := HomedirVar
// 	HomedirVar = func() (string, error) {
// 		return "", errors.New("Customised error")
// 	}
// 	main()
// 	defer func() {
// 		HomedirVar = tempDir
// 	}()

// }

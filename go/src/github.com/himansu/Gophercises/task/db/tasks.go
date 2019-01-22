package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")

//Dba is required for testing. thats by intilise caps
var Dba *bolt.DB

// Task struct with Id and Name
type Task struct {
	Key   int
	Value string
}

//Init its use to connect DB and create bucket if not exists
func Init(dbPath string) error {
	var err error
	Dba, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return Dba.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

//CreateTask  Task in task list
func CreateTask(task string) (int, error) {
	var id int
	err := Dba.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	//db.Close()
	return id, nil
}

//AllTasks Get List of all remaining task
func AllTasks() ([]Task, error) {
	var tasks []Task
	err := Dba.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//DeleteTask using task id
func DeleteTask(key int) error {
	return Dba.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

//itob its use to conversion int to byte
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

//btoi its use to conversion byte to int
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

package random

import (
	"math/rand"
)

//GetRandom Integer, Funv return int64
func GetRandom() int {
	//fmt.Println("My favorite number is", rand.Intn(10))
	return rand.Intn(10)
}

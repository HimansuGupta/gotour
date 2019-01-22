package sandbox

import "fmt"
import "time"

// Sandboxfun for testing calling out side
func Sandboxfun() {
	fmt.Println("I am in SandBOX")
	fmt.Println(time.Now())
	internalsand()
}

// Temp for Test var
var Temp = internalsand

//Test Internat
func internalsand() {
	fmt.Println("Test")
}

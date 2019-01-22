package calculate

import (
	"fmt"
	"math"
)

//GetSquareRoot of any digit
func GetSquareRoot(value float64) float64 {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(value))
	return math.Sqrt(value)
}

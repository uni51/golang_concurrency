package calculator

import "fmt"

var offset float64 = 1 // Private 変数
var Offset float64 = 1 // Public 変数

func Sum(a float64, b float64) float64 {
	fmt.Println("multiply: ", multiply(a, b))
	fmt.Println("Multiply: ", Multiply(a, b))
	return a + b + offset
}

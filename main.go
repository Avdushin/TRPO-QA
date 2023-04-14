package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Println("Введите коэффициенты квадратного уравнения ax^2 + bx + c = 0:")
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Print("b: ")
	fmt.Scan(&b)
	fmt.Print("c: ")
	fmt.Scan(&c)

	if a == 0 {
		fmt.Println("Неразрешимое уравнение")
		return
	}

	root1, root2, D := solveQE(a, b, c)

	// если x1 или x2 == 0 -> нулевой корень
	if root1 == 0 || root2 == 0 {
		fmt.Printf("\nНулевой корень\n\nx1 = %.2f, x2 = %.2f\n", root1, root2)
		fmt.Println("\nДискриминант: ", D)
		return
	}

	switch {
	case D < 0:
		fmt.Println("\nУравнение не имеет корней...")
	case D == 0:
		fmt.Printf("\nУравнение имеет один корень: x = %v\n", root1)
	default:
		fmt.Printf("\nУравнение имеет два корня: x1 = %.2f, x2 = %.2f\n", root1, root2)
	}
	fmt.Println("\nДискриминант: ", D)
}

// решаем уравнение по формуле
func solveQE(a, b, c float64) (float64, float64, float64) {
	// формула дискриминанта
	D := b*b - 4*a*c

	// корни уравнения
	root1 := (-b + math.Sqrt(D)) / (2 * a)
	root2 := (-b - math.Sqrt(D)) / (2 * a)

	return root1, root2, D
}

// сравниваем 2 числа float64
func floatEquals(a, b, epsilon float64) bool {
	if math.IsNaN(a) && math.IsNaN(b) {
		return true
	}
	return math.Abs(a-b) < epsilon
}

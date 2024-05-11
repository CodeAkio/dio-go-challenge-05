package main

import "fmt"

// Funções de operações básicas
func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Teste de Operações da Calculadora:")
	fmt.Printf("5 + 3 = %v\n", add(5, 3))
	fmt.Printf("5 - 3 = %v\n", subtract(5, 3))
	fmt.Printf("5 * 3 = %v\n", multiply(5, 3))
	result, err := divide(5, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("5 / 3 = %v\n", result)
	}
}

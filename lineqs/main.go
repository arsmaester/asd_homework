package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readMatrix(filename string) ([][]float64, []float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var matrix [][]float64
	var b []float64

	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		var row []float64
		for i, num := range numbers {
			val, err := strconv.ParseFloat(num, 64)
			if err != nil {
				return nil, nil, err
			}
			if i < len(numbers)-1 {
				row = append(row, val)
			} else {
				b = append(b, val)
			}
		}
		matrix = append(matrix, row)
	}

	return matrix, b, scanner.Err()
}

func gaussianElimination(matrix [][]float64, b []float64) []float64 {
	n := len(matrix)

	// convert to upper-triangular
	for i := 0; i < n; i++ {

		for k := i + 1; k < n; k++ {
			ratio := matrix[k][i] / matrix[i][i]
			for j := i; j < n; j++ {
				matrix[k][j] -= ratio * matrix[i][j]
			}
			b[k] -= ratio * b[i]
		}
	}

	// finding variables
	x := make([]float64, n)
	for i := n - 1; i >= 0; i-- {
		sum := b[i]
		for j := i + 1; j < n; j++ {
			sum -= matrix[i][j] * x[j]
		}
		x[i] = sum / matrix[i][i]
	}

	return x
}

func main() {
	matrix, b, err := readMatrix("input.txt")
	if err != nil {
		fmt.Println("Error reading matrix:", err)
		return
	}

	solution := gaussianElimination(matrix, b)
	fmt.Println("Solution:")
	for i, val := range solution {
		fmt.Printf("x%d = %.6f\n", i+1, val)
	}
}

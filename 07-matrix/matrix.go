package main

import "fmt"

func main() {
	var matrixA = [][]int{{1, 2, 3}, {4, 5, 6}}
	var matrixB = [][]int{{2, 3}, {4, 5}, {6, 1}}
	var matrixC [2][2]int
	if len(matrixA) == len(matrixB[0]) {
		fmt.Println("Rows in matrixA:", len(matrixA), "is match to number of column in matrixB:", len(matrixB[0]))
	} else {
		fmt.Println("Incompatible dimension")
	}
	i := 0
	for i < len(matrixA) {
		j := 0
		for j < len(matrixA[1]) {
			fmt.Println(matrixA[i][j])
			j++
		}
		i++
	}

	i = 0
	for i < len(matrixB) {
		j := 0
		for j < len(matrixB[0]) {
			fmt.Println(matrixB[i][j])
			j++
		}
		i++
	}

	i = 0
	for i < len(matrixA) {
		j := 0
		for j < len(matrixA) {
			k := 0
			for k < len(matrixA[0]) {
				fmt.Println("C[", i, j, k, "]:", matrixA[i][k], matrixB[k][j])
				matrixC[i][j] = matrixC[i][j] + matrixA[i][k]*matrixB[k][j]
				k++
			}
			j++
			fmt.Println(i, j, k)
		}

		i++
	}
	fmt.Println(matrixC)
}

/*
	var matrixFinalC [][]int
	for i, matA := range matrixA {
		for j, matB := range matrixB {
			matrixFinalC = matA[0] * matB[0]+matA[1]
			j++
		}
		i++
	}
*/

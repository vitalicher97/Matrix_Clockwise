package main

import "fmt"

func clockwiseMatrix(matrix [][]int) ([]int, bool) {

	var i int // iterator

	startRowIndex := 0
	startColIndex := 0
	endRowIndex := len(matrix[0])
	endColIndex := len(matrix)

	if (endColIndex != endRowIndex) {
		emptySlice := []int{0}
		return emptySlice, false
	}

	resultSlice := make([]int, endRowIndex * endColIndex)

	slIndex := 0
	for startRowIndex < endRowIndex && startColIndex < endColIndex {

		for i = startColIndex; i < endColIndex; i++ {
			resultSlice[slIndex] = matrix[startRowIndex][i]
			slIndex++
		}
		startRowIndex++

		for i = startRowIndex; i < endRowIndex; i++ {
			resultSlice[slIndex] = matrix[i][endColIndex - 1]
			slIndex++
		}
		endColIndex--

		if startRowIndex < endRowIndex {
			for i = endColIndex - 1; i >= startColIndex; i-- {
				resultSlice[slIndex] = matrix[endRowIndex - 1][i]
				slIndex++
			}
			endRowIndex--
		}

		if startColIndex < endColIndex {
			for i = endRowIndex - 1; i >= startRowIndex; i-- {
				resultSlice[slIndex] = matrix[i][startColIndex]
				slIndex++
			}
			startColIndex++
		}

	}

	return resultSlice, true


}

func main() {

	twoDSlice := make([][]int, 4)

	for i := range twoDSlice {
		twoDSlice[i] = make([]int, 4)
	}

	value := 1

	for i := 0; i < cap(twoDSlice); i++ {
		for j := 0; j < cap(twoDSlice[0]); j++ {
			twoDSlice[i][j] = value
			value++
		}
	}

	for i := 0; i < len(twoDSlice); i++ {
		for j := 0; j < (len(twoDSlice[0])); j++ {
			fmt.Print(twoDSlice[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}

	fmt.Println()



	resultMatrix, flag := clockwiseMatrix(twoDSlice)

	if flag {
		for i := 0; i < len(resultMatrix); i++ {
			fmt.Print(resultMatrix[i])
			fmt.Print(" ")
		}
	} else {
		fmt.Println("Matrix must be n*n")
	}



}

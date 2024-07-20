package main

import (
	"fmt"
	"sync"
	"time"
)

func multiplyMatrices(n [][]int, m [][]int) ([][]int, error) {
	nRows, nCols := len(n), len(n[0])
	mCols := len(m[0])

	result := make([][]int, nRows)
	for i := 0; i < nRows; i++ {
		result[i] = make([]int, mCols)
	}

	for resultRow := 0; resultRow < nRows; resultRow++ {
		for resultCol := 0; resultCol < mCols; resultCol++ {
			total := 0
			for nCol := 0; nCol < nCols; nCol++ {
				total += n[resultRow][nCol] * m[nCol][resultCol]
			}
			result[resultRow][resultCol] = total
		}
	}

	return result, nil
}

func multiplyMatricesParallel(n, m [][]int) ([][]int, error) {
	nRows, nCols := len(n), len(n[0])
	mCols := len(m[0])

	result := make([][]int, nRows)
	for i := 0; i < nRows; i++ {
		result[i] = make([]int, mCols)
	}

	var wg sync.WaitGroup
	wg.Add(nRows * mCols)

	for resultRow := 0; resultRow < nRows; resultRow++ {
		for resultCol := 0; resultCol < mCols; resultCol++ {
			go func(resultRow, resultCol int) {
				defer wg.Done()
				total := 0
				for nCol := 0; nCol < nCols; nCol++ {
					total += n[resultRow][nCol] * m[nCol][resultCol]
				}
				result[resultRow][resultCol] = total
			}(resultRow, resultCol)
		}
	}

	wg.Wait()

	return result, nil
}

func main() {
	size := 700
	n := make([][]int, size)
	for i := 0; i < size; i++ {
		n[i] = make([]int, size)
		for j, _ := range n[i] {
			n[i][j] = 2
		}
	}
	start := time.Now()
	m, _ := multiplyMatrices(n, n)
	fmt.Println(time.Since(start))
	_ = m

	start = time.Now()
	m, _ = multiplyMatricesParallel(n, n)
	fmt.Println(time.Since(start))
}

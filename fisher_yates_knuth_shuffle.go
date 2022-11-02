package shuffle

import (
	"math/rand"
	"time"
)

// FisherYatesKnuthShuffle Fisher–Yates-Knuth Shuffle或 算法对一维数组洗牌，O(n)
func FisherYatesKnuthShuffle[T any](slice []T) {
	rand.Seed(time.Now().Unix())
	for index := len(slice) - 1; index > 0; index-- {
		chooseIndex := rand.Intn(index + 1)
		slice[chooseIndex], slice[index] = slice[index], slice[chooseIndex]
	}
}

// FisherYatesShuffleMatrix Fisher–Yates-Knuth shuffle算法对矩阵洗牌
func FisherYatesShuffleMatrix[T any](matrix [][]T) {
	rand.Seed(time.Now().Unix())
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			chooseI := rand.Intn(i + 1)
			chooseJ := rand.Intn(j + 1)
			matrix[chooseI][chooseJ], matrix[i][j] = matrix[i][j], matrix[chooseI][chooseJ]
		}
	}
}

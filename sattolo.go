package shuffle

import (
	"math/rand"
	"time"
)

// SattoloShuffle Scatology算法
func SattoloShuffle[T any](slice []T) {
	rand.Seed(time.Now().Unix())
	for index := len(slice) - 1; index > 0; index-- {
		chooseIndex := rand.Intn(index)
		slice[chooseIndex], slice[index] = slice[index], slice[chooseIndex]
	}
}

func SattoloShuffleMatrix[T any](matrix [][]T) {
	rand.Seed(time.Now().Unix())
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[i]) - 1; j >= 0; j-- {
			chooseI := rand.Intn(i + 1)
			chooseJ := rand.Intn(j + 1)
			matrix[chooseI][chooseJ], matrix[i][j] = matrix[i][j], matrix[chooseI][chooseJ]
		}
	}
}

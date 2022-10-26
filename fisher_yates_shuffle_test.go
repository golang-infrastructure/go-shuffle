package shuffle

import "testing"

func TestFisherYatesShuffle(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	FisherYatesShuffle(slice)
	t.Log(slice)
}

func TestFisherYatesShuffleMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36, 37, 38, 39, 40},
	}
	FisherYatesShuffleMatrix(matrix)
	t.Log(matrix)
}

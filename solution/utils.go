package solution

import (
	"math/rand"
	"time"
)

func GenerateRandomSlice(length int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, length)
	for i := 0; i < length; i++ {
		slice[i] = rand.Intn(2)
	}
	return slice
}

func Flip(index int, data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	result[index] ^= 1
	return result
}

package main

import (
	"fmt"
	"math/rand"
)

func mergeSort(data []int) []int {
	length := len(data)
	if length <= 1 {
		return data
	}
	num := length / 2
	left := mergeSort(data[:num])
	right := mergeSort(data[num:])
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

func main() {
	s := make([]int, 0, 16)
	for i := 0; i < 16; i++ {
		s = append(s, rand.Intn(1000))
	}
	fmt.Println(s)
	s = mergeSort(s)
	fmt.Println(s)
}

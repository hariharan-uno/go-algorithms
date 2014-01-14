package main

import "fmt"

func main() {
	x := []int{2, 9, 4, 2, 5, 1, 56, 24, 6, 7, 1}
	fmt.Println(mergeSort(x))
}

func mergeSort(A []int) []int {
	if len(A) < 2 {
		return A
	}
	q := len(A) / 2
	c := mergeSort(A[:q])
	d := mergeSort(A[q:])
	return merge(c, d)
}

func merge(c []int, d []int) []int {
	b := make([]int, (len(c) + len(d)))
	i, j := 0, 0
	for i < len(c) && j < len(d) {
		if c[i] < d[j] {
			b[i+j] = c[i]
			i++
		} else {
			b[i+j] = d[j]
			j++
		}
	}
	for i < len(c) {
		b[i+j] = c[i]
		i++
	}
	for j < len(d) {
		b[i+j] = d[j]
		j++
	}
	return b
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var array []int64
	// open input file
	file, err := os.Open("IntegerArray.txt")
	if err != nil {
		panic(err)
	}
	// close file on exit and check for its returned error
	// make a read buffer
	r := bufio.NewReader(file)
	scanner := bufio.NewScanner(r)
	num := 0
	for scanner.Scan() {
		num, err = strconv.Atoi(scanner.Text())
		array = append(array, int64(num))
	}
	_, p := sortCount(array)
	fmt.Printf("%v\n", p)

}

func sortCount(A []int64) ([]int64, int64) {
	if len(A) == 1 {
		return A, 0
	} else {
		B, x := sortCount(A[:len(A)/2])
		C, y := sortCount(A[len(A)/2:])
		D, z := countSplitInv(B, C)
		return D, x + y + z
	}
	return A, 0
}

func countSplitInv(c []int64, d []int64) ([]int64, int64) {
	b := make([]int64, (len(c) + len(d)))
	i, j := 0, 0
	var inv int64
	for i < len(c) && j < len(d) {
		if c[i] < d[j] {
			b[i+j] = c[i]
			i++
		} else {
			b[i+j] = d[j]
			j++
			inv = inv + int64((len(c) - i))
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
	return b, inv
}

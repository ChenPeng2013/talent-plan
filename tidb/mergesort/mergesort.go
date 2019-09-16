package main

import "sync"

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
	if len(src) <= 1 {
		return
	}

	mid := len(src) / 2
	left := src[0:mid]
	right := src[mid:]

	if len(src) > 1000000 {
		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			MergeSort(left)
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			MergeSort(right)
			wg.Done()
		}()

		wg.Wait()
	} else {
		MergeSort(left)
		MergeSort(right)
	}

	a := []int64{}
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			a = append(a, left[i])
			i++
			continue
		}
		a = append(a, right[j])
		j++
	}

	if i < len(left) {
		a = append(a, left[i:]...)
	}
	if j < len(right) {
		 a = append(a, right[j:]...)
	}

	copy(src, a)
	return
}

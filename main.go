package main

import "fmt"

// CountInversions returns the number of pairs (i, j) with i < j and
// nums[i] > nums[j], counted during a merge sort. Input is not modified.
func CountInversions(nums []int) int64 {
	buf := append([]int{}, nums...)
	tmp := make([]int, len(buf))
	return sortCount(buf, tmp)
}

func sortCount(a, tmp []int) int64 {
	n := len(a)
	if n < 2 {
		return 0
	}
	mid := n / 2
	count := sortCount(a[:mid], tmp[:mid]) + sortCount(a[mid:], tmp[mid:])
	i, j, k := 0, mid, 0
	for i < mid && j < n {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			i++
		} else {
			// a[i..mid) are all greater than a[j]: that many inversions at once.
			count += int64(mid - i)
			tmp[k] = a[j]
			j++
		}
		k++
	}
	for i < mid {
		tmp[k] = a[i]
		i++
		k++
	}
	copy(a[:k], tmp[:k]) // any tail a[j:] is already in place
	return count
}

func main() {
	fmt.Println(CountInversions([]int{5, 3, 2, 4, 1}))
}

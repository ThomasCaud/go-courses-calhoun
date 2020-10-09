package module02

import "sort"

// Big O (without binary search): O(N^2), where N is the size of the list.
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.
func InsertionSortInt(list []int) {
	// list not a pointer, so we can't affect it to another slice
	// so we have to create another slice...
	var sorted []int

	// ... which is sorted...
	for _, val := range list {
		sorted = Insert(sorted, val)
	}

	// ... then we just affect sorted value in initial slice
	for i, val := range sorted {
		list[i] = val
	}
}

func Insert(list []int, newValue int) []int {
	for i, val := range list {
		if newValue < val {
			return append(list[:i], append([]int{newValue}, list[i:]...)...)
		}
	}
	return append(list, newValue)
}

// InsertionSort uses the standard library's sort.Interface to sort.
func InsertionSort(list sort.Interface) {
	for i := 0 ; i < list.Len() ; i++ {
		for j := 0 ; j < i ; j++ {
			if list.Less(i, j) {
				for k := j ; k < i ; k++ {
					list.Swap(k, i)
				}
			}
		}
	}
}

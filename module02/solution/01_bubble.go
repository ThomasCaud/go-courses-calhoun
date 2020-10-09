package module02

import "sort"

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	for sweepIndex, _ := range list {
		swapped := false
		for i := 0 ; i < len(list) - 1 - sweepIndex ; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			// optimisation: if elements are sorted, no need to go through the list again 
			break;
		}
	}
}

// BubbleSortString is a bubble sort for string slices.
func BubbleSortString(list []string) {
	for sweepIndex, _ := range list {
		swapped := false
		for i := 0 ; i < len(list) - 1 - sweepIndex ; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				swapped = true
			}
		}
		if !swapped {
			break;
		}
	}
}

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName.
func BubbleSortPerson(people []Person) {
	isSorted := func(a, b Person) bool {
		if a.Age != b.Age {
			return a.Age < b.Age
		}
		if a.LastName != b.LastName {
			return a.LastName < b.LastName
		}
		return a.FirstName < b.FirstName
	}

	for sweepIndex, _ := range people {
		swapped := false
		for i := 0 ; i < len(people) - 1 - sweepIndex ; i++ {
			if isSorted(people[i+1], people[i]) {
				people[i], people[i+1] = people[i+1], people[i]
				swapped = true
			}
		}
		if !swapped {
			break;
		}
	}
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
	for s := 0; s < list.Len(); s++ {
		swapped := false
		for i := 0; i < list.Len()-1-s; i++ {
			if list.Less(i+1, i) {
				list.Swap(i, i+1)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
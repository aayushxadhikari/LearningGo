package main

import "fmt"

// Define a struct to hold two slices for comparison
type CompareSlices struct {
	a, b []int
}

// Method to compare the slices
func (cs *CompareSlices) Compare() {
	if len(cs.a) != len(cs.b) {
		fmt.Println("A and B slices have different lengths")
		return
	}

	eq := true
	for i, valueA := range cs.a {
		if valueA != cs.b[i] {
			eq = false
			break
		}
	}

	if eq {
		fmt.Println("A and B slices are equal")
	} else {
		fmt.Println("A and B slices are not equal")
	}
}

// Defined a struct to manage slices and append/copy operations
type SliceManager struct {
	numbers []int
}

// Method to append elements to the numbers slice
func (sm *SliceManager) AppendElements(elements ...int) {
	sm.numbers = append(sm.numbers, elements...)
}

// Method to perform a copy operation between slices
func (sm *SliceManager) CopySlices(src, dst []int) (int, []int, []int) {
	n := copy(src, dst)
	return n, src, dst
}

func main() {
	// Comparing two slices
	cs := CompareSlices{
		a: []int{1, 2, 3},
		b: []int{1, 2, 3, 4},
	}
	cs.Compare() // Call the Compare method to compare the slices

	// Working with appending and copying
	sm := SliceManager{numbers: []int{2, 3}}

	// Append elements using the AppendElements method
	fmt.Println(sm.numbers)
	sm.AppendElements(10)
	fmt.Println(sm.numbers)
	sm.AppendElements(10, 20, 30)
	fmt.Println(sm.numbers)

	// Append another slice using the spread operator ...
	n := []int{100, 200}
	sm.AppendElements(n...)
	fmt.Println(sm.numbers)

	// Perform a copy operation using the CopySlices method
	src := []int{10, 20, 30}
	dst := []int{100, 200}
	nn, src, dst := sm.CopySlices(src, dst)
	fmt.Println(src, dst, nn)
}

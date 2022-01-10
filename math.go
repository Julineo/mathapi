package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// Helper function to calculate average
func Avg(arr []int) float64 {
	// Variable to store the sum
	var sum int

	// Calculate the sum
	for _, num := range arr {
		sum += num
	}

	// Calculate the average
	return float64(sum) / float64(len(arr))
}

// Helper function to calculate median
func Median(arr []int) float64 {
	/* To calculate median we sort the slice.
	Then check if the length is odd or even.
	If it is odd we return the middle element.
	If it is even we take average of both middle elements.
	*/
	sort.Ints(arr)
	mNumber := len(arr)/2
	if len(arr) % 2 != 0 {
		// Odd length. Return the middle element
		return float64(arr[mNumber])
	}
	return Avg(arr[mNumber - 1 : mNumber + 1])
}

// Custom function to calculate n min values
func Min(arr []int, quantifier int) ([]int, error) {
	/* Since task requires to get n min/max
	numbers, we will use custom functions to calculate it.
	We will use heap data structure to get the best performance.
	*/
	if quantifier > len(arr) || quantifier < 0 {
		err := fmt.Errorf("Input is outside of range.")
		return nil, err
	}

	h := &MinHeap{}
	heap.Init(h)
	for _, element := range arr {
		heap.Push(h, element)
	}
	result := make([]int, quantifier)
	for idx := 0; idx < quantifier; idx++ {
		result[idx] = heap.Pop(h).(int)
	}
	return result, nil
}

// Min-heap
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Custom function to calculate n max values
func Max(arr []int, quantifier int) ([]int, error) {
	/* Since task requires to get n min/max
	numbers, we will use custom functions to calculate it.
	We will use heap data structure to get the best performance.
	*/
	if quantifier > len(arr) || quantifier < 0 {
		err := fmt.Errorf("Input is outside of range.")
		return nil, err
	}

	h := &MaxHeap{}
	heap.Init(h)
	for _, element := range arr {
		heap.Push(h, element)
	}
	result := make([]int, quantifier)
	for idx := 0; idx < quantifier; idx++ {
		result[idx] = heap.Pop(h).(int)
	}
	return result, nil
}

// Max-heap
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
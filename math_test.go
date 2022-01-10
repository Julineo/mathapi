package main

import (
	"reflect"
	"testing"
)

// Testing the Avg function
func TestAvg(t *testing.T) {
	avg := Avg([]int{1,2,3,4})
	if avg != 2.5 {
		t.Errorf("Average is incorrect, got: %v, want: %v", avg, 2.5)
	}
}

// Testing the MedianFunction
func TestMedian(t *testing.T) {
	medianEven := Median([]int{1,2,3,4})
	if medianEven != 2.5 {
		t.Errorf("Median is incorrect, got: %v, want: %v", medianEven, 2.5)
	}
	medianOdd := Median([]int{1,2,3,4,5})
	if medianOdd != 3 {
		t.Errorf("Median is incorrect, got: %v, want: %v", medianOdd, 3)
	}
}

// Testing the Min function
func TestMin(t *testing.T) {
	arr := []int{10,1,4,10,8,3}
	res, err := Min(arr, 3)
	if !reflect.DeepEqual(res, []int{1,3,4}) {
		t.Errorf("Min is incorrect, got %v, want %v", res, []int{1,3,4})
	}
	if err != nil {
		t.Errorf("Min should not return error.")
	}
}

// Testing the Max function
func TestMax(t *testing.T) {
	arr := []int{10,1,4,10,8,3}
	res, err := Max(arr, 3)
	if !reflect.DeepEqual(res, []int{10,10,8}) {
		t.Errorf("Max is incorrect, got %v, want %v", res, []int{10,10,8})
	}
	if err != nil {
		t.Errorf("Max should not return error.")
	}
}

// Benchmarking the Avg function
func BenchmarkAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Avg([]int{1,2,3,4})
	}
}

// Benchmarking the Median function
func BenchmarkMedian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Median([]int{1,2,3,4,5})
	}
}

// Benchmarking the Min function
func BenchmarkMin(b *testing.B) {
	arr := []int{10,1,4,10,8,3}
	for i := 0; i < b.N; i++ {
		Min(arr, 3)
	}
}

// Benchmarking the Max function
func BenchmarkMax(b *testing.B) {
	arr := []int{10,1,4,10,8,3}
	for i := 0; i < b.N; i++ {
		Max(arr, 3)
	}
}
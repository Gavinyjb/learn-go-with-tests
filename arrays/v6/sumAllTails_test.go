package main

import (
	"fmt"
	"reflect"
	"testing"
)

//
//func TestSumAllTails(t *testing.T) {
//	got := SumAllTails([]int{1,2}, []int{0,9})
//	want := []int{2, 9}
//
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf("got %v want %v", got, want)
//	}
//}
func TestSumAllTails(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("测试一", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		//if !reflect.DeepEqual(got, want) {
		//	t.Errorf("got %v want %v", got, want)
		//}
		assertCorrectMessage(t, got, want)
	})

	t.Run("测试二", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 11})
		want := []int{0, 20}

		//if !reflect.DeepEqual(got, want) {
		//	t.Errorf("got %v want %v", got, want)
		//}
		assertCorrectMessage(t, got, want)
	})
}
func ExampleSumAllTails() {
	sums := SumAllTails([]int{1, 2}, []int{0, 9, 11})
	fmt.Println(sums)
	// Output: [2 20]
}

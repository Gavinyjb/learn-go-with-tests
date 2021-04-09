package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	//在 Go 中不能对切片使用等号运算符。你可以写一个函数迭代每个元素来检查它们的值。但是一种比较简单的办法是使用 reflect.DeepEqual，它在判断两个变量是否相等时十分有用

	//if got != want {
	//	t.Errorf("got %v want %v", got, want)
	//}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

//func BenchmarkSum(b *testing.B) {
//	numbers := []int{1, 2, 3}
//	for i := 0; i < b.N; i++ {
//		Sum(numbers)
//	}
//
//}

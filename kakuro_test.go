// Test of Sumtupel
//
// Ralf Poeppel 2023
package main

import (
	"reflect"
	"testing"
)

func Test_Sumtupel(t *testing.T) {
	mx := 9
	cases := []struct {
		mx   int
		size int
		sum  int
		want [][]int
	}{
		{mx, 2, 3, [][]int{[]int{1, 2}}},
		{mx, 2, 4, [][]int{[]int{1, 3}}},
		{mx, 2, 5, [][]int{[]int{1, 4}, []int{2, 3}}},
		{mx, 2, 6, [][]int{[]int{1, 5}, []int{2, 4}}},
		{mx, 2, 7, [][]int{[]int{1, 6}, []int{2, 5}, []int{3, 4}}},
		{mx, 2, 8, [][]int{[]int{1, 7}, []int{2, 6}, []int{3, 5}}},
		{mx, 2, 9, [][]int{[]int{1, 8}, []int{2, 7}, []int{3, 6}, []int{4, 5}}},
		{mx, 2, 10, [][]int{[]int{1, 9}, []int{2, 8}, []int{3, 7}, []int{4, 6}}},
		{mx, 2, 11, [][]int{[]int{2, 9}, []int{3, 8}, []int{4, 7}, []int{5, 6}}},
		{mx, 3, 6, [][]int{[]int{1, 2, 3}}},
		{mx, 3, 7, [][]int{[]int{1, 2, 4}}},
		{mx, 3, 8, [][]int{[]int{1, 2, 5}, []int{1, 3, 4}}},
		{mx, 3, 9, [][]int{[]int{1, 2, 6}, []int{1, 3, 5}, []int{2, 3, 4}}},
		{mx, 3, 10, [][]int{[]int{1, 2, 7}, []int{1, 3, 6}, []int{1, 4, 5},
			[]int{2, 3, 5}}},
		{mx, 3, 11, [][]int{[]int{1, 2, 8}, []int{1, 3, 7}, []int{1, 4, 6},
			[]int{2, 3, 6}, []int{2, 4, 5}}},
	}
	for _, c := range cases {
		tpl := Sumtupel(c.mx, c.size, c.sum)
		if !reflect.DeepEqual(tpl, c.want) {
			t.Errorf("sumtupel(%v, %v, %v) = \n%v want \n%v",
				c.mx, c.size, c.sum, tpl, c.want)
		}
	}
}

// Package kakuro computes the possible solutions of all kakuro lines
//
// Ralf Poeppel 2023
package main

import (
	"fmt"
)

const MAX_NUMBER = 9

// Sumtupel returns all tupels of size with given sum
// from values in range 1 to mx
//
// example
// mx = 9
// size = 3
// sumexp = 9
//
// pos = 0
// 1                     val = 1  rem=2*1+3,                             pos = 1,
//    2                  val = 2, rem=1*1+1,                             pos = 2,
//      3 -              val = 3, rem=0,     sum = 6  tup = 1,2,6        pos = 1
//    3                  val = 3  rem=1*3+1,                             pos = 2,
//      4 -              val = 4  rem=0,     sum = 8  tup = 1,3,5        pos = 1
//    4                  val = 4  rem=1*4+1,                             pos = 0
// 2                     val = 2  rem=2*2+3,                             pos = 1
//    3                  val = 3  rem=1*3+1,                             pos = 2
//      4 +              val = 4  rem=0,     sum = 9, tup = 2,3,4        pos = 1
//    4                  val = 4  rem=1*4+1,                             pos = 0
// 3                     val = 3  rem=2*3+3,                             pos = -1
//
func Sumtupel(mx, size, sumexp int) [][]int {
	// fmt.Println("SumTupel", size, sumexp)
	res := make([][]int, 0, 12) // result array of tupels
	tup := make([]int, size)    // tupel used for iteration
	// initialize tupel to smallest possibility
	for i := range tup {
		tup[i] = i
	}
	pos := 0        // index position in the tupel
	lst := size - 1 // last position in a tupel
	val := 0        // current value in iteration at position
	for pos >= 0 {
		val++
		tup[pos] = val
		sum := 0
		for i := 0; i <= pos; i++ {
			sum += tup[i]
		}
		// fmt.Println("> pos=", pos, "val", val, "tup=", tup[:pos+1], "sum=", sum)
		if pos != lst {
			n := (lst - pos)
			rem := n*tup[pos] + n*(n+1)/2 // minimal value of remaining
			// fmt.Println("  pos=", pos, "lst=", lst, "n=", n, "rem=", rem)
			if (rem + sum) > sumexp {
				// go back one position, tupel cannot be completed, will exceed given sum
				pos--
				if pos >= 0 {
					val = tup[pos]
				}
			} else {
				pos++
			}
		} else {
			// last position
			tup[pos] += sumexp - sum // compute last value
			if tup[pos] <= mx {
				// valid tupel
				tupClone := append([]int{}, tup...)
				res = append(res, tupClone)
				// fmt.Println("  res=", res)
			}
			// go back one position for next iteration
			pos--
			val = tup[pos]
		}
		// fmt.Println("  pos=", pos)
	}

	return res
}

// possibleNumbersForSum possible number combinations for sum
func possibleNumbersForSum(count, sum int) {
	fmt.Printf("%2d:", sum)
	combinations := Sumtupel(MAX_NUMBER, count, sum)
	for i := range combinations {
		fmt.Printf(" ")
		for j := range combinations[i] {
			if j != 0 {
				fmt.Printf(",")
			}
			fmt.Printf("%d", combinations[i][j])
		}
	}
	fmt.Println()
}

// possibilitiesWithCount with count numbers
func possibilitiesWithCount(count int) {
	min := count * (count + 1) / 2
	max := (MAX_NUMBER-count)*count + min
	fmt.Printf("\nsums in range %d to %d from %d possible numbers\n", min, max, count)

	for sum := min; sum <= max; sum++ {
		possibleNumbersForSum(count, sum)
	}
}

func main() {
	fmt.Println("List of possible Kakuro tupels")

	count := 9
	for i := 2; i <= count; i++ {
		possibilitiesWithCount(i)
	}
}

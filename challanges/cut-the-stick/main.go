package main

import "fmt"

func main() {
	arr := []int32{8, 8, 14, 10, 3, 5, 14, 12}

	fmt.Println(arr)
	fmt.Printf("cut: %v\n", cutTheSticks(arr))
}

// add the current length
// find the shortest
// decrement each by the shortest (including the shortest)
//
//	P1, P2 pointers
//
// 	P1 shortest number for the round (entire loop)
//  P2 cuts each by the shortest (entire loop)
// 	add length of final loop result
//  repeat until no number left

// NOTE: https://www.hackerrank.com/challenges/cut-the-sticks/problem?isFullScreen=true
func cutTheSticks(arr []int32) []int32 {
	result := []int32{}

	// NOTE: pointers
	p1, p2 := 0, 0

	// NOTE: min number for the round
	min := int32(0)

	result = append(result, int32(len(arr)))

	// NOTE: while there is a number
	for len(arr) > 0 {
		// NOTE: move P1 and find min
		if p1 <= len(arr)-1 {
			if min == 0 || min > arr[p1] {
				min = arr[p1]
			}

			p1++
		}

		// NOTE: when P1 cannot move, then move P2 and cut sticks
		if p1 > len(arr)-1 {

			// NOTE: when P2 can move
			if p2 <= len(arr)-1 {

				arr[p2] -= min

				if len(arr) > 0 && arr[p2] <= 0 {
					arr = append(arr[:p2], arr[p2+1:]...)
				} else {
					p2++
				}
			} else {
				result = append(result, int32(len(arr)))

				p1 = 0
				p2 = 0
				min = 0
			}
		}
	}

	return result
}

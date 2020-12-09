package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("go nuances!")
	mySlice := make([]int, 5)
	fmt.Println(mySlice)
	mySlice = append(mySlice, 10)
	fmt.Println(mySlice)

	// Slice Copy: copying mySlice to cpySlice, remember it has a reference to the mySlice, so will get modified along with mySlice
	cpySlice := make([]int, len(mySlice))
	cpySlice = mySlice
	fmt.Println(cpySlice)

	// Slice Cut: new cutSlice should be [1, 2, 3, 8, 9, 10], see the ... spread operator at the end as we can't just append a slice at the end
	cutSlice := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	cutSlice = append(cutSlice[:3], cutSlice[7:]...) // made a cut in indices: [3, 4, 5, 6]
	fmt.Println(cutSlice)

	// Slice Delete: deletes an element or consecutive elements in a range: similar to the cut, just use it accordingly.
	// To get the best efficiency, it is best to do the insertion without using append, in particular when the number of the inserted elements is known:
	copy(cutSlice, []int{100, 200})
	fmt.Println(cutSlice)

	// Pop Front/Shift: it's like a deque operation in a queue, we get the front value & remove the front value to update the queue
	var queue []int
	for i := 1; i <= 3; i++ {
		queue = append(queue, i)
	}
	x := -1
	x, queue = queue[0], queue[1:]
	fmt.Println("x: ", x, "queue: ", queue)

	// Reversing a Slice: it's pretty easy, remember it's not a good practice to calculate len(cutSlice) every time, calc. it once & store it.
	for i := 0; i < len(cutSlice) / 2; i++ {
		cutSlice[i], cutSlice[len(cutSlice) - i - 1] = cutSlice[len(cutSlice) - i - 1], cutSlice[i]
	}
	fmt.Println(cutSlice)
	// same thing, reversing the slice, except with two indices.
	for left, right := 0, len(cutSlice)-1; left < right; left, right = left+1, right-1 {
		cutSlice[left], cutSlice[right] = cutSlice[right], cutSlice[left]
	}
	fmt.Println(cutSlice)
	// Shuffling slice or array, used rand.Intn() from math/rand library
	for i := len(cutSlice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		cutSlice[i], cutSlice[j] = cutSlice[j], cutSlice[i]
	}
	fmt.Println(cutSlice)
}
package main

import (
	"fmt"
)

func changeVal2(myPtr *int) {
	*myPtr = 12
}

func dblArrVal(arr []int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, val := range nums {
		sum += val
	}
	return sum / numSize
}

func main() {
	pl := fmt.Println
	f4 := 10

	pl("f4: ", f4)
	pl("Address of f4: ", &f4)

	f4Ptr := &f4

	pl("f4Ptr: ", f4Ptr)
	pl("Value at f4Ptr: ", *f4Ptr)

	*f4Ptr = 20

	pl("f4Ptr after changing value at f4Ptr: ", *f4Ptr)
	pl("f4Ptr Address: ", &f4Ptr)

	pl("f4 before function: ", f4)
	changeVal2(&f4)
	pl("f4 after function: ", f4)

	pArr := []int{1, 2, 3, 4}
	dblArrVal(pArr)
	pl(pArr)

	iSlice := []float64{11, 12, 13, 14, 15}
	fmt.Printf("Average of iSlice: %.3f\n", getAverage(iSlice...))
}

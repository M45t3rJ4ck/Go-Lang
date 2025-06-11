package main

import (
	"fmt"
)

func main() {
	slc1 := make([]string, 6)
	slc1[0] = "Society"
	slc1[1] = "of"
	slc1[2] = "the"
	slc1[3] = "simulated"
	slc1[4] = "universe"

	fmt.Println("Slice Size: ", len(slc1))

	for i := 0; i < len(slc1); i++ {
		fmt.Println("Index", i, "of slc1: ", slc1[i])
	}

	slc2 := []int{27, 05, 1983}
	fmt.Println("Slice slc2: ", slc2)

	sArr := [5]int{1, 2, 3, 4, 5}
	slc3 := sArr[0:4]
	fmt.Println("First till forth index of slice 3: ", slc3)
	slc4 := sArr[:3]
	fmt.Println("First 3 indexes of slice 4: ", slc4)
	slc5 := sArr[2:]
	fmt.Println("Last 3 indexes of slice 5: ", slc5)

	sArr[0] = 10
	fmt.Println("Slice 3: ", slc3)

	slc3 = append(slc3, 12)
	fmt.Println("Slice 3 after appending 12: ", slc3)

	slc6 := make([]string, 6)
	fmt.Println("Slice 4: ", slc6)
}

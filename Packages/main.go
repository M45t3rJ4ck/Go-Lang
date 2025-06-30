package main

import (
	stuff "example/project/my_packages"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello", stuff.Name)

	intArr := []int{1, 2, 3, 4, 5}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println("strArr Typeof: ", reflect.TypeOf(strArr))
	fmt.Println("IntArr Typeof: ", reflect.TypeOf(intArr))
}

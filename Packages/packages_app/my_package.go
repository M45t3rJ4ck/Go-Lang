package stuff

import "strconv"

// Capitalized variable names can be accessed from outside this file.
var Name string = "Riaan"

// Capitalized function names can be accessed from outside this file.
func IntArrToStrArr(intArr []int) []string {
	var strArr []string

	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}

	return strArr
}

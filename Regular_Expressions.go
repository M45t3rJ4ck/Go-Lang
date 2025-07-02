package main

import (
	"fmt"
	"regexp"
)

func main() {
	pl := fmt.Println

	reStr1 := "The ape was at the apex."
	match, _ := regexp.MatchString("(ape[^ ]?)", reStr1)
	pl("Match String:", match)

	reStr2 := "Cat rat mat fat pat sat"
	// Case sensitive - lower case c was searched for not upper case c.
	compiler, _ := regexp.Compile("([crmfps]at)")
	pl("Complier:", compiler.MatchString(reStr2))

	pl("Find String:", compiler.FindString(reStr2))
	pl("Find String Index:", compiler.FindStringIndex(reStr2))
	pl("Find All Matching Strings:", compiler.FindAllString(reStr2, -1))
	pl("Find First 2 Matching Strings:", compiler.FindAllString(reStr2, 2))
	pl("Find All Index:", compiler.FindAllStringSubmatchIndex(reStr2, -1))
	pl("Replace All Strings:", compiler.ReplaceAllString(reStr2, "Zap"))
}

package utils

import "fmt"

func CharacterLimitFix(stringLength int) {
	data := "Imagine this is more than 120 Characters"

	if len(data) > stringLength {
		fmt.Println(data[:stringLength] + " ...")
	} else {
		fmt.Println(data)
	}
}

/*
NOTE:
If character length is 25 then output is :
Imagine this is more than ...

*/

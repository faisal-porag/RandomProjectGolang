package StringFormatter

import "fmt"

func StringSliceExample() {
	numbers := make([]int, 3, 5)
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

	//This line will cause a runtime error index out of range [4] with length 3
	//numbers[4] = 5

	//Increasing the length from 3 to 5
	numbers = numbers[0:5]
	fmt.Println("\nIncreasing length from 3 to 5")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))

	//Decresing the length from 3 to 2
	numbers = numbers[0:2]
	fmt.Println("\nDecreasing length from 3 to 2")
	fmt.Printf("numbers=%v\n", numbers)
	fmt.Printf("length=%d\n", len(numbers))
	fmt.Printf("capacity=%d\n", cap(numbers))
}

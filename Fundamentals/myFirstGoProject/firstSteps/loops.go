package firstSteps

import "fmt"

func Loops() {
	//Initial statement -> Conditional Expression -> Post-Statement
	//Any of above, is optionally. You will understand
	var output int
	//var i int
	//for i < 10 ...
	for i := 0; i < 10; i++ {
		output++
	}
	fmt.Println(output)
}

func Range() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Blank Identifier -> underline. sometimes we dont want the 'i' return.
	for _, elements := range arr {
		fmt.Println(elements)
	}
}

func NewLoopUpdated() {
	for i := range 10 {
		fmt.Println(i)
	}
}

func ArrayPointer() {
	arr := [3]int{1, 2, 3}
	for i, elem := range arr {
		fmt.Println(&i, &elem)
	}
}

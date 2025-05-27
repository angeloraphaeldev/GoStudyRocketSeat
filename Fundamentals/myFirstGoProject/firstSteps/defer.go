package firstSteps

import "fmt"

// Put some function into await while the main function is executed.
func Dodefer() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}

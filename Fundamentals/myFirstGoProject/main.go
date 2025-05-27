package main

import (
	"fmt"
	fS "myFirstGoProject/firstSteps"
	"myFirstGoProject/packageTest"
	//"myFirstGoProject/packageTest/internal/foo" -------- Internal Package can't be imported on main file.
)

func main() {
	fmt.Println("Hello World!")
	packageTest.Nick()
	packageTest.Print()
	fmt.Println(packageTest.Name)

	//foo.Foo
	fmt.Println("------------------------------ FUNCTION --------------------------------")
	fmt.Println(fS.SumMath(1, 2))
	c, d := fS.Swap(10, 20)
	fmt.Println(c, d)

	res, rem := fS.DivisionMath(10, 3)
	fmt.Println(res, rem)

	fS.Arrays()

	fS.Loops()

	fS.Range()

	fS.Loops()

	fS.NewLoopUpdated()

	fS.ArrayPointer() // Memory Allocation of Pointer.

	fS.Condition()

	fS.Dodefer() // Defer test

}

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
	fmt.Println(fS.Somar(1, 2))
	c, d := fS.Swap(10, 20)
	fmt.Println(c, d)

	res, rem := fS.Dividir(10, 3)
	fmt.Println(res, rem)

}

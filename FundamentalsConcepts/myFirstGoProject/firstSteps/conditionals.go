package firstSteps

import (
	"fmt"
	"math"
)

func Condition() {
	/* 	if 1 < 2 {
		fmt.Println("Is Higher")
	}
	*/
	if x := math.Sqrt(4); x < 1 {
		fmt.Println(x)
	} else if x > 0 {
		fmt.Println("Higher than zero")
	}

}

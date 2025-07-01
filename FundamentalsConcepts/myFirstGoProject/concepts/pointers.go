package concepts

import (
	"fmt"
)

/*
Memory Stack -> called stack of a thread. Last item of stack will be the first. LIFO
Memory Heap -> Dynamic when we need more memory on our programs. Sometimes the maps type of data is used here, sometimes in Stack if is too small.
*/

func Pointers() {
	x := 10
	take(x)
	fmt.Println(x)
}

func take(x int) {
	x = 100

}

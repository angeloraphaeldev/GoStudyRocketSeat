package firstSteps

import "fmt"

func Arrays() {
	// var a = [4]int{1,2,3,4} -> Should have the array size (in time of compilation)
	// var b = [...]int{1,4,2,1,5,6,7,3,2,1,2,3,4,1,2,3}
	// var c = [10]int{5: 400} specify the value of index
	// arrays like that is static value, dynamic values is called slices

	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	arr1 := [12]string{10: "Hello World, Im Back"}
	fmt.Println(arr1)
	fmt.Println(arr[2] + arr[3])
	fmt.Println(arr1[10] + arr1[3])
	arr2 := [2]string{"testing", " xcy"}
	fmt.Println(arr1[10] + arr2[1])
}

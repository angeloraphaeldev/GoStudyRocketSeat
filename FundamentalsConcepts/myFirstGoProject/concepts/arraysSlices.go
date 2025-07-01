package concepts

import "fmt"

func ArraySlices() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9} //Limited
	fmt.Println(arr)
	slic := []int{1, 3, 2, 5, 8, 0} //Unlimited
	fmt.Println(slic)
	slic = append(slic, 0, 3, 2, 1, 4, 3)
	fmt.Println(slic)

}

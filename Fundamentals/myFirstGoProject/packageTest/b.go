package packageTest

import (
	"fmt"
	"myFirstGoProject/packageTest/internal/foo"
)

var Name = "Raph"
var nickName = "None"

func Nick() {
	fmt.Println(Name + " " + nickName)
	fmt.Println(foo.Foo)
}

package firstSteps

var (
	num1 int
	num2 int
)

// There's some ways to declare/create an function in golang. I'll show some of them below
func SumMath(a int, b int) int {
	return a + b
}

func Swap(c, d int) (int, int) {
	return d, c
}

func DivisionMath(e, f int) (int, int) {
	res := e / f
	rem := e % f
	return res, rem
}

package mathutils

func Factorial(x int) (r int) {

	r = 1

	for i := 1; i <= x; i++ {
		r *= i
	}

	return r
}

// fibonacci calculates the n'th element of the famous sequence.
func fibonacci(n int) int {
	/*
		[]test{
			{1, 1},
			{2, 1},
			{3, 2},
			{7, 13},
			{11, 89},
		}
	*/
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

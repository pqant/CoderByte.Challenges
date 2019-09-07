package TryMe

func FibGenerator(val int) <-chan int {
	result := make(chan int)
	go func() {
		if val < 2 {
			result <- val
			return
		}
		result <- <-FibGenerator(val - 1) + <-FibGenerator(val - 2)
	}()
	return result
}

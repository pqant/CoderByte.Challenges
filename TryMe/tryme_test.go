package TryMe

import (
	"testing"
)


func BenchmarkFibGenerator(b *testing.B) {
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		<-FibGenerator(10)
		/*
			go func(i int) {
				log.Printf("result %v for %v \n", <-FibGenerator(i), i)
			}(i)

		*/
	}
}

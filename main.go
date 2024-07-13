package main

import (
	"fmt"
	"sync"
)

func square(num int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- num * num
}
func main() {
	numbers := []int{2, 4, 6, 8, 10}

	ch := make(chan int, len(numbers))

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go square(num, &wg, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	sum := 0
	for result := range ch {
		sum += result
	}
	fmt.Println("Сумма квадратов:", sum)
}

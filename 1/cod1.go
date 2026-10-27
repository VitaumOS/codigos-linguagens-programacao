package main
import ("fmt";"time")

func main() {
	ch := make(chan string)
	for i := 1; i <= 3; i++ {
		go func(n int) {
			time.Sleep(time.Duration(4-n) * 100 * time.Millisecond)
			ch <- fmt.Sprint("tarefa ", n)
		}(i)
	}
	fmt.Println(<-ch, <-ch, <-ch)
}
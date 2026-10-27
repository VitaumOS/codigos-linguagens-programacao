package main
import ("fmt";"time")

func main() {
    c := make(chan string,1)
    go func() { time.Sleep(2 * time.Second); c <- "pronto" }()
    select {
    case msg := <-c:
        fmt.Println(msg)
    case <-time.After(time.Second):
        fmt.Println("timeout")
    }
}

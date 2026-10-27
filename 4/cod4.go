package main
import "fmt"

func raiz(x int) (int, error) {
    for i := 1; i*i <= x; i++ {
        if i*i == x { return i, nil }
    }
    return 0, fmt.Errorf("%d não tem raiz inteira positiva", x)
}
func main() {
    if r, err := raiz(16); err == nil { fmt.Println(r) }
    if _, err := raiz(10); err != nil { fmt.Println(err) }
}
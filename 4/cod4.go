package main
import "fmt"

func divmod(a, b int) (q, r int) {
	q, r = a/b, a%b
	return
}
func main() {
	q, r := divmod(17, 5)
	_, r2 := divmod(9, 4)
	fmt.Println(q, r, r2)
}
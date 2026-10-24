package main
import "fmt"

type Number interface{ ~int | ~float64 }

func Soma[T Number](xs []T) T {
	var s T
	for _, x := range xs {
		s += x
	}
	return s
}
func main() {fmt.Println(Soma([]int{1, 2, 3}), Soma([]float64{1.5, 2.5}))}
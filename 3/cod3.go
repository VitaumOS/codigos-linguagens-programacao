package main
import "fmt"

func dividir(a, b int) (r int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("recuperado: %v", e)
		}
	}()
	return a / b, nil
}
func main() {fmt.Println(dividir(10, 0))}

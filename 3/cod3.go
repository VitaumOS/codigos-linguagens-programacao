package main
import "fmt"

func dividir(a, b int) (_ int, err error) {
    defer fmt.Println("saindo")
    defer func() {
        if e := recover(); e != nil { 
        err = fmt.Errorf("recuperado: %v", e) }
    }()
    if b == 0 { panic("divisão por zero") }
    return a / b, nil
}
func main() { fmt.Println(dividir(10, 0)) }
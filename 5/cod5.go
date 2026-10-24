package main
import "fmt"

type Falante interface{ Falar() string }
type Cao struct{}
type Gato struct{}

func (Cao) Falar() string  { return "Au" }
func (Gato) Falar() string { return "Miau" }

func main() {
	for _, f := range []Falante{Cao{}, Gato{}} {fmt.Println(f.Falar())}
}
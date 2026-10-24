package main
import "fmt"

type Animal struct{ Nome string }

func (a Animal) Apresentar() string { return "Sou " + a.Nome }

type Cachorro struct {
	Animal
	Raca string
}
func main() {
	c := Cachorro{Animal{"Rex"}, "Labrador"}
	fmt.Println(c.Apresentar(), c.Nome)
}

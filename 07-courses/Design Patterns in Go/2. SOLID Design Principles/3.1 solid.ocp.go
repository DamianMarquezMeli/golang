package main

import "fmt"

// combination of OCP and Repository demo

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name string
	color Color
	size Size
}

type Filter struct {

}

func (f *Filter) filterByColor(
	products []Product, color Color)[]*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) filterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) filterBySizeAndColor(
	products []Product, size Size,
	color Color)[]*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// filterBySize, filterBySizeAndColor

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

type SizeSpecification struct {
	size Size
}

func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(p *Product) bool {
	return spec.first.IsSatisfied(p) &&
		spec.second.IsSatisfied(p)
}

type BetterFilter struct {}

func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main_() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{ "House", blue, large}

	products := []Product{apple, tree, house}

	fmt.Print("Green products (old):\n")
	f := Filter{}
	for _, v := range f.filterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}
	// ^^^ BEFORE

	// vvv AFTER
	fmt.Print("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}

	largeGreenSpec := AndSpecification{largeSpec, greenSpec}
	fmt.Print("Large blue items:\n")
	for _, v := range bf.Filter(products, largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.name)
	}
}


////////////////////////////////////////
// es mejor usar shape poliformicamente que tener un metodo are que hay que modificar

package main

import (
	"fmt"
)

// Interface
// las interfaces SI puede recibir parametros de entrada
type forma interface {
	area() float32
	perimetro() float32
}

type rectangulo struct {
	largo float32
	ancho float32
}

type cuadrado struct {
	lado float32
}

type triangulo struct {
	base   float32
	altura float32
}

//////////////////////
// Medodos

// metodo area rectangulo
func (r rectangulo) area( /*parametros de entreda*/ ) float32 {
	return r.largo * r.ancho
}

// Si el comento, o elimino el metodo perimetro de rectangulo,
// rectangulo ya no sera de tipo forma
func (r rectangulo) perimetro() float32 {
	return r.largo * r.ancho * 2
}

// Aunque este metodo NO esta dentro del methodset de forma,
// si hay polimofismo
func (r rectangulo) rotar() int {
	return 30
}

// metodo area cuadrado
func (c cuadrado) area() float32 {
	return c.lado * 2
}

func (c cuadrado) perimetro() float32 {
	return c.lado * 4
}

// triángulo isósceles tiene dos lados idénticos
func (t triangulo) area() float32 {
	return (t.base * t.altura) / 2
}

// el perimetro de un triangulo tiene mas complejidad, pero fines de demostracion
// diremos que es base x altura x 3
func (t triangulo) perimetro() float32 {
	return (t.base * t.altura) * 3
}

// no es un metodo, es un funcion, no tiene reciber.
// aqui es donde se puede ver el uso del polimorfismo
// como el methodset de la interface forma
// esta incluido dentro de los metodos implmentados en triangulo, rectangulo y cuadrodo
// estos TAMBIEN son de TIPO forma, por lo que le podemos enviar
// o sea, triangulo, cuadrado o rectangulo, puede tener mas metodos y no esta en la interface,
// pero todos los metodos en la interface DEBEN estar implementados en los tipos para
// que se cumpla el polimorfismo.
func imprimirArea(f forma) string {
	return fmt.Sprintf("Area: %f - Perimetro: %f", f.area(), f.perimetro())
}

func main() {

	rec := rectangulo{
		largo: 10.0,
		ancho: 50.0,
	}

	cua := cuadrado{
		lado: 40,
	}

	tri := triangulo{
		base:   20.0,
		altura: 30.0,
	}

	fmt.Println("Area rectangulo: ", imprimirArea(rec))
	fmt.Println("Area cuadrado: ", imprimirArea(cua))
	fmt.Println("Area triangulo:", imprimirArea(tri))
}




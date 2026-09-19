// Structs, methods aur interfaces. Go me classes nahi hoti, ye teeno milke kaam karte hain.
package main

import (
	"fmt"
	"math"
)

// Shape interface: jis type me Area() hoga wo automatically Shape ban jayega
// (implicit implementation, "implements" likhne ki zaroorat nahi)
type Shape interface {
	Area() float64
}

type Rect struct{ W, H float64 }
type Circle struct{ R float64 }

func (r Rect) Area() float64   { return r.W * r.H }
func (c Circle) Area() float64 { return math.Pi * c.R * c.R }

type Counter struct{ n int }

// pointer receiver (*Counter): original struct badalta hai.
// value receiver hota to sirf copy badalti.
func (c *Counter) Inc() { c.n++ }

func main() {
	shapes := []Shape{Rect{2, 3}, Circle{1}}
	for _, s := range shapes {
		fmt.Printf("%T area = %.2f\n", s, s.Area())
	}

	c := Counter{}
	c.Inc()
	c.Inc()
	fmt.Println("counter:", c.n)

	// type switch: interface ke andar actual type pata karne ke liye
	var x any = Rect{1, 1}
	switch v := x.(type) {
	case Rect:
		fmt.Println("rect hai", v.W)
	case Circle:
		fmt.Println("circle hai")
	}
}

package main

// Bridge
// Connecting components together through abstractions

// avoid the complexity of crossmultiple
// imagine you have 5 types of figures and 2 ways of drawing
// you would have to create 10 different method to do that
// with bridge we have an interface that handles that

import "fmt"

// the interface that allows us to specify the type that we want
type Renderer interface {
	RenderCircle(radius float32)
}

// one type of renderer
type VectorRenderer struct {
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

// other type of renderer
type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

// has a renderer, and you can define it as you want
type Circle struct {
	renderer Renderer
	radius   float32
}

// use the abstraction
func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

// receive the renderer
func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func main() {
	raster := RasterRenderer{}
	vector := VectorRenderer{}

	circle := NewCircle(&vector, 5)
	circle.Draw()

	circle = NewCircle(&raster, 10)
	circle.Draw()
}

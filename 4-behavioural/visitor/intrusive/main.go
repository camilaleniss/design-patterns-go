package visitor

import (
	"fmt"
	"strings"
)

type Expression interface {
	Print(sb *strings.Builder)
}

// let's call this number... or something
type DoubleExpression struct {
	value float64
}

// just prints the value
func (d *DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", d.value))
}

type AdditionExpression struct {
	left, right Expression
}

// it's intrusive cause you pass the sb over the hierarchy
// violating the open-closed principle
func (a *AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteString("(")
	// prints an expression
	a.left.Print(sb)
	sb.WriteString("+")
	// prints an expression
	a.right.Print(sb)
	sb.WriteString(")")
}

func main() {
	// 1+(2+3)
	e := AdditionExpression{
		&DoubleExpression{1},
		&AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}

	sb := strings.Builder{}
	e.Print(&sb)

	fmt.Println(sb.String())
}

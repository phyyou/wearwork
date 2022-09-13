package object

import "fmt"

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

const (
	DOUBLE_OBJ = "DOUBLE"
	NULL_OBJ   = "NULL"
)

type Double struct {
	Value float64
}

func (i *Double) Type() ObjectType { return DOUBLE_OBJ }
func (i *Double) Inspect() string  { return fmt.Sprintf("%f", i.Value) }

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

package object

import "fmt"

type ObjectType string

const (
	BOOLEAN_OBJ      = "BOOLEAN"
	ERROR_OBJ        = "ERROR"
	INTEGER_OBJ      = "INTEGER"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

type Null struct{}

func (n *Null) Inspect() string {
	return "null"
}
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}
func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

type Error struct {
	Message string
}

func (e *Error) Inspect() string {
	return "ERROR:" + e.Message
}
func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}

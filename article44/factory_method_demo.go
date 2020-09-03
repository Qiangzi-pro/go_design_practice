package main

import (
	"fmt"
)

type Operation struct {
	a float64
	b float64
}

type OperationI interface {
	GetResult() float64
	SetA(float64)
	SetB(float64)
}

func (op *Operation) SetA(a float64) {
	op.a = a
}

func (op *Operation) SetB(b float64) {
	op.b = b
}

type AddOperation struct {
	Operation
}

func (this *AddOperation) GetResult() float64 {
	return this.a + this.b
}

type SubOperation struct {
	Operation
}

func (this *SubOperation) GetResult() float64 {
	return this.a - this.b
}

//type IFactory interface {
//	CreateOperation() Operation
//}

type AddFactory struct {
}

func (this *AddFactory) CreateOperation() OperationI {
	return &(AddOperation{})
}

type SubFactory struct {
}

func (this *SubFactory) CreateOperation() OperationI {
	return &(SubOperation{})
}

func main() {
	fac := &(AddFactory{})
	oper := fac.CreateOperation()
	fmt.Printf("%v, %T\n", oper, oper)
	oper.SetA(1)
	oper.SetB(5)
	fmt.Println(oper.GetResult())
}

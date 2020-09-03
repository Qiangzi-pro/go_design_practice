package main

import (
	"fmt"
)

type Operator interface {
	Operate(int, int) int
}

type AddOperate struct{}

func (ao *AddOperate) Operate(rhs int, lhs int) int {
	return rhs + lhs
}

type MultipleOperate struct{}

func (mo *MultipleOperate) Operate(rhs int, lhs int) int {
	return rhs * lhs
}

type OperateFactory struct{}

func NewOperateFactory() *OperateFactory {
	return &OperateFactory{}
}

func (of *OperateFactory) CreateOperate(operatename string) Operator {
	switch operatename {
	case "+":
		return &AddOperate{}
	case "*":
		return &MultipleOperate{}
	default:
		panic("无效运算符号")
		return nil
	}
}

var cachedOperator map[string]Operator

func init() {
	cachedOperator = make(map[string]Operator)
	cachedOperator["+"] = &AddOperate{}
	cachedOperator["*"] = &MultipleOperate{}
}

// 简单工厂的第二种实现
func (of *OperateFactory) createOperate2(name string) Operator {
	if len(name) == 0 {
		return nil
	}

	var operator Operator = cachedOperator[name]
	return operator

}

func main() {
	//Operator := NewOperateFactory().CreateOperate("+")
	Operator := NewOperateFactory().createOperate2("+")

	fmt.Printf("add result is %d\n", Operator.Operate(1, 2))
}

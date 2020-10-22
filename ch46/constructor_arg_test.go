package ch46

import "testing"

type ConstructorArg struct {
	isRef bool
	kind  interface{}
	arg   interface{}
}

type ConstructorArgBuilder struct {
	isRef bool
	kind  interface{}
	arg   interface{}
}

func (builder *ConstructorArgBuilder) build() *ConstructorArg {
	if builder.isRef == true {
		if _, ok := builder.arg.(string); !ok {
			panic("error arg")
		}
	} else {
		if builder.kind == nil || builder.arg == nil {
			panic("kind and arg need set nue")
		}
	}

	return &ConstructorArg{
		isRef: builder.isRef,
		kind:  builder.kind,
		arg:   builder.arg,
	}
}

func (builder *ConstructorArgBuilder) setIsRef(isRef bool) *ConstructorArgBuilder {
	builder.isRef = isRef
	return builder
}

func (builder *ConstructorArgBuilder) setKind(kind interface{}) *ConstructorArgBuilder {
	builder.kind = kind
	return builder
}

func (builder *ConstructorArgBuilder) setArg(arg interface{}) *ConstructorArgBuilder {
	builder.arg = arg
	return builder
}

func TestInterfaceFuc(t *testing.T) {
	s := "today is friday"

	var i interface{} = s
	t.Log(i)

}

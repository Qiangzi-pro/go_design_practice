package article68_visitor_desgin_pattern

import (
	"fmt"
)

type ResourceFile struct {
	FilePath string
}

type PdfFile struct {
	*ResourceFile
}

type PPTFile struct {
	*ResourceFile
}

type WordFile struct {
	*ResourceFile
}

type Visitor interface {
	Visitor(f interface{})
}

func (r *ResourceFile) accept(visitor Visitor) {
	visitor.Visitor(r)
}

type Extractor struct {
	SupportObjs []interface{}
}

// 待进一步调研
func (e Extractor) Visitor(f interface{}) {
	//for _, obj := range e.SupportObjs {
	//	val, ok := f.(obj)
	//}
	//switch r := f.(type) {
	//case string:
	//	fmt.Println(r)
	//}
}




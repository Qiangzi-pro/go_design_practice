package version2

import (
	"testing"
)

//type Employee struct {
//	salary int
//}

//func EmployeeByID(id int) Employee  {
//	return Employee{11}
//}

func TestFile(t *testing.T) {
	//dirsTmp, _ := ioutil.ReadDir("/Users/yuqiang/go_workspace/go_learning/src/Design_Pattern")
	//for _, f := range dirsTmp {
	//	t.Log(f.Name())
	//}

	//files, _ := filepath.Glob("*")
	//t.Log(files)
	//pp := new(Employee)
	//
	//*pp = Employee{salary: 10}

	//EmployeeByID(2).salary = 1

	//var a [3]int
	//t.Log(len(a))

	s := "/Users/yuqiang/go_workspace/go_learning/src/Design_Pattern/"
	t.Log(len(s), s[len(s)-1] == '/')

	d := NewDirectory(s)
	t.Log(d.FileSize())
	t.Log(d.FileNum())

	t.Logf("%#v", d)

	path2 := "/Users/yuqiang/go_workspace/go_learning/src/Design_Pattern/dependency_injection"
	d2 := NewDirectory(path2)

	t.Log(d2.FileSize())
	t.Log(d2.FileNum())

	d.RemoveSubNode(d2)

	t.Log(d.FileSize())
	t.Log(d.FileNum())

}

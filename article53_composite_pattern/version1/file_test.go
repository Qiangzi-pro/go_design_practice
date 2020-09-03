package version1

import (
	"testing"
)

func TestFile(t *testing.T) {
	//dirsTmp, _ := ioutil.ReadDir("/Users/yuqiang/go_workspace/go_learning/src/Design_Pattern")
	//for _, f := range dirsTmp {
	//	t.Log(f.Name())
	//}

	//files, _ := filepath.Glob("*")
	//t.Log(files)

	s := "/Users/yuqiang/go_workspace/go_learning/src/Design_Pattern/dependency_injection"
	t.Log(len(s), s[len(s)-1] == '/')

	d := NewDirectory(s)
	t.Log(d.CountSizeOfFiles())
	t.Log(d.CountNumOfFiles())
}

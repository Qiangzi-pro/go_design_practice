package version2

import (
	"io/ioutil"
	"os"
	"strings"
)

const (
	Separator = string(os.PathSeparator)
)

type IFileSystem interface {
	GetPath() string
	Parent() IFileSystem
	SetParent(node IFileSystem)
	FileSize() int64
	SetFileSize(fileSize int64)
	FileNum() int
}

type FileSystemNode struct {
	path     string
	fileSize int64
	parent   IFileSystem
}

func (fsn *FileSystemNode) SetParent(node IFileSystem) {
	fsn.parent = node
}

func (fsn *FileSystemNode) Parent() IFileSystem {
	return fsn.parent
}

func (fsn *FileSystemNode) FileSize() int64 {
	return fsn.fileSize
}

func (fsn *FileSystemNode) SetFileSize(fileSize int64) {
	fsn.fileSize = fileSize
}

func (fsn *FileSystemNode) GetPath() string {
	return fsn.path
}

func (fsn *FileSystemNode) FileNum() int {
	return 1
}

type File struct {
	FileSystemNode
}

func NewFile(path string) *File {
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		panic(err)
	}
	return &File{
		FileSystemNode: FileSystemNode{path: path, fileSize: fileInfo.Size()},
	}
}

type IDirSystem interface {
	IFileSystem
	AddSubNode(node IFileSystem)
	RemoveSubNode(node IFileSystem)
	SetFileNum(fileNum int)
}

type Directory struct {
	FileSystemNode
	subNodes []IFileSystem
	fileNum  int
}

func (dir *Directory) FileNum() int {
	return dir.fileNum
}

func (dir *Directory) SetFileNum(fileNum int) {
	dir.fileNum = fileNum
}

func NewDirectory(path string) *Directory {
	//fileInfo, err := os.Stat(path)
	//if os.IsNotExist(err) || !fileInfo.IsDir() {
	//	errors.New("目录路径错误，请检查")
	//}
	if path[len(path)-1] != os.PathSeparator {
		path += Separator
	}

	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	root := &Directory{
		FileSystemNode: FileSystemNode{path: path},
		subNodes:       make([]IFileSystem, 0),
	}

	for _, fileInfo := range dirs {
		var node IFileSystem
		currentPath := path + fileInfo.Name()
		if fileInfo.IsDir() {
			node = NewDirectory(currentPath)
		} else {
			node = NewFile(currentPath)
		}
		node.SetParent(root)
		root.AddSubNode(node)
	}

	return root
}

func (dir *Directory) adjustSizeAndNum(node IFileSystem, isAdd bool) {
	switch isAdd {
	case true:
		dir.SetFileSize(dir.fileSize + node.FileSize())
		dir.SetFileNum(dir.fileNum + node.FileNum())
	case false:
		dir.SetFileSize(dir.fileSize - node.FileSize())
		dir.SetFileNum(dir.fileNum - node.FileNum())
	}

	if pNode, ok := dir.Parent().(*Directory); ok {
		pNode.adjustSizeAndNum(node, isAdd)
	}
}

func (dir *Directory) AddSubNode(node IFileSystem) {
	dir.subNodes = append(dir.subNodes, node)

	dir.adjustSizeAndNum(node, true)
}

func (dir *Directory) RemoveSubNode(node IFileSystem) {
	i := 0
	for ; i < len(dir.subNodes); i++ {
		t, t2 := dir.subNodes[i].GetPath(), node.GetPath()
		if strings.EqualFold(t, t2) {
			break
		}
	}

	if i < len(dir.subNodes) {
		// 切片删除操作
		dir.subNodes = append(dir.subNodes[:i], dir.subNodes[i+1:]...)

		dir.adjustSizeAndNum(node, false)
	}
}

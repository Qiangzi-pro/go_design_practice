package version1

import (
	"io/ioutil"
	"os"
	"strings"
)

const (
	Separator = string(os.PathSeparator)
)

type IFileSystem interface {
	CountNumOfFiles() int
	CountSizeOfFiles() int64
	GetPath() string
}

type FileSystemNode struct {
	path string
}

func (fsn *FileSystemNode) GetPath() string {
	return fsn.path
}

type File struct {
	FileSystemNode
}

func NewFile(path string) *File {
	return &File{
		FileSystemNode{path: path},
	}
}

func (node *File) CountNumOfFiles() int {
	return 1
}

func (node *File) CountSizeOfFiles() int64 {
	fileInfo, err := os.Stat(node.path)
	if os.IsNotExist(err) {
		return 0
	} else {
		return fileInfo.Size()
	}
}

type IDirSystem interface {
	IFileSystem
	AddSubNode(node IFileSystem)
	RemoveSubNode(node IFileSystem)
}

type Directory struct {
	FileSystemNode
	subNodes []IFileSystem
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
		currentPath := path + Separator + fileInfo.Name()
		if fileInfo.IsDir() {
			node = NewDirectory(currentPath)
		} else {
			node = NewFile(currentPath)
		}
		root.AddSubNode(node)
	}

	return root
}

func (dir *Directory) AddSubNode(node IFileSystem) {
	dir.subNodes = append(dir.subNodes, node)
}

func (dir *Directory) RemoveSubNode(node IFileSystem) {
	i := 0
	for ; i < len(dir.subNodes); i++ {
		if strings.EqualFold(dir.subNodes[i].GetPath(), node.GetPath()) {
			break
		}
	}

	if i < len(dir.subNodes) {
		// 切片删除操作
		dir.subNodes = append(dir.subNodes[:i], dir.subNodes[i+1:]...)
	}
}

func (dir *Directory) CountNumOfFiles() int {
	numOfFiles := 0
	for _, fileOrDir := range dir.subNodes {
		numOfFiles += fileOrDir.CountNumOfFiles()
	}
	return numOfFiles
}

func (dir *Directory) CountSizeOfFiles() int64 {
	SizeOfFiles := int64(0)
	for _, fileOrDir := range dir.subNodes {
		SizeOfFiles += fileOrDir.CountSizeOfFiles()
	}
	return SizeOfFiles
}

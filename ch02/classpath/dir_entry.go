package classpath

import (
	"io/ioutil"
	"path/filepath"
)

//存放类的目录
type DirEntry struct {
	absDir string
}

//生成对象
//存放类的目录
func newDirEntry(path string) *DirEntry {
	//判断文件目录是否存在
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{
		absDir: absPath,
	}
}

//实现Entry接口
//读取类文件为字节数组
//注意： 指针接收器与非指针接收器的区别
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	//目录+类名(包含java package目录的全限定名)
	// /opt + java/lang/Object.class
	fileName := filepath.Join(self.absDir, className)
	//读取文件为字节
	fileBytes, err := ioutil.ReadFile(fileName)
	return fileBytes, self, err
}

//实现Entry接口
//存放class的绝对路径目录
func (self *DirEntry) String() string {
	return self.absDir
}

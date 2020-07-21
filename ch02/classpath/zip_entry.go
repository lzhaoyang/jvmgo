package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

//zip或者jar形式的类路径
type ZipEntry struct {
	absPath string
}

func newZipOrJarEntry(path string) *ZipEntry {
	abs, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{
		absPath: abs,
	}
}

//实现entry接口
//读取jar包中的文件 找到类全限定的类名字，然后读取出字节数组
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	//从压缩文件中读取文件
	zipFile, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer zipFile.Close()
	//读取文件列表
	for _, file := range zipFile.File {
		if file.Name == className {
			targetFile, err := file.Open()
			if err != nil {
				return nil, nil, err
			}
			defer targetFile.Close()
			bytes, err := ioutil.ReadAll(targetFile)
			if err != nil {
				return nil, nil, err
			}
			return bytes, self, nil
		}
	}
	//没找到指定的类
	return nil, nil, errors.New("class not found: " + className)
}

//实现entry接口
func (self *ZipEntry) String() string {
	return self.absPath
}

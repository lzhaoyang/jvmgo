package classpath

import (
	"os"
	"strings"
)

const PathListSeperator = string(os.PathListSeparator)

type Entry interface {
	//读取类
	readClass(className string) ([]byte, Entry, error)
	//返回entry变量的字符串表示
	String() string
}

//创建不同类型的Entry
func newEntry(path string) Entry {
	//如果包含路径分隔符，说明是有多个类路径，这时候需要拆开
	///usr/local/Cellar/go/1.14/libexec/bin:/Users/sunrise/CodeGround/GolandProjects/Go/bin
	// windows下 abc/c;cd/e
	if strings.Contains(path, PathListSeperator) {
		return newCompositeEntry(path)
	}
	//含有通配符的类路径
	if strings.HasSuffix(path, "*") {
		return newWildCardEntry(path)
	}
	//含有压缩文件的类路径 jar包或者zip包
	if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, "JAR") ||
		strings.HasSuffix(path, "zip") || strings.HasSuffix(path, "ZIP") {
		return newZipOrJarEntry(path)
	}
	//java/lang/Object.class
	//普通默认的类路径
	return newDirEntry(path)
}

//含有压缩格式的类路径
func newZipOrJarEntry(path string) Entry {
	return nil
}

//含有类路径通配符的类路径
func newWildCardEntry(path string) Entry {
	return nil
}

//含有路径分割符号的类路径
func newCompositeEntry(path string) Entry {
	return nil
}

package classpath

import (
	"os"
	"strings"
)

const PathListSeperator = string(os.PathListSeparator)

//readClass方法的参数是class文件的相对路径，路径之间用斜线（/）分隔，
//文件名有.class后缀。比如要读取java.lang.Object类，传入的参数应该是java/lang/Object.class。
//返回值是读取到的字节数据、最终定位到class文件的Entry，以及错误信息。
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
	// /opt/java/*
	if strings.HasSuffix(path, "*") {
		return newWildCardEntry(path)
	}
	//含有压缩文件的类路径 jar包或者zip包
	// /opt/java/base.jar
	if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, "JAR") ||
		strings.HasSuffix(path, "zip") || strings.HasSuffix(path, "ZIP") {
		return newZipOrJarEntry(path)
	}
	//普通默认的类路径
	// /opt (相对于类全限定名的相对路径)，即/opt下有java编译后的class包
	return newDirEntry(path)
}

//含有类路径通配符的类路径
func newWildCardEntry(path string) Entry {
	return nil
}

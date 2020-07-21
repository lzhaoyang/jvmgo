package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

//构造函数
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntries := []Entry{}
	for _, path := range strings.Split(pathList, PathListSeperator) {
		entryPath := newEntry(path)
		compositeEntries = append(compositeEntries, entryPath)

	}
	return compositeEntries
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		bytes, entry, err := entry.readClass(className)
		//读取到所需要的类信息
		if err == nil {
			return bytes, entry, nil
		}
	}
	//没读取到所需要的类信息
	return nil, nil, errors.New("class not found: " + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()

	}
	return strings.Join(strs, PathListSeperator)
}

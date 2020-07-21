package classpath

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestDirEntry_readClass(t *testing.T) {
	a := "dir_entry.go"
	abs, err := filepath.Abs(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(abs)
}

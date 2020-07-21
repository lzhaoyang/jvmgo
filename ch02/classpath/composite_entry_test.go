package classpath

import (
	"fmt"
	"testing"
)

func TestCompositeEntry_String(t *testing.T) {
	entry := newCompositeEntry("/root/*:/opt")
	fmt.Println(entry)
}

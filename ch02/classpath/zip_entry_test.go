package classpath

import (
	"archive/zip"
	"fmt"
	"testing"
)

//测试读取压缩文件的
//=== RUN   TestZipEntry_ReadJarOrZip
//0
//META-INF/
//1
//META-INF/MANIFEST.MF
//2
//Test.class
//--- PASS: TestZipEntry_ReadJarOrZip (0.00s)
//PASS
func TestZipEntry_ReadJarOrZip(t *testing.T) {
	reader, err := zip.OpenReader("java-sample-1.0-SNAPSHOT.jar")
	if err != nil {
		panic(err)
	}
	defer reader.Close()
	for i, file := range reader.File {
		fmt.Println(i)
		fmt.Println(file.Name)
	}

}

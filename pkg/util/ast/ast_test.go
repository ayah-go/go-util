package ast

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type Transact struct {
}

// @Transactional
func (*Transact) Before() {
}
func getCurrentPath() string {
	s, err := os.Getwd()
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := s[0 : i+1]
	return path
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func TestScanFuncDeclByComment(t *testing.T) {

	var urlList []string
	println(getCurrentPath())
	// 遍历api文件夹下的所有go文件
	filepath.Walk("D:\\goland_workspaces\\go-temp\\internal\\api\\", func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			bt, _ := ioutil.ReadFile(path)
			src := string(bt)
			result := ScanFuncDeclByComment(``, src, "@CheckToken")
			for _, url := range result {
				urlList = append(urlList, url)
			}
		}
		return nil
	})

	fmt.Print(urlList)
}

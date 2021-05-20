package ast

import (
	"go-util/util/logger"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"
)

/*
实现ast语法树分析,本工具实现提取带指定注释后的文本内容
@example // @CheckToken /url 提取后返回 /url
@author wyy
@see https://github.com/handsomestWei/go-annotation
*/
const commentPrefix = string("// ")

type AnalysisResult struct {
	PkgName     string
	RecvMethods map[string][]MethodInfo // key RecvName
	Funcs       []FuncInfo
}

type MethodInfo struct {
	PkgName    string
	RecvName   string
	MethodName string
	Comment    []string
}

type FuncInfo struct {
	PkgName  string
	FuncName string
	Comment  []string
}

func GetUrlListFromControllerPathByComment(apiPath string, targetComment string) []string {
	var urlList []string
	// 遍历api文件夹下的所有go文件
	err := filepath.Walk(apiPath, func(path string, info fs.FileInfo, err error) error {
		if strings.HasSuffix(path, "_controller.go") {
			bt, _ := ioutil.ReadFile(path)
			logger.L().Debug("扫描文件:" + path)
			src := string(bt)
			result := ScanFuncDeclByComment(``, src, targetComment)
			for _, url := range result {
				urlList = append(urlList, url)
			}
		}
		if err != nil {
			logger.L().Error("注释扫描出现异常", err)
			return nil
		}
		return nil
	})
	if err != nil {
		logger.L().Error("注释扫描出现异常", err)
		return nil
	}
	logger.L().Debug("注释：", targetComment, " : ", urlList)
	return urlList
}

// ScanFuncDeclByComment find func and method in go file by target comment
func ScanFuncDeclByComment(fileName, src, targetComment string) []string {
	fSet := token.NewFileSet()
	f, err := parser.ParseFile(fSet, fileName, src, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	// 定义一个空切片存放url、 list
	var urlList []string
	result := &AnalysisResult{
		RecvMethods: make(map[string][]MethodInfo),
	}
	result.PkgName = f.Name.String()
	for _, d := range f.Decls {
		switch decl := d.(type) {
		case *ast.FuncDecl:
			if decl.Doc != nil {
				param := getContainCommentParam(decl.Doc.List, targetComment)
				if param != "" {
					urlList = append(urlList, param)
				}
			}
		}
	}
	return urlList
}

// 获取指定注解后面的参数
func getContainCommentParam(lines []*ast.Comment, targetComment string) string {
	for _, l := range lines {
		c := strings.TrimSpace(strings.TrimLeft(l.Text, commentPrefix))
		// 首位包含 targetComment 注解 返回后面的内容
		if strings.Index(c, targetComment) == 0 {
			return strings.TrimSpace(strings.TrimLeft(c, targetComment))
		}
	}
	return ""
}

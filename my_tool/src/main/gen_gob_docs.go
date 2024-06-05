package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

// 生成gob的md文件
func main() {
	rootDir := `D:\dev-project\0.demo\go-learning\my_tool\src\main\gen`
	rootDir, _ = filepath.Abs("./my_tool/src/main/gen")
	genGobDocs(rootDir)
}

func genGobDocs(path string) {

	// get packages
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("ParseDir 出错:", err)
		return
	}

	// 开始创建文件夹
	genPath := filepath.Join(path, fmt.Sprintf("gen-%d", time.Now().Unix()))
	if err := os.Mkdir(genPath, 0700); err != nil {
		fmt.Println("创建文件夹出错:", err)
		return
	}

	for _, pkg := range pkgs {
		for _, astFile := range pkg.Files {
			var key, code string
			readme := strings.TrimPrefix(astFile.Doc.List[0].Text, "/*")
			readme = strings.TrimSuffix(readme, "*/")

			for _, v := range astFile.Scope.Objects {
				if v.Kind == ast.Con {
					d := v.Decl.(*ast.ValueSpec)
					key = strings.Trim(d.Values[0].(*ast.BasicLit).Value, `"`)
				}
				if v.Kind == ast.Typ {
					d := v.Decl.(*ast.TypeSpec)

					// 考虑是否能把源代码截取原文处理

					buffer := bytes.NewBufferString("")
					if err = astToGo(buffer, d); err != nil {
						fmt.Println("astToGo 出错:", err)
						return
					}
					code = "```go " + buffer.String() + " ```"
				}
			}

			filePath := filepath.Join(genPath, "aaa.md")
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Println("创建文件出错:", err)
				return
			}

			tpl := template.Must(template.New("first").Parse(md))
			p := make(map[string]string)
			p["key"] = key
			p["readme"] = readme
			p["code"] = code
			if err := tpl.Execute(file, p); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Println("创建md成功, 文件夹地址:", genPath)
	}
}

var md = `---
lang: zh-CN
title: {{.key}}
description:
---
# {{.key}}

{{.readme}}

{{.code}}
`

func astToGo(buffer *bytes.Buffer, node interface{}) error {
	addNewline := func() {
		err := buffer.WriteByte('\n') // add newline
		if err != nil {
			log.Panicln(err)
		}
	}
	addNewline()
	if err := format.Node(buffer, token.NewFileSet(), node); err != nil {
		return err
	}
	addNewline()
	return nil
}

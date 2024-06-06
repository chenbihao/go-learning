package main

import (
	"bufio"
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
)

// 生成 gob 的 md 文件
func main() {
	rootDir, _ := filepath.Abs("./my_tool/src/main/gen")
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
	genPath := filepath.Join(path, fmt.Sprintf("gen"))
	if err = CreateFolderIfNotExists(genPath); err != nil {
		fmt.Println("创建文件 出错:", err)
		return
	}

	for _, pkg := range pkgs {
		for filePath, astFile := range pkg.Files {
			var key, readme, code string
			readme = strings.TrimPrefix(astFile.Doc.List[0].Text, "/*")
			readme = strings.TrimSuffix(readme, "*/")

			for _, v := range astFile.Scope.Objects {
				if v.Kind == ast.Con {
					d := v.Decl.(*ast.ValueSpec)
					key = strings.Trim(d.Values[0].(*ast.BasicLit).Value, `"`)
				}
				if v.Kind == ast.Typ {
					d := v.Decl.(*ast.TypeSpec)

					// 把源代码截取原文处理
					code, err = astToGo2(filePath, fset, d)
					if err != nil {
						fmt.Println("astToGo 出错:", err)
						return
					}
					code = "```go \n" + code + "\n```"
				}
			}

			fileName := filepath.Base(filePath)
			fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
			genFilePath := filepath.Join(genPath, fileName+".md")
			file, err := os.Create(genFilePath)
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

func astToGo2(filePath string, fset *token.FileSet, d *ast.TypeSpec) (s string, err error) {
	pos := fset.Position(d.Type.Pos())
	end := fset.Position(d.Type.End())

	f, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		return s, err
	}
	defer f.Close()

	buf := make([]byte, 1024*4)
	f.Read(buf)
	reader := bufio.NewReader(bytes.NewReader(buf))

	sb := strings.Builder{}
	for i := 1; i <= end.Line; i++ {
		line, _, _ := reader.ReadLine()
		if i >= pos.Line {
			sb.Write(line)
			if i != end.Line {
				sb.Write([]byte("\n"))
			}
		}
	}
	return sb.String(), nil
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	// os.Stat 获取文件信息
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 如果不存在，则创建文件夹
func CreateFolderIfNotExists(folder string) error {
	if !Exists(folder) {
		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

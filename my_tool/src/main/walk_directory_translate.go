package main

import (
	"fmt"
	alimt2 "github.com/alibabacloud-go/alimt-20181012/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"os"
	"path/filepath"
)

type Strategy = int

const (
	strategy_onlyPrint Strategy = iota
	strategy_renameSuffix
	strategy_renameAndSave
	strategy_renameAndSaveTxt
)

const (
	regionId        = "cn-shenzhen"
	accessKeyID     = ""
	accessKeySecret = ""
)

// 遍历并重命名翻译
func main() {
	// 指定要遍历的目录
	rootDir := `Q:\其他\文件\`
	// 调用遍历函数
	walkDirectorysAndTranslate(rootDir, 1, strategy_renameAndSaveTxt)
}

// 遍历并翻译
func walkDirectorysAndTranslate(path string, depth int, strategy Strategy) {
	// 递归结束条件：达到指定的深度
	if depth < 0 {
		return
	}

	// 打印当前目录
	//fmt.Println(path)

	// 读取当前目录下的文件和子目录
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 遍历当前目录下的文件和子目录
	for _, file := range files {

		// 打印子目录下的文件和子目录
		filepathStr := filepath.Join(path, file.Name())
		if file.IsDir() {
			walkDirectorysAndTranslate(filepathStr, depth-1, strategy)
		}

		ts, _ := translate(file.Name())
		fmt.Printf("原名：%s 翻译：%s \n", file.Name(), ts)

		switch strategy {
		case strategy_onlyPrint:
		case strategy_renameSuffix:
			_ = os.Rename(filepathStr, filepathStr+" "+ts)
		case strategy_renameAndSave:
			// 保存 【原名.txt】
			newTxtPath := filepath.Join(filepathStr, "原名："+file.Name())
			f, err := os.Create(newTxtPath + ".txt")
			if err != nil {
				fmt.Printf("Create err：", err.Error())
			}
			_ = f.Close()
			err = os.Rename(filepathStr, filepath.Join(path, ts))
			if err != nil {
				fmt.Printf("Rename err：", err.Error())
			}
		case strategy_renameAndSaveTxt:
			// 保存 file.Name() .txt
			f, err := os.Create(filepath.Join(filepathStr, "原名.txt"))
			if err != nil {
				fmt.Printf("Create err：", err.Error())
			}
			f.WriteString(file.Name())
			_ = f.Close()
			err = os.Rename(filepathStr, filepath.Join(path, ts))
			if err != nil {
				fmt.Printf("Rename err：", err.Error())
			}
		}
	}
	return
}

func translate(source string) (result string, err error) {
	client, err := CreateClient(accessKeyID, accessKeySecret, regionId)
	if err != nil {
		return result, err
	}

	request := &alimt2.TranslateGeneralRequest{}
	request.SetFormatType("text")
	request.SetSourceLanguage("auto")
	request.SetTargetLanguage("zh")
	request.SetSourceText(source)

	response, err := client.TranslateGeneral(request)
	if err != nil {
		return result, err
	}
	result = *response.Body.Data.Translated
	return
}

func CreateClient(accessKeyId string, accessKeySecret string, regionId string) (_result *alimt2.Client, _err error) {
	config := &openapi.Config{}
	config.AccessKeyId = &accessKeyId
	config.AccessKeySecret = &accessKeySecret
	config.RegionId = &regionId
	_result = &alimt2.Client{}
	_result, _err = alimt2.NewClient(config)
	return _result, _err
}

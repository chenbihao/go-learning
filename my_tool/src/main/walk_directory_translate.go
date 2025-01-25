package main

import (
	"fmt"
	alimt2 "github.com/alibabacloud-go/alimt-20181012/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"os"
	"path/filepath"
	"strings"
)

type RenameStrategy = string
type BackupStrategy = string

const (
	rename_replace RenameStrategy = "替换"
	rename_prefix                 = "翻译在前"
	rename_suffix                 = "翻译在后"
)
const (
	backup_non          BackupStrategy = "暂无备份"
	backup_source_name                 = "源备份1"
	backup_source_name2                = "源备份2"
)

// 阿里云 翻译API 接入
const (
	regionId        = "cn-shenzhen"
	accessKeyID     = ""
	accessKeySecret = ""
)

/*
遍历文件并翻译重命名
- 支持预览
- 多种重源格式（例如掺杂下划线）
- 多种重命名格式（替换、前缀、后缀）
- todo：
- 初始版本名称回滚等操作
- 屏蔽某些文件： ".DS_Store"、".xxxx"
- 屏蔽某些文件夹： "__MACOSX"
- 支持自定义字典、自定义保留不翻译文本
- 翻译字符预估、检测是否已翻译
*/
func main() {

	// 指定要遍历的目录
	rootDir := `D:\DevProjects\TempFile`

	onlyPreview := false
	onlyFileNotDir := false
	depth := 10
	renameStrategy := rename_suffix
	backupStrategy := backup_non

	fmt.Print("====================================================== \n")
	fmt.Printf("========= 只预览：%t 只文件：%t 遍历层数：%d  \n", onlyPreview, onlyFileNotDir, depth)
	fmt.Printf("========= 重命名策略：%s 备份策略：%s  \n", renameStrategy, backupStrategy)
	fmt.Print("====================================================== \n")

	walkDirectorysAndTranslate(rootDir, onlyPreview, onlyFileNotDir, depth, renameStrategy, backupStrategy)
}

// 遍历并翻译
func walkDirectorysAndTranslate(pathStr string, onlyPreview, onlyFileNotDir bool, depth int,
	renameStrategy RenameStrategy, backupStrategy BackupStrategy) {

	// 递归结束条件：达到指定的深度
	if depth < 0 {
		return
	}

	// 读取当前目录下的文件和子目录
	files, err := os.ReadDir(pathStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 遍历当前目录下的文件和子目录
	for _, file := range files {
		filePathStr := filepath.Join(pathStr, file.Name())
		if file.IsDir() {
			walkDirectorysAndTranslate(filePathStr, onlyPreview, onlyFileNotDir, depth-1, renameStrategy, backupStrategy)
		}
		rename(onlyPreview, onlyFileNotDir, file, pathStr, filePathStr, renameStrategy, backupStrategy)
	}
	return
}

func rename(onlyPreview, onlyFileNotDir bool, file os.DirEntry, pathStr, filePathStr string,
	renameStrategy RenameStrategy, backupStrategy BackupStrategy) {

	filename := strings.ReplaceAll(file.Name(), "_", " ")

	var nonExtensionFileName, extension string
	translateStr := file.Name()

	if file.IsDir() {
		if onlyFileNotDir {
			fmt.Printf("================ 忽略文件夹： %s ================ \n", filename)
			return
		}
		translateStr, _ = translate(filename)
		nonExtensionFileName = filename
		fmt.Print("文件夹 ---------------- \n")
	} else {
		// 获取后缀
		extension = filepath.Ext(filename)
		// 获取非后缀部分
		nonExtensionFileName = strings.TrimSuffix(filename, extension)
		translateStr, _ = translate(nonExtensionFileName)
	}
	fmt.Printf("原名：%s\n", filename)

	switch renameStrategy {
	case rename_replace:
		newName := translateStr + extension
		fmt.Printf("翻译：%s\n", translateStr+extension)
		if !onlyPreview {
			_ = os.Rename(filePathStr, filepath.Join(pathStr, newName))
		}
	case rename_prefix: // 翻译在前
		newName := translateStr + " " + file.Name()
		fmt.Printf("新名：%s\n", newName)
		if !onlyPreview {
			_ = os.Rename(filePathStr, filepath.Join(pathStr, newName))
		}
	case rename_suffix: // 翻译在后
		newName := nonExtensionFileName + " " + translateStr + extension
		fmt.Printf("新名：%s\n", newName)
		if !onlyPreview {
			_ = os.Rename(filePathStr, filepath.Join(pathStr, newName))
		}
	}

	if file.IsDir() {
		fmt.Print("文件夹翻译结束 ================ \n")
	}

	//switch backupStrategy {
	//case backup_non:
	//case backup_source_name:
	//	// 保存 file.Name().txt
	//	newTxtPath := filepath.Join(filepathStr, "原名："+filename)
	//	f, err := os.Create(newTxtPath + ".txt")
	//	if err != nil {
	//		fmt.Printf("Create err：", err.Error())
	//	}
	//	_ = f.Close()
	//case backup_source_name2:
	//	// 保存 【原名.txt】
	//	f, err := os.Create(filepath.Join(filepathStr, "原名.txt"))
	//	if err != nil {
	//		fmt.Printf("Create err：", err.Error())
	//	}
	//	f.WriteString(filename)
	//	_ = f.Close()
	//}

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

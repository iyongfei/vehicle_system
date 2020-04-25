package main

import (
	"encoding/csv"
	"log"
	"os"
)

//func Exists(path string) bool {
//	_, err := os.Stat(path)    //os.Stat获取文件信息
//	if err != nil {
//		if os.IsExist(err) {
//			return true
//		}
//		return false
//	}
//	return true
//}
//
//// 判断所给路径是否为文件夹
//func IsDir(path string) bool {
//	s, err := os.Stat(path)
//	if err != nil {
//		return false
//	}
//	return s.IsDir()
//}
//
//// 判断所给路径是否为文件
//func IsFile(path string) bool {
//	return !IsDir(path)
//}
//

func IsFolderDir(path string) bool {
	fi, e := os.Stat(path)
	if e != nil {
		return false
	}
	return fi.IsDir()
}
func IsFileDir(path string) bool {
	fi, e := os.Stat(path)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

type A struct {
	Name string
	Sex  string
}

func main() {

	//filename := "test.csv"
	//columns := [][]string{{"姓名", "电话", "公司", "职位", "加入时间"}, {"1", "2", "刘犇,刘犇,刘犇", "4", "5"}}
	//ExportCsv(filename, columns)
}

func ExportCsv(filePath string, data [][]string) {
	fp, err := os.Create(filePath) // 创建文件句柄
	if err != nil {
		log.Fatalf("创建文件["+filePath+"]句柄失败,%v", err)
		return
	}
	defer fp.Close()
	fp.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(fp)         //创建一个新的写入文件流
	w.WriteAll(data)
	w.Flush()
}

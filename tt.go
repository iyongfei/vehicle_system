package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
)

func main() {
	a := "wnn"
	//rr := a[1:]
	//fmt.Println(rr)
	fmt.Println(a[:0])

	//fmt.Println(r)
	//filename := "test.csv"
	//columns := [][]string{{"姓名", "电话", "公司", "职位", "加入时间"}, {"1", "2", "刘犇,刘犇,刘犇", "4", "5"}}
	//ExportCsv(filename, columns)

	//AA("1", "2")
}

func AA(args ...interface{}) {
	for k, v := range args {
		fmt.Println(reflect.TypeOf(v))
		switch v.(type) {

		case string:
			fmt.Println(k, v)

		}

	}
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

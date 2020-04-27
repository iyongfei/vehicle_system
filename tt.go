package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

func ExampleReader() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	recordLine := 0
	for {

		if recordLine == 0 {
			record, err := r.Read()
			if err == io.EOF {
				fmt.Println(record, "1")
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(record)
		} else {
			record, err := r.Read()
			if err == io.EOF {
				fmt.Println()
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(record)
		}

		recordLine++

	}
	// Output:
	// [first_name last_name username]
	// [Rob Pike rob]
	// [Ken Thompson ken]
	// [Robert Griesemer gri]
}
func main() {

	iportsMap := map[string]map[string][]uint32{}
	var a map[string][]uint32
	a["v"] = []uint32{1, 2, 3}

	iportsMap["a"] = a
	fmt.Println(iportsMap["a"])

	//filename := "test.csv"
	//columns := [][]string{{"姓名", "电话", "公司", "职位", "加入时间"}, {"1", "2", "刘犇,刘犇,刘犇", "4", "5"}}
	//ExportCsv(filename, columns)
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

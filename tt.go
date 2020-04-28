package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
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

const (
	AddCsv  = iota
	EditCsv = 2
)

func main() {

	r := strings.Trim("   ew  ", " ")
	fmt.Println(r)

	//filename := "test.csv"
	//columns := [][]string{{"姓名", "电话", "公司", "职位", "加入时间"}, {"1", "2", "刘犇,刘犇,刘犇", "4", "5"}}
	//ExportCsv(filename, columns)
}
func StampUnix2Time(timestamp int64) time.Time {
	datetime := time.Unix(timestamp, 0)
	return datetime
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

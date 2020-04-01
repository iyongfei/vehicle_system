package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func WriteContent2File(filename string, str_content string) (bool ,error) {
	fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	buf := []byte(str_content)
	_,err := fd.Write(buf)
	//defer fd.Close()

	if err != nil{
		return false,err
	}
	return true,nil
}

func Write(file *os.File ,data string)(bool ,error)  {
	_,err := file.WriteString(data)
	if err != nil{
		return false,err
	}
	return true,nil
}




func IsDir(dirname string) bool  {
	fhandler, err := os.Stat(dirname);
	if(! (err == nil || os.IsExist(err)) ) {
		return false
	}else {
		return fhandler.IsDir()
	}
}

func IsFile(filename string) bool  {
	fhandler, err := os.Stat(filename);
	if(! (err == nil || os.IsExist(err)) ) {
		return false
	}else if (fhandler.IsDir()){
		return false
	}
	return true
}


func IsExist(path string) (bool, error){
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileByteSize(filename string) (bool,int64) {
	if (! IsFile(filename)) {
		return false,0
	}
	fhandler, _ := os.Stat(filename);
	return true,fhandler.Size()
}


func getAllFiles(folder string) []string{
	filesNameList:=[]string{}
	files,err:= ioutil.ReadDir(folder)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _,file := range files{
		if file.IsDir(){
			continue
		}else {
			filesNameList = append(filesNameList,file.Name())
		}
	}
	return filesNameList
}



func GetCurPath() string {
	file, _ := exec.LookPath(os.Args[0])

	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
	path, _ := filepath.Abs(file)

	rst := filepath.Dir(path)

	return rst
}
func GetCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])
	//checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}
func GetCurrentPath1() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		//return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}
func GetCurrentDirectory1() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

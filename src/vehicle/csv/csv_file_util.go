package csv

import (
	"os"
	"path"
)

//判断是否是一个合法的文件路径(非文件夹)

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

func IsExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Getwd() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return dir

}

//有该目录会报错，创建下级目录也会报错
func Mkdir(path string) error {
	err := os.Mkdir(path, os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}

//可以创建单级目录，或者多级目录，而且不会覆盖目录下其他文件
func MkdirAll(path string) error {
	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}

func Rename() {
}

func Remove() {

}

func RemoveAll() {

}

//会清空文件，不能创建子级
func Create() {

}

//打开只读
func Open() {
}

//创建不会清空
func NewFile() {

}

//特定模式
func OpenFile() {

}

func IsAbs(pathName string) bool {
	return path.IsAbs(pathName)
}

func PathDir(pathName string) string {
	return path.Dir(pathName)

}

//判断目录是否存在

//创建目录

//判断文件是否存在

//创建文件

package util

import (
	"archive/tar"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)


/*
读取文件内容
参数：路径
返回：内容
*/
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

/*
生成文件并写入
*/
func WriteWithIoutil(filname string , content []byte) {
	ioutil.WriteFile(filname, content, 0644)

}

func VHardwareTempFileExist() bool {
	uncompresPath := "vfiles/uncompress_temp"
	files := VGetAllFiles(uncompresPath)
	if len(files) <= 0{
		return false
	}
	return true
}

func VHardwareFileExist() bool {
	uncompresPath := "vfiles/uncompress"
	files := VGetAllFiles(uncompresPath)
	if len(files) <= 0{
		 return false
	}
	return true
}

//path路径文件是否存在
func VFileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

//获取path路径下的所有文件路径
func VGetAllFilePaths(path string) (files []string) {
	_ = filepath.Walk(path, func(filePath string, f os.FileInfo, err error) error {

		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		r, _ := regexp.Compile(`^\.(.)*`)

		if r.MatchString(f.Name()) {
			return nil
		}

		if f.IsDir() == false {
			files = append(files, filePath)
		}

		return nil
	})
	return files
}

//获取path路径下的所有文件名称
func VGetAllFiles(path string) (files []string) {
	_ = filepath.Walk(path, func(filePath string, f os.FileInfo, err error) error {

		if f == nil {
			return err
		}

		if f.IsDir() {

			return nil
		}

		r, _ := regexp.Compile(`^\.(.)*`)
		if r.MatchString(f.Name()) {
			return nil
		}

		if f.IsDir() == false {
			fmt.Println(filePath)
			files = append(files, f.Name())
		}

		return nil
	})

	return files
}

//获取文件修改时间戳
//path 文件路径
func VGetFileModTime(path string) int64 {
	f, err := os.Open(path)
	if err != nil {
		log.Println("open file error")
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		log.Println("stat fileinfo error")
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}

///获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".go")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}

//循环遍历分割字符串
//str 初始字符串，component 分割字符串，sli 保存切片
func TracerString(str string,component string, sli []string) []string {
	slices := []string{}
	if len(sli) >0 {
		slices = sli
	}
	tracer := str
	comma := strings.Index(tracer, component)
	if comma >= 0{
		pos := strings.Index(tracer[comma:], "")
		if pos >= 0{
			slices =append(slices,tracer[0:comma])
			slices = TracerString(tracer[comma+pos+1:],component,slices)
		}
	}else {
		if len(tracer) >0{
			slices =append(slices,tracer[:])
		}
	}
	return slices
}

func UncompressFile(path string) error{
	// file read
	//fr, err := os.Open("vfiles/compress/lin_golang_src.tar.gz")
	fr, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fr.Close()
	// gzip read
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()
	// tar read
	tr := tar.NewReader(gr)
	// 读取文件
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
		}
		// 显示文件
		if strings.Contains(h.Name,"/") {

		}
		filename := "vfiles/uncompress_temp/"
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			os.Mkdir(filename,os.ModePerm)
		}
		res:= strings.Index(h.Name,"/")
		if res>0 {
			content := h.Name[0 : res]
			plusName := filename+content+"/"
			if _, err := os.Stat(plusName); os.IsNotExist(err) {
				os.Mkdir(plusName,os.ModePerm)
			}
		}

		// 打开文件vfiles/uncompress_temp/Venus_1.2.3.20819_53c223dc_mt7621_tianqi-R201b_1d930cfd6ba311da5ef6e42ce80b491b_sysupgrade.bin
		fw, err := os.OpenFile(filename + h.Name, os.O_CREATE | os.O_WRONLY, 0644/*os.FileMode(h.Mode)*/)
		if err != nil {

		}
		defer fw.Close()
		// 写文件
		_, err = io.Copy(fw, tr)
		if err != nil {
			return err
		}
	}
	return nil
}

func HashFileMd5(filePath string) (md5Str string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return
	}
	hashInBytes := hash.Sum(nil)[:16]
	md5Str = hex.EncodeToString(hashInBytes)
	return
}

func Compress()  {
	// file write
	fw, err := os.Create("vfiles/compress/src.tar.gz")
	if err != nil {
		fmt.Println("createerror",err)
	}
	defer fw.Close()
	// gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()
	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()
	// 打开文件夹
	dir, err := os.Open("vfiles/uncompress_temp/src")
	if err != nil {
		panic(nil)
	}
	defer dir.Close()
	// 读取文件列表
	fis, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}
	// 遍历文件列表
	for _, fi := range fis {
		// 逃过文件夹, 我这里就不递归了
		if fi.IsDir() {
			continue
		}
		// 打印文件名称
		fmt.Println(fi.Name())
		// 打开文件
		fr, err := os.Open(dir.Name() + "/" + fi.Name())
		if err != nil {
			panic(err)
		}
		defer fr.Close()
		// 信息头
		h := new(tar.Header)
		h.Name = fi.Name()
		h.Size = fi.Size()
		h.Mode = int64(fi.Mode())
		h.ModTime = fi.ModTime()
		// 写信息头
		err = tw.WriteHeader(h)
		if err != nil {
			panic(err)
		}
		// 写文件
		_, err = io.Copy(tw, fr)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("tar.gz ok")
}


/**
获取bin包的md5
 */
func GetFirmwareMd5(firmwareName string) string {
	firmwareBinSlice := []string{}
	var firmwareMd5 string
	//bin包格式
	if strings.Contains(firmwareName,".bin") {
		//切割bin包
		firmwareBinSlice = TracerString(TracerString(firmwareName,".bin",
			firmwareBinSlice)[0],"_",firmwareBinSlice)
		//md5长度
		if len(firmwareBinSlice[5])== 32 {
			Upload_Uncompressed_Path := "vfiles/uncompress"

			listFilePath := VGetAllFilePaths(Upload_Uncompressed_Path)
			for _,filePath := range listFilePath{
				if strings.Contains(filePath,".bin") {
					//教验
					md5Valid,_ := HashFileMd5(filePath)
					if md5Valid == firmwareBinSlice[5] {
						firmwareMd5 = firmwareBinSlice[5]
					}
				}
			}
		}
	}
	return firmwareMd5
}
package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"syscall"
	//"github.com/pquerna/ffjson/ffjson"
)

// ReadFile read buf from path and return string buf
// return string
func ReadFile(path string) (string, error) {
	fi, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return string(fd), nil
}

// ReadFileByte read buf from path and return string buf
// return byte
func ReadFileByte(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return fd, nil
}

// PathExist check file exist or not
func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// WriteFileWithLock write buf to file with LINUX LOCK
func WriteFileWithLock(path string, buf []byte) error {
	if !PathExist(path) {
		_, err := os.Create(path)
		if err != nil {
			return err
		}
	}

	fwriter, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer fwriter.Close()

	syscall.Flock(int(fwriter.Fd()), syscall.LOCK_EX)
	defer syscall.Flock(int(fwriter.Fd()), syscall.LOCK_UN)

	_, err = fwriter.Write(buf)
	if err != nil {
		return err
	}

	return nil
}

// WriteFile write buf to file no lock
func WriteFile(path string, buf []byte) error {
	if !PathExist(path) {
		_, err := os.Create(path)
		if err != nil {
			return err
		}
	}

	fwriter, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer fwriter.Close()

	_, err = fwriter.Write(buf)
	if err != nil {
		return err
	}

	return nil
}

// WriteAppendWithLock write content append filename with file mutx lock
func WriteAppendWithLock(path string, buf []byte) error {
	if !PathExist(path) {
		_, err := os.Create(path)
		if err != nil {
			return err
		}
	}
	fwriter, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer fwriter.Close()

	syscall.Flock(int(fwriter.Fd()), syscall.LOCK_EX)
	defer syscall.Flock(int(fwriter.Fd()), syscall.LOCK_UN)
	_, err = fwriter.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

// ReadFileWithLock read file with LINUX LOCK
func ReadFileWithLock(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	syscall.Flock(int(fi.Fd()), syscall.LOCK_EX)
	defer syscall.Flock(int(fi.Fd()), syscall.LOCK_UN)
	fd, err := ioutil.ReadAll(fi)
	return fd, err
}

// FileMd5Checksum only check file's md5sum
func FileMd5Checksum(checkTargetFilename string) ([]byte, error) {
	if !PathExist(checkTargetFilename) {
		return nil, fmt.Errorf("file %s not found", checkTargetFilename)
	}
	file, err := os.Open(checkTargetFilename)
	if err != nil {
		return nil, err
	}

	md5h := md5.New()
	io.Copy(md5h, file)

	//filename 	md5checksum
	md5Result := md5h.Sum([]byte(""))

	return md5Result, nil
}

// Md5Checksum2File write the md5 sums of checkTargetFilename file
// into md5Filename
// checkTargetFilename: file need to be md5 checksum
func Md5Checksum2File(checkTargetFilename, md5Filename string) ([]byte, error) {
	file, err := os.Open(checkTargetFilename)
	if err != nil {
		return nil, err
	}

	md5h := md5.New()
	io.Copy(md5h, file)

	//md5checksum  file base name
	fileBaseName := path.Base(checkTargetFilename)
	md5Result := fmt.Sprintf("%x\t%s\n", md5h.Sum([]byte("")), fileBaseName)

	md5File, err := os.OpenFile(md5Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return nil, err
	}
	defer md5File.Close()

	// Acquire the lock
	syscall.Flock(int(md5File.Fd()), syscall.LOCK_EX)
	defer syscall.Flock(int(md5File.Fd()), syscall.LOCK_UN)
	_, err = md5File.Write([]byte(md5Result)) //md5
	if err != nil {
		return nil, err
	}

	return []byte(md5Result), nil
}

// ListFiles return the file list of th path
// the result will egnore directories, only files, and only filename
func ListFiles(path string) ([]string, error) {
	var flist []string

	err := filepath.Walk(path, func(filepath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		isDir := f.IsDir()
		if isDir == false {
			flist = append(flist, f.Name())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return flist, nil
}

// ListDirectorys return the directorys under path
// it will ignore files
func ListDirectorys(path string) ([]string, error) {
	var flist []string

	err := filepath.Walk(path, func(filepath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		isDir := f.IsDir()
		if isDir == true {
			dirName := f.Name()
			if dirName == "." || dirName == ".." {
				return nil
			}

			flist = append(flist, f.Name())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return flist, nil
}

// ListFilesAbsPath list the files under path
// the result will return the file name include ABS name
func ListFilesAbsPath(path string) ([]string, error) {
	var absflist []string

	err := filepath.Walk(path, func(filepath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		isDir := f.IsDir()
		if isDir == false {
			absflist = append(absflist, filepath)
			return nil
		}

		return nil
	})

	if err != nil {
		return absflist, err
	}

	return absflist, nil
}

// WriteJSONFileWithLock trans v into JSON-DATA and save into config path
//func WriteJSONFileWithLock(v interface{}, path string) error {
//	jdata, err := ffjson.Marshal(v)
//	if err != nil {
//		return err
//	}
//
//	err = WriteFileWithLock(path, jdata)
//	return err
//}
//
//// ReadJSONFileWithLock trans file-data into JSON
//func ReadJSONFileWithLock(v interface{}, path string) error {
//	jdata, err := ReadFileWithLock(path)
//	if err != nil {
//		return err
//	}
//
//	err = ffjson.Unmarshal(jdata, v)
//	return err
//}

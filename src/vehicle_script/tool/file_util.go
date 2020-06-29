package tool

import (
	"io/ioutil"
	"os"
)

// PathExist check file exist or not
func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
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

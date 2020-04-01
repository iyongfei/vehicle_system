package util

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"time"
)




func GetUniqID() string {
	return strconv.FormatUint((uint64)(time.Now().UnixNano()), 10)
}

//生成32位md5字串
func Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}


func RandomAlternativeString(ln int) string {
	letters := []rune("12")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}

func RandomString(ln int) string {
	letters := []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, ln)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}

	return string(b)
}

///////////////////////////////////////////
func CreatRandowUUId()uuid.UUID{
	const UTag = "B64UUID"
	newUUID,_:=uuid.NewRandom()

	return newUUID
}

func CreatUUId()string{
	const UTag = "B64UUID"
	newUUID,err:=uuid.NewRandom()

	if err!=nil{
		Vprintf(UTag,"error:%v\n",err)
	}
	newUUIDBinary,_:=newUUID.MarshalBinary()
	return Base64Encode(newUUIDBinary)
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Decode(str string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	return data, err
}


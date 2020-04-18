package util

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//快速排序（倒序：[4,3,2,1]）
func QuickSort(array []int, left, right int) {
	if left >= right {
		return
	}
	index := partition(array, left, right)
	QuickSort(array, left, index-1)
	QuickSort(array, index+1, right)
}

func partition(array []int, left, right int) int {
	baseNum := array[left]
	for left < right {
		for (array[right] <= baseNum && right > left) {
			right--
		}
		array[left] = array[right]
		for (array[left] >= baseNum && right > left) {
			left++
		}
		array[right] = array[left]
	}
	array[right] = baseNum
	return right
}

//切片去重
//var arr = []string{"hello", "hi", "world", "hi", "china", "hello", "hi"}
func RemoveRepeatedForSlice(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			if arr[i] != "" {
				newArr = append(newArr, arr[i])
			}
		}
	}
	return
}

//截取之后的字符串
func GetSomeString(strs string, split string) string {
	begin := strings.Index(strs, split)
	content := strs[begin+1 : len(strs)]
	return content
}

//截取之前的字符串 unused
func SubStringTopic(strs string, split string) string {
	begin := strings.Index(strs, split)
	content := strs[:begin ]
	return content
}

/**
字符串是否包含字母
 */

func ContainsLetters(s string) bool {
	pattern := ".*[a-zA-z].*" //反斜杠要转义
	result, _ := regexp.MatchString(pattern, s)
	return result

}

/**
获取特定的固件版本号
 */
func GetPrecedingStr(str string, indexStr string) string {
	begin := strings.LastIndex(str, indexStr)

	if begin == -1 {
		return str
	}
	lastPart := str[begin+1:]
	if !ContainsLetters(lastPart) {
		return str
	}

	content := str[:begin]
	return content
}

func FilterSqlWhere(sqlStr string) string {
	trimStr := strings.Replace(sqlStr, " ", "", -1)
	replaceRet := strings.Replace(trimStr, "?", "? and ", -1)
	andLastIndex := strings.LastIndex(replaceRet, "and")
	var fatr string
	if andLastIndex != -1 {
		fatr = replaceRet[:andLastIndex]
	} else {
		fatr = sqlStr
	}
	return fatr
}

/**
切片中是否存在某个元素
 */
func IsEleExistInSlice(valueParam interface{}, array []interface{}) bool {
	for _, value := range array {
		if value == valueParam {
			return true
		}
	}
	return false
}

/**
以某条件拼接切片
 */
func JoinSlice(slice []string, condition string) string {
	ret := ""
	for index, indexV := range slice {
		if index < len(slice)-1 {
			ret += indexV + ","
		} else if index == len(slice)-1 {
			ret += indexV
		}
	}
	return ret
}

/**
切片中是否存在某个值
*/
func IsExistInSlice(valueParam interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		arraySlice := reflect.ValueOf(array)
		for i := 0; i < arraySlice.Len(); i++ {
			if reflect.DeepEqual(valueParam, arraySlice.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

/**
字符串是否以某值结尾
strings.HasPrefix(s, prefix string) bool
HasSuffix 判断字符串 s 是否以 suffix 结尾：
strings.HasSuffix(s, suffix string) bool
 */

func HasSuffix(str, suffix string) bool {
	return strings.HasSuffix(str, suffix)
}


/**
是否是邮件格式
 */

func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

/**
是否是电话格式
 */
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

/**
是否是url格式
 */
func VerifyUrl(urls []string) bool {
	var urlReg bool = true
	for _, url := range urls {
		validURL := govalidator.IsURL(url)
		if !validURL {
			urlReg = false
		}
	}
	return urlReg
}


/**
字符串是否包含数字
 */

func ContainsNumber(s string) bool {
	pattern := ".*[0-9].*" //反斜杠要转义
	result, _ := regexp.MatchString(pattern, s)
	return result

}

/**
字符串是否都是中文
 */

func VerifyChineseFormat(s string) bool {
	pattern := "[\\u4e00-\\u9fa5]+" //反斜杠要转义
	result, _ := regexp.MatchString(pattern, s)
	return result

}

/**
字符串是否都是英文
 */

func VerifyWordsFormat(s string) bool {
	pattern := "[a-zA-Z]+" //反斜杠要转义
	result, _ := regexp.MatchString(pattern, s)
	return result

}

/**
字符串是否只有英文数字的字符串
 */

func VerifyWordsAndNumberFormat(s string) bool {
	pattern := "[a-zA-Z0-9]+" //反斜杠要转义
	result, _ := regexp.MatchString(pattern, s)
	return result

}
func IsSingleDigit(data string) bool {
	digit := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, item := range digit {
		if data == item {
			return true
		}
	}
	return false
}

func IsDigit(data string) bool {
	for _, item := range data {
		if IsSingleDigit(string(item)) {
			continue
		} else {
			return false
		}
	}
	return true
}


/**
是否是15位数字
 */

func VerifyLengthNumberFormat(num string,length int) bool {
	regular := fmt.Sprintf("^([0-9])\\d{%d}$",length)
	reg := regexp.MustCompile(regular)
	return reg.MatchString(num)
}

func VerifyNumberFormat(num string)bool  {
	pattern := "\\d+"
	result,_ := regexp.MatchString(pattern,num)

	return result
}


/**
生成多少以内的数字
[0,n) //[0,65535)
[1,65535]
 */
func RandToMaxNumber(max int) int32 {
	r:=rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(max))
	return r
}

//传10，生成[1,10]的随机数

func RandOneToMaxNumber(max int) int32 {
	r:=rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(int32(max))
	r ++
	return r
}


//判断是不是一个端口
func VerifyIpPort(num string)bool  {

	pattern := "\\d+"
	_,err := regexp.MatchString(pattern,num)
	if err!=nil{
		return false
	}

	parseRt,err:=strconv.Atoi(num)

	if err!=nil{
		return  false
	}


	if 0  < parseRt && parseRt <= 65535{
		return true
	}
	return false
}

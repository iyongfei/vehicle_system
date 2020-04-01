package util

import (
	"fmt"
	"reflect"
	"strings"
)

/**
slice的共同因素
 */

func SliceCommonElements(left interface{},right interface{})(interface{},error){

	leftType:=reflect.TypeOf(left)
	rightType:=reflect.TypeOf(right)

	if leftType != rightType{
		return nil,fmt.Errorf("leftType %s inconsistency rightType %s ",leftType,rightType)

	}

	var leftTypeData []string
	var rightTypeData []string

	switch left.(type) {
	case []string:
		switch right.(type) {
		case []string:
			leftTypeData = left.([]string)
			rightTypeData = right.([]string)
		}

	}

	var result []interface{}
	for _,leftE:=range leftTypeData{
		for _,rightE:= range rightTypeData{
			if leftE == rightE{
				result = append(result,leftE)
			}
		}
	}
	return result,nil
}


/**
反转切片
s1:=[]string{"s","b","SF","jwke"}
common_util.ReverseSlice(s1[:2])
结果为[b s SF jwke]
 */
func ReverseSlice(ele []string)  {
	for i,j:=0,len(ele)-1; i<j;i,j = i +1 ,j-1{
		ele[i],ele[j] = ele[j],ele[i]
	}
}

/**
判断2个切片是否拥有相同元素
 */
func EqualSlice(left,right []string) bool {
	if len(left) != len(right){
		return false
	}

	for leftIndex:=range left{
		if left[leftIndex] != right[leftIndex]{
			return false
		}
	}
	return true
}


/**
去除空字符串1
 */
func RemoveSliceEmptyByCutting(slice []string) []string  {
	i:=0
	for _,ele:=range slice{
		if ele != "" && strings.Trim(ele," ") != ""{
			slice[i] = ele
			i ++
		}
	}

	return slice[:i]
}




/**
去除空字符串2
*/

func RemoveSliceEmptyByAppend(slice []string) []string  {
	retSlice:=slice[:0]
	for _,ele:=range slice{
		if  ele != "" && strings.Trim(ele," ") != ""{
			retSlice = append(retSlice,ele)
		}
	}
	return  retSlice
}

/**
移除某个元素
 */

func RemoveSliceEle(slice []string,index int) interface{} {
	copy(slice[index:],slice[index+1:])
	return slice[:len(slice)-1]
}



/**
判断2个map是否有相同的键值
!ok用来区分元素不存在,元素存为零值
 */

func EqualMap(left ,right map[string]int) bool {
	if len(left) != len(right){
		return false
	}

	for k,xv:=range left{
		if yv,ok:= right[k]; !ok|| yv != xv{

			return  false
		}
	}
	return true
}

/**
比較版本大小
 */
func CompareVersion(left,right string) int {
	minus := -2
	if left < right{
		minus = 1
	}else if left == left {
		minus = 0
	}else {
		minus = -1
	}
	return minus
}






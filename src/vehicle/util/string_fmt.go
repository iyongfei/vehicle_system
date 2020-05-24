package util

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"math/rand"

	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// IsPhoneNumber check value is phone number or not
func IsPhoneNumber(value string) bool {
	phlen := len(value)
	tmp := "12"
	head := tmp[0]

	if phlen < 11 {
		return false
	} else if value[0] != head || (govalidator.IsNumeric(value) == false) {
		return false
	}

	return true
}

// IsMacAddress check mac is valid or not
// Only allowed format:
// 		01:23:45:67:89:AB
//		01.23.45.67.89.AB
// 		0123456789AB
// 00000000000 will be trated as invalid mac
func IsMacAddress(mac string) bool {
	for i, perc := range []byte(mac) {
		if perc != '0' {
			break
		}

		i++
		if i == 12 {
			return false
		}
	}

	if govalidator.IsMAC(mac) == true {
		macLenMax := len("01:23:45:67:89:ab")
		if len(mac) > macLenMax {
			return false
		}

		dotFieldLen := 0
		for i := 0; i < len(mac); i++ {
			if mac[i] == '-' {
				return false
			}

			if mac[i] == '.' {
				if dotFieldLen > 2 {
					return false
				}
				dotFieldLen = 0
			}
			dotFieldLen++
		}
	} else {
		if len(mac) != len("0123456789AB") {
			return false
		}

		if govalidator.IsHexadecimal(mac) == false {
			return false
		}
	}

	return true
}

// IsBase64 return true if str is base64 encoded
func IsBase64(str string) bool {
	return govalidator.IsBase64(str)
}

// ToHexadecimalMac to Hexadecimal Mac address with upper character
func ToHexadecimalMac(macStr string) (string, error) {
	macRet := ""
	var err error
	if IsMacAddress(macStr) == false {
		err = errors.New("mac source is invalid")

		return macRet, err
	}

	var macBytes []byte
	for i := 0; i < len(macStr); i++ {
		if macStr[i] == '.' || macStr[i] == ':' {
			continue
		} else {
			macBytes = append(macBytes, byte(macStr[i]))
		}
	}

	macRet = string(macBytes)
	macRet = strings.ToUpper(macRet)
	return macRet, nil
}

// ToColonFormatedMac to mac splited with colon
// 112233aabbcc to 11:22:33:aa:bb:cc
func ToColonFormatedMac(macStr string) (string, error) {
	mac, err := ToHexadecimalMac(macStr)
	if err != nil {
		return "", err
	}

	macByte := []byte(mac)
	macRet := fmt.Sprintf("%c%c:%c%c:%c%c:%c%c:%c%c:%c%c",
		macByte[0], macByte[1],
		macByte[2], macByte[3],
		macByte[4], macByte[5],
		macByte[6], macByte[7],
		macByte[8], macByte[9],
		macByte[10], macByte[11])

	return macRet, nil
}

// VersionCompare check version
// version a > version b  return 1, else return -1
// if version equal, return 0
//func VersionCompare(vera, verb []byte) int {
//	veraLen := len(vera)
//	verbLen := len(verb)
//
//	if verbLen == 0 { // no ipk in local
//		return -1
//	}
//
//	ret := 0
//	verLenMin := MinInt(veraLen, verbLen)
//
//	for i := 0; i < verLenMin; i++ {
//		if vera[i] > verb[i] {
//			ret = 1
//			break
//		} else if vera[i] < verb[i] {
//			ret = -1
//			break
//		}
//	}
//
//	if ret == 0 {
//		if veraLen > verbLen {
//			ret = 1
//		} else if veraLen < verbLen {
//			ret = -1
//		}
//	}
//
//	return ret
//}

// RandString rand string
func RandString(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// CheckURLValid check the url invalied or not
func CheckURLValid(urlStr string) bool {
	hasDOt := false
	for i, pb := range []byte(urlStr) {
		if pb == '.' {
			if i == 0 {
				continue
			}

			hasDOt = true
		}
		if pb >= '0' && pb <= '9' || pb == ':' { // is IP host
			continue
		}

		if hasDOt == true {
			return true
		}
	}

	return false
}

// IsIP check IP format valid or invalid
func IsIP(IP string) bool {
	return govalidator.IsIP(IP)
}

package main

import (
	"strings"
)

// default path init
var (
	MacOrgFile = "/etc/vgwsys/mac-prefixes"
	MacOrgData []byte
	MacOrgMap  map[string]string
)

func init() {
	MacOrgMap = make(map[string]string)
	initMacOrgMap()
}

// initMacOrgMap init MacOrgMap
func initMacOrgMap() {
	if !utils.PathExist(MacOrgFile) {
		return
	}

	var err error
	MacOrgData, err = utils.ReadFileByte(MacOrgFile)
	if err != nil {
		log.Error("File [%s] read error:%v", MacOrgFile, err)
	}
	macOrgLines := strings.Split(string(MacOrgData), "\n")
	for _, macOrg := range macOrgLines {
		if len(macOrg) < 7 {
			continue
		}
		if string([]byte(strings.TrimSpace(macOrg))[0]) == "#" {
			continue
		}
		fields := strings.SplitN(macOrg, " ", 2)
		if len(fields) != 2 {
			continue
		}

		MacOrgMap[fields[0]] = fields[1]
	}
}

// GetOrgByMAC get orginazation by mac address
// if not found, return void string
func GetOrgByMAC(mac string) string {
	hmac, err := utils.ToHexadecimalMac(mac)
	if err != nil {
		return ""
	}

	macHead := string([]byte(hmac)[:6])
	if v, ok := MacOrgMap[macHead]; ok {
		return v
	}

	return ""
}

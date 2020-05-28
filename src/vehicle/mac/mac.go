package mac

import (
	"strings"
	"vehicle_system/src/vehicle/util"
)

var (
	MacOrgFile = "macs"
	MacOrgData []byte
	MacOrgMap  map[string]string
)

func Setup() {
	MacOrgMap = make(map[string]string)
	initMacOrgMap()
}

// initMacOrgMap init MacOrgMap
func initMacOrgMap() {
	if !util.PathExist(MacOrgFile) {
		return
	}

	var err error
	MacOrgData, err = util.ReadFileByte(MacOrgFile)
	if err != nil {
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
	hmac, err := util.ToHexadecimalMac(mac)
	//hmac := []string{"102C83", "4455B1", "A88038", "F80D60", "A47B85", "D06A1F"}
	if err != nil {
		return ""
	}
	macHead := string([]byte(hmac)[:6])
	if v, ok := MacOrgMap[macHead]; ok {
		return v
	}

	return ""
}

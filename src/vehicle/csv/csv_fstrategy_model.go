package csv

import (
	"strconv"
	"strings"
)

//csv头信息
type CsvFstrategyModelHeader struct {
	VehicleId   string
	FstrategyId string
	Ip          string
	Port        string //通过,分割
}

//csv所有记录
type CsvFstrategyModelBody struct {
	CsvFstrategyModelBody []CsvFstrategyModelHeader
}

func CreateCsvFstrategyHeader() CsvFstrategyModelHeader {
	return CsvFstrategyModelHeader{
		VehicleId:   "VehicleId",
		FstrategyId: "FstrategyId",
		Ip:          "Ip",
		Port:        "Port",
	}
}

func CreateCsvFstrategyBody(vehicleId string, fstrategyId string, ipPortMap map[string][]uint32) CsvFstrategyModelBody {
	var csvFstrategyBody []CsvFstrategyModelHeader
	for ip, ports := range ipPortMap {
		var ipPortSlice []string
		for _, port := range ports {
			portStr := strconv.Itoa(int(port))
			ipPortSlice = append(ipPortSlice, portStr)
		}
		ipPorts := strings.Join(ipPortSlice, ":")
		csvFstrategyBodyTemp := CsvFstrategyModelHeader{
			VehicleId:   vehicleId,
			FstrategyId: fstrategyId,
			Ip:          ip,
			Port:        ipPorts,
		}
		csvFstrategyBody = append(csvFstrategyBody, csvFstrategyBodyTemp)
	}
	csvFstrategyModelBody := CsvFstrategyModelBody{
		CsvFstrategyModelBody: csvFstrategyBody,
	}
	return csvFstrategyModelBody
}

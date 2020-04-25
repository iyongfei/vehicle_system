package csv

import "strconv"

//csv头信息
type CsvFstrategyModelHeader struct {
	VehicleId   string
	FstrategyId string
	Ip          string
	Port        string //通过,分割
}

//csv单条记录
//type CsvFstrategyModelBodyItem struct {
//	CsvFstrategyModelHeader
//}

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
		for _, port := range ports {

			portStr := strconv.Itoa(int(port))
			csvFstrategyBodyTemp := CsvFstrategyModelHeader{
				VehicleId:   vehicleId,
				FstrategyId: fstrategyId,
				Ip:          ip,
				Port:        portStr,
			}
			csvFstrategyBody = append(csvFstrategyBody, csvFstrategyBodyTemp)
		}

	}
	csvFstrategyModelBody := CsvFstrategyModelBody{
		CsvFstrategyModelBody: csvFstrategyBody,
	}
	return csvFstrategyModelBody
}

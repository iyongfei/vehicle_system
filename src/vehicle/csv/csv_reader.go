package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

var csvFormat = fmt.Errorf("%s", "csv format error")

const CsvAddColumn = 3
const CsvEditColumn = 4

const (
	AddCsv = iota
	EditCsv
)

type CsvReaderModel struct {
	csvReader *csv.Reader
	csvFile   *os.File
}

func CreateCsvReader(csvFilePathName string) *CsvReaderModel {

	csvFile := getCsvFile(csvFilePathName)
	csvReader := csv.NewReader(csvFile)

	csvReaderModel := &CsvReaderModel{
		csvReader: csvReader,
		csvFile:   csvFile,
	}

	return csvReaderModel
}

func getCsvFile(csvFilePathName string) *os.File {
	csvFile, csvFileErr := os.OpenFile(csvFilePathName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if csvFileErr != nil {
		logger.Logger.Print("%s newCsvReader err:%+v", util.RunFuncName(), csvFileErr)
		logger.Logger.Info("%s  newCsvReader:%+v", util.RunFuncName(), csvFileErr)
		return nil
	}
	return csvFile

}

func (csvReaderModel *CsvReaderModel) ParseCsvFile(handleCsvMode int) (map[string]map[string][]uint32, error) {
	csvReader := csvReaderModel.csvReader
	defer csvReaderModel.csvFile.Close()

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var filterData map[string]map[string][]uint32
	var filterDataErr error

	if handleCsvMode == AddCsv {
		filterData, filterDataErr = filterCsvData(records, handleCsvMode)
		if filterDataErr != nil {
			return nil, filterDataErr
		}
	} else if handleCsvMode == EditCsv {
		filterData, filterDataErr = filterCsvData(records, handleCsvMode)
		if filterDataErr != nil {
			return nil, filterDataErr
		}
	}
	return filterData, nil
}

/**

VehicleId,FstrategyId,Ip,Port
754d2728b4e549c5a16c0180fcacb800,dwUF8MhOcJDDuXWaDYsQXW1aNtzHSMlp,192.167.1.3,123:125:23
754d2728b4e549c5a16c0180fcacb800,dwUF8MhOcJDDuXWaDYsQXW1aNtzHSMlp,192.168.1.5,123:125:23
*/
func filterCsvData(records [][]string, handleCsvMode int) (map[string]map[string][]uint32, error) {
	diportsMap := map[string]map[string][]uint32{}
	for _, record := range records {
		if handleCsvMode == AddCsv {
			if len(record) != CsvAddColumn {
				return nil, csvFormat
			}
		} else if handleCsvMode == EditCsv {
			if len(record) != CsvEditColumn {
				return nil, csvFormat
			}
		}
	}
	for _, record := range records {
		vehicleId := record[0]
		ip := record[1]
		ports := record[2]

		if len(diportsMap[vehicleId]) == 0 {
			diportsMap[vehicleId] = map[string][]uint32{}

			ipFormat := util.IpFormat(ip)
			if ipFormat {
				if len(diportsMap[vehicleId][ip]) == 0 {
					diportsMap[vehicleId][ip] = []uint32{}
					portSlice := strings.Split(ports, ":")
					for _, port := range portSlice {
						if util.VerifyIpPort(port) {
							portInt, _ := strconv.Atoi(port)
							diportsMap[vehicleId][ip] = append(diportsMap[vehicleId][ip], uint32(portInt))
						}
					}
				}
			}
		} else {
			ipFormat := util.IpFormat(ip)

			if ipFormat {
				if len(diportsMap[vehicleId][ip]) == 0 {
					diportsMap[vehicleId][ip] = []uint32{}
					portSlice := strings.Split(ports, ":")
					mapFilter := map[string]int{}
					for _, port := range portSlice {
						mapFilter[port] = 1
					}
					for mapK, _ := range mapFilter {
						if util.VerifyIpPort(mapK) {

						}
						portInt, _ := strconv.Atoi(mapK)
						diportsMap[vehicleId][ip] = append(diportsMap[vehicleId][ip], uint32(portInt))
					}
				} else {
					portSlice := strings.Split(ports, ":")
					for _, port := range portSlice {
						portInt, _ := strconv.Atoi(port)
						diportsMap[vehicleId][ip] = append(diportsMap[vehicleId][ip], uint32(portInt))
					}
				}
			}

		}
	}
	return diportsMap, nil
}

//
//func ExampleReader() {
//	in := `first_name,last_name,username
//"Rob","Pike",rob
//Ken,Thompson,ken
//"Robert","Griesemer","gri"
//`
//	r := csv.NewReader(strings.NewReader(in))
//
//	for {
//		record, err := r.Read()
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			log.Fatal(err)
//		}
//
//	}
//	// Output:
//	// [first_name last_name username]
//	// [Rob Pike rob]
//	// [Ken Thompson ken]
//	// [Robert Griesemer gri]
//}

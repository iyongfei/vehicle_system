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
	CsvColumnZeroIndex = iota
	CsvColumnTwoIndex
	CsvColumnThreeIndex
	CsvColumnFourIndex
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

func (csvReaderModel *CsvReaderModel) ParseEditCsvFile(fstrategyId string, vehicleId string) (map[string]map[string][]uint32, error) {
	csvReader := csvReaderModel.csvReader
	defer csvReaderModel.csvFile.Close()

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		if len(record) != CsvEditColumn {
			return nil, csvFormat
		}
	}

	filterRecords := [][]string{}
	//过滤非本fstrategyId
	for _, record := range records {
		csvFstrategyId := record[CsvColumnTwoIndex]
		csvVehicleId := record[CsvColumnZeroIndex]

		if util.RrgsTrim(csvFstrategyId) == util.RrgsTrim(fstrategyId) &&
			util.RrgsTrim(csvVehicleId) == util.RrgsTrim(vehicleId) {
			filterRecords = append(filterRecords, record)
		}
	}

	var filterData map[string]map[string][]uint32
	var filterDataErr error

	filterData, filterDataErr = filterEditCsvData(filterRecords)
	if filterDataErr != nil {
		return nil, filterDataErr
	}

	return filterData, nil
}

/**

VehicleId,FstrategyId,Ip,Port
754d2728b4e549c5a16c0180fcacb800,dwUF8MhOcJDDuXWaDYsQXW1aNtzHSMlp,192.167.1.3,123:125:23
754d2728b4e549c5a16c0180fcacb800,dwUF8MhOcJDDuXWaDYsQXW1aNtzHSMlp,192.168.1.5,123:125:23
*/
func filterEditCsvData(records [][]string) (map[string]map[string][]uint32, error) {
	diportsMap := map[string]map[string][]uint32{}
	for _, record := range records {
		if len(record) != CsvEditColumn {
			return nil, csvFormat
		}
	}

	for _, record := range records {
		vehicleId := record[CsvColumnZeroIndex]
		//fstrategyId := record[CsvColumnTwoIndex]
		ip := record[CsvColumnThreeIndex]
		ports := record[CsvColumnFourIndex]

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
func (csvReaderModel *CsvReaderModel) ParseAddCsvFile() (map[string]map[string][]uint32, error) {
	csvReader := csvReaderModel.csvReader
	defer csvReaderModel.csvFile.Close()

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var filterData map[string]map[string][]uint32
	var filterDataErr error

	filterData, filterDataErr = filterAddCsvData(records)
	if filterDataErr != nil {
		return nil, filterDataErr
	}

	return filterData, nil
}

func filterAddCsvData(records [][]string) (map[string]map[string][]uint32, error) {
	diportsMap := map[string]map[string][]uint32{}
	for _, record := range records {
		if len(record) != CsvAddColumn {
			return nil, csvFormat
		}
	}

	for _, record := range records {
		vehicleId := record[CsvColumnZeroIndex]
		ip := record[CsvColumnTwoIndex]
		ports := record[CsvColumnThreeIndex]

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

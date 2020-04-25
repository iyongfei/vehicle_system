package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

const FStrategyCsvFolder = "fstrategy"
const FStrategyCsvSuffix = ".csv"

type CSV struct {
	csvWriter   *csv.Writer
	csvFile     *os.File
	CsvFilePath string
}

func NewCsvWriter(vehicleId string, fstrategyId string) *CSV {

	csvFile, csvFolderFileName, err := GetFstrategyCsvFile(vehicleId, fstrategyId)
	if err != nil {
		logger.Logger.Print("%s newCsvWriter err:%+v", util.RunFuncName(), err)
		logger.Logger.Info("%s  newCsvWritererr:%+v", util.RunFuncName(), err)
	}
	var csvModel *CSV
	csvWriter := csv.NewWriter(csvFile)

	csvModel = &CSV{
		csvWriter:   csvWriter,
		csvFile:     csvFile,
		CsvFilePath: csvFolderFileName,
	}

	return csvModel
}

func (csvModel *CSV) ParseCsvWritData(csvDatas ...interface{}) [][]string {

	fCsvDatas := [][]string{}

	for _, scvData := range csvDatas {
		switch scvData.(type) {
		case CsvFstrategyModelHeader:

			csvHeaderdata := scvData.(CsvFstrategyModelHeader)
			var scvHeader []string
			scvHeader = append(scvHeader,
				csvHeaderdata.VehicleId, csvHeaderdata.FstrategyId,
				csvHeaderdata.Ip, csvHeaderdata.Port)

			fCsvDatas = append(fCsvDatas, scvHeader)

		case CsvFstrategyModelBody:
			csvBodydata := scvData.(CsvFstrategyModelBody)
			for _, csvBodyItem := range csvBodydata.CsvFstrategyModelBody {
				var scvBodyItemTemp []string
				scvBodyItemTemp = append(scvBodyItemTemp,
					csvBodyItem.VehicleId, csvBodyItem.FstrategyId,
					csvBodyItem.Ip, csvBodyItem.Port)

				fCsvDatas = append(fCsvDatas, scvBodyItemTemp)
			}
		}
	}

	return fCsvDatas
}

func (csvModel *CSV) SetCsvWritData(csvDatas ...interface{}) {

	csvData := csvModel.ParseCsvWritData(csvDatas...)
	csvModel.CsvWritData(csvData)

}
func (csvModel *CSV) CsvWritData(csvDatas [][]string) {
	logger.Logger.Print("%s csvWritData err:%+v", util.RunFuncName(), csvDatas)
	logger.Logger.Info("%s  csvWritData:%+v", util.RunFuncName(), csvDatas)

	_, fileWriteErr := csvModel.csvFile.WriteString("\xEF\xBB\xBF")
	if fileWriteErr != nil {
		logger.Logger.Print("%s fileWrite bom err:%+v", util.RunFuncName(), fileWriteErr)
		logger.Logger.Info("%s  fileWrite bom:%+v", util.RunFuncName(), fileWriteErr)

	}

	csvWriteAllerr := csvModel.csvWriter.WriteAll(csvDatas)
	if csvWriteAllerr != nil {
		logger.Logger.Print("%s csvWriteAll err:%+v", util.RunFuncName(), csvWriteAllerr)
		logger.Logger.Info("%s  csvWriteAll err:%+v", util.RunFuncName(), csvWriteAllerr)

	}

	csvModel.csvWriter.Flush()

	csvModel.Close()
}
func (csvModel *CSV) Close() {
	if csvModel != nil {
		if csvModel.csvFile != nil {
			csvModel.csvFile.Close()
		}
	}
}

func CreateCsvFolder() (string, error) {
	wd := Getwd()

	if strings.Trim(wd, " ") == "" {
		return "", fmt.Errorf("%s get_wd:%s null", util.RunFuncName(), wd)

	}
	csvFileFolderPath := fmt.Sprintf("%s/%s", wd, FStrategyCsvFolder)
	csvFileFolderPathErr := MkdirAll(csvFileFolderPath)
	if csvFileFolderPathErr != nil {
		return "", csvFileFolderPathErr
	}
	return csvFileFolderPath, nil
}

//获取fstrategyCsv文件
func GetFstrategyCsvFile(vehicleId string, fstrategyId string) (*os.File, string, error) {
	csvFolderPath, err := CreateCsvFolder()
	if err != nil {
		return nil, "", err
	}
	csvFileName := fmt.Sprintf("%s%s%s", vehicleId, fstrategyId, FStrategyCsvSuffix)
	csvFilePath := fmt.Sprintf("%s/%s", csvFolderPath, csvFileName)

	csvFile, err := os.OpenFile(csvFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)

	if err != nil {
		return nil, "", err
	}
	csvFolderFileName := fmt.Sprintf("%s/%s", FStrategyCsvFolder, csvFileName)
	return csvFile, csvFolderFileName, nil
}

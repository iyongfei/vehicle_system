package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/util"
)

const FStrategyCsvFolder = "fstrategy_csv"
const FStrategyCsvSuffix = ".csv"
const FileTruncate = 0
const FileAppend = 1

type CsvWriter struct {
	csvWriter   *csv.Writer
	csvFile     *os.File
	CsvFilePath string
}

func NewCsvWriter(fstrategyId string, fileMode int) *CsvWriter {

	csvFile, csvFolderFileName, err := GetFstrategyCsvFile(fstrategyId, fileMode)
	if err != nil {
		logger.Logger.Print("%s newCsvWriter err:%+v", util.RunFuncName(), err)
		logger.Logger.Info("%s  newCsvWritererr:%+v", util.RunFuncName(), err)
	}
	var csvModel *CsvWriter
	csvWriter := csv.NewWriter(csvFile)

	csvModel = &CsvWriter{
		csvWriter:   csvWriter,
		csvFile:     csvFile,
		CsvFilePath: csvFolderFileName,
	}

	return csvModel
}

//获取fstrategyCsv文件
func GetFstrategyCsvFile(fstrategyId string, fileMode int) (*os.File, string, error) {
	csvFolderPath, err := CreateCsvFolder()
	if err != nil {
		return nil, "", err
	}
	csvFileName := fmt.Sprintf("%s%s", fstrategyId, FStrategyCsvSuffix)
	csvFilePath := fmt.Sprintf("%s/%s", csvFolderPath, csvFileName)

	var csvFile *os.File
	var csvFileErr error
	if fileMode == FileTruncate {
		csvFile, csvFileErr = os.OpenFile(csvFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	} else if fileMode == FileAppend {
		csvFile, csvFileErr = os.OpenFile(csvFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	} else {
		csvFile, csvFileErr = os.OpenFile(csvFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	}

	if csvFileErr != nil {
		return nil, "", csvFileErr
	}
	localHost := util.GetLocalHost()
	//http://192.168.100.2:7001/fstrategy/754d2728b4e549c5a16c0180fcacb800_LDmpxpPaarSHf2dvgVjQWNJHTewnGXEz.csv
	csvFolderFileName := fmt.Sprintf("%s/%s/%s", localHost, FStrategyCsvFolder, csvFileName)
	return csvFile, csvFolderFileName, nil
}

func (csvModel *CsvWriter) ParseCsvWritData(csvDatas ...interface{}) [][]string {

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

func (csvModel *CsvWriter) SetCsvWritData(csvDatas ...interface{}) {

	csvData := csvModel.ParseCsvWritData(csvDatas...)
	csvModel.CsvWritData(csvData)

}
func (csvModel *CsvWriter) CsvWritData(csvDatas [][]string) {
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
func (csvModel *CsvWriter) Close() {
	if csvModel != nil {
		if csvModel.csvFile != nil {
			csvModel.csvFile.Close()
		}
	}
}

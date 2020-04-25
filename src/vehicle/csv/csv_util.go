package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"vehicle_system/src/vehicle/util"
)

type CSV struct {
	csvWriter *csv.Writer
	csvFile   *os.File
}

func NewCsvWriter(vehicleId string, fstrategyId string) *CSV {

	csvFile, _ := GetFstrategyCsvFile(vehicleId, fstrategyId)
	csvWriter := csv.NewWriter(csvFile)

	csvModel := &CSV{
		csvWriter: csvWriter,
		csvFile:   csvFile,
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

	csvData := csvModel.ParseCsvWritData(csvDatas)

	_, fileWriteErr := csvModel.csvFile.WriteString("\xEF\xBB\xBF")
	fmt.Println("csvFile WriteString err", fileWriteErr)
	csvWriteAllerr := csvModel.csvWriter.WriteAll(csvData)
	fmt.Println("csvWriteAllerr", csvWriteAllerr)
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

func CreateCSV(vehicleId string, fstrategyId string) (string, error) {
	wd := Getwd()
	if strings.Trim(wd, " ") == "" {
		return "", fmt.Errorf("%s get_wd:%s null", util.RunFuncName(), wd)

	}
	csvFilePath := fmt.Sprintf("%s%s%s", wd, vehicleId, fstrategyId)
	csvFilePathErr := MkdirAll(csvFilePath)
	if csvFilePathErr != nil {
		return csvFilePath, csvFilePathErr
	}
	return "", nil
}

//获取fstrategyCsv文件
func GetFstrategyCsvFile(vehicleId string, fstrategyId string) (*os.File, error) {
	csvFilePath, err := CreateCSV(vehicleId, fstrategyId)
	if err != nil {
		return nil, err
	}

	csvFile, err := os.OpenFile(csvFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)

	if err != nil {
		return nil, err
	}
	return csvFile, nil
}

func ExampleWriter() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	// Output:
	// first_name,last_name,username
	// Rob,Pike,rob
	// Ken,Thompson,ken
	// Robert,Griesemer,gri
}

func ExampleWriter_WriteAll() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
	// Output:
	// first_name,last_name,username
	// Rob,Pike,rob
	// Ken,Thompson,ken
	// Robert,Griesemer,gri
}

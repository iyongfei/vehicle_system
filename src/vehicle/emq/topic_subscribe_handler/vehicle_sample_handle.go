package topic_subscribe_handler

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"vehicle_system/src/vehicle/emq/protobuf"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/model/model_base"
	"vehicle_system/src/vehicle/util"
)

func HandleVehicleSample(vehicleResult protobuf.GWResult, vehicleId string) error {
	//parse
	sampleParam := &protobuf.SampleParam{}
	err := proto.Unmarshal(vehicleResult.GetParam(), sampleParam)
	if err != nil {
		logger.Logger.Print("%s unmarshal sampleParam err:%s", util.RunFuncName(), err.Error())
		logger.Logger.Error("%s unmarshal sampleParam err:%s", util.RunFuncName(), err.Error())
		return fmt.Errorf("%s unmarshal sampleParam err:%s", util.RunFuncName(), err.Error())
	}

	logger.Logger.Print("%s unmarshal sampleParam:%+v", util.RunFuncName(), sampleParam)
	logger.Logger.Info("%s unmarshal sampleParam:%+v", util.RunFuncName(), sampleParam)
	//create
	sampleInfo := &model.Sample{
		SampleId:  sampleParam.GetId(),
		VehicleId: vehicleId,
	}
	modelBase := model_base.ModelBaseImpl(sampleInfo)

	_, recordNotFound := modelBase.GetModelByCondition("sample_id = ? and vehicle_id = ?",
		[]interface{}{sampleInfo.SampleId, sampleInfo.VehicleId}...)

	modelBase.CreateModel(sampleParam)
	if recordNotFound {
		if err := modelBase.InsertModel(); err != nil {
			return fmt.Errorf("%s insert vehicle sample err:%s", util.RunFuncName(), err.Error())
		}
	} else {
		attrs := map[string]interface{}{
			"start_time":      sampleInfo.StartTime,
			"remain_time":     sampleInfo.RemainTime,
			"total_time":      sampleInfo.TotalTime,
			"status":          sampleInfo.Status,
			"timeout":         sampleInfo.Timeout,
			"name":            sampleInfo.Name,
			"introduce":       sampleInfo.Introduce,
			"check":           sampleInfo.Check,
			"study_origin_id": sampleInfo.StudyOriginId,
		}
		if err := modelBase.UpdateModelsByCondition(attrs, "sample_id = ? and vehicle_id = ?",
			[]interface{}{sampleInfo.SampleId, sampleInfo.VehicleId}...); err != nil {
			return fmt.Errorf("%s update vehicle sample err:%s", util.RunFuncName(), err.Error())
		}
	}

	switch sampleParam.GetStatus() {
	case protobuf.SampleParam_COLLECT_OK:
		for _, sampleItem := range sampleParam.GetSampleItem() {

			sampleItemModel := &model.SampleItem{
				SampleItemId: util.RandomString(32),
				SampleId:     sampleParam.GetId(),
			}
			modelBase := model_base.ModelBaseImpl(sampleItemModel)

			_, recordNotFound := modelBase.GetModelByCondition("sample_id = ? and sample_item_id = ?",
				[]interface{}{sampleItemModel.SampleId, sampleItemModel.SampleItemId}...)

			modelBase.CreateModel(sampleItem)
			if recordNotFound {
				if err := modelBase.InsertModel(); err != nil {
					continue
					//return fmt.Errorf("%s insert vehicle sample item err:%s",util.RunFuncName(),err.Error())
				}
			}

		}
	case protobuf.SampleParam_STATUSDEFAULT:
	case protobuf.SampleParam_COLLECTING:
	case protobuf.SampleParam_COLLECT_FAILED:

	default:

	}
	//for _, threatItem := range vehicleParam.GetThreatItem() {
	//	threat := &model.Threat{
	//		ThreatId:util.RandomString(32),
	//		VehicleId:vehicleId,
	//	}
	//
	//	modelBase := model_base.ModelBaseImpl(threat)
	//	modelBase.CreateModel(threatItem)
	//	_, recordNotFound := modelBase.GetModelByCondition("threat_id = ?", threat.ThreatId)
	//	if !recordNotFound {
	//		continue
	//	}
	//
	//
	//	err := modelBase.InsertModel()
	//	if err != nil {
	//		continue
	//	}
	//}
	return nil
}

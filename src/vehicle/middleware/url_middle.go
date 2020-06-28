package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"vehicle_system/src/vehicle/auth"
	"vehicle_system/src/vehicle/db/mysql"
	"vehicle_system/src/vehicle/logger"
	"vehicle_system/src/vehicle/model"
	"vehicle_system/src/vehicle/response"
	"vehicle_system/src/vehicle/util"
)

func UrlMiddle() gin.HandlerFunc {
	return urlMiddleHandlerFunc
}

func urlMiddleHandlerFunc(c *gin.Context) {

	//未授权
	vehicleIds := []string{}
	_ = mysql.QueryPluckByModelWhere(&model.VehicleInfo{}, "vehicle_id", &vehicleIds,
		"", []interface{}{}...)

	unAuth := auth.VehicleAllUnAuth(vehicleIds)
	if unAuth {
		ret := response.StructResponseObj(response.VStatusUnauthorized, response.Unauthorized, "")
		c.JSON(http.StatusOK, ret)
		c.Abort()
		logger.Logger.Error("%s auth file:%s not exist", util.RunFuncName(), auth.AuthFile)
		logger.Logger.Print("%s auth file:%s not exist", util.RunFuncName(), auth.AuthFile)
		return
	}
	//授权过期
	authExpire := auth.AuthVehicleAllExpire()
	if authExpire {
		ret := response.StructResponseObj(response.VStatusExpiredUnauthorized, response.AuthorizedExpire, "")
		c.JSON(http.StatusOK, ret)
		c.Abort()
		logger.Logger.Error("%s auth file:%s not exist", util.RunFuncName(), auth.AuthFile)
		logger.Logger.Print("%s auth file:%s not exist", util.RunFuncName(), auth.AuthFile)
		return
	}

	///api/v1/vehicles/47aa3baa36d94269b9d32ed9bcf8a2f4
	path := c.Request.URL.Path
	method := c.Request.Method

	urlGetSlice := UrlGetSlice()
	urlPutSlice := UrlPutSlice()
	urlPostSlice := UrlPostSlice()

	if method == "GET" {
		for _, url := range urlGetSlice {
			index := strings.Index(path, url)

			if index != -1 {
				qVehicleId := c.Query("vehicle_id")
				pVehicleId := c.Param("vehicle_id")
				fVehicle := ""

				if qVehicleId != "" {
					fVehicle = qVehicleId
				}
				if pVehicleId != "" {
					fVehicle = pVehicleId
				}

				if fVehicle != "" {

					//判断是否是为授权
					if !auth.VehicleAuth(fVehicle) {
						ret := response.StructResponseObj(response.VStatusUnauthorized, response.Unauthorized, "")
						c.JSON(http.StatusOK, ret)
						c.Abort()
						logger.Logger.Error("%s auth file:%s not exist", util.RunFuncName(), auth.AuthFile)
						logger.Logger.Print("%s auth file:%s not exist", util.RunFuncName(), auth.AuthFile)
						return
					}
				}
			}
		}
	} else if method == "PUT" {
		for _, url := range urlPutSlice {
			index := strings.Index(path, url)
			if index != -1 {
				qVehicleId := c.PostForm("vehicle_id")
				pVehicleId := c.Param("vehicle_id")
				fVehicle := ""

				if qVehicleId != "" {
					fVehicle = qVehicleId
				}
				if pVehicleId != "" {
					fVehicle = pVehicleId
				}

				if fVehicle != "" {
					//判断是否是为授权
					if !auth.VehicleAuth(fVehicle) {
						ret := response.StructResponseObj(response.VStatusUnauthorized, response.Unauthorized, "")
						c.JSON(http.StatusOK, ret)
						c.Abort()
						logger.Logger.Error("%s auth file:%s not exist", util.RunFuncName(), auth.AuthFile)
						logger.Logger.Print("%s auth file:%s not exist", util.RunFuncName(), auth.AuthFile)
						return
					}
				}
			}
		}

	} else if method == "DELETE" {

	} else if method == "POST" {
		for _, url := range urlPostSlice {
			index := strings.Index(path, url)
			if index != -1 {
				qVehicleId := c.PostForm("vehicle_id")
				pVehicleId := c.Param("vehicle_id")
				fVehicle := ""

				if qVehicleId != "" {
					fVehicle = qVehicleId
				}
				if pVehicleId != "" {
					fVehicle = pVehicleId
				}
				if fVehicle != "" {
					//判断是否是为授权
					if !auth.VehicleAuth(fVehicle) {
						ret := response.StructResponseObj(response.VStatusUnauthorized, response.Unauthorized, "")
						c.JSON(http.StatusOK, ret)
						c.Abort()
						logger.Logger.Error("%s auth file:%s not exist", util.RunFuncName(), auth.AuthFile)
						logger.Logger.Print("%s auth file:%s not exist", util.RunFuncName(), auth.AuthFile)
						return
					}
				}

			}

		}

	}

	c.Next()
}

//dele
//"fstrategys",

//post
func UrlPostSlice() []string {

	urlPostSlice := []string{
		"fstrategys",
	}
	return urlPostSlice
}

//put
///assets/:asset_id
//
func UrlPutSlice() []string {

	urlPutSlice := []string{
		"fstrategys",
		"fstrategy_csvs",
		"vehicles",
		"vehicle_info",
		"asset_info",
	}
	return urlPutSlice
}

/**
get
pagination/fstrategys do
pagination/vehicles do
fstrategy_csvs/:fstrategy_id do 半do
pagination/assets do
all/assets do
assets/:asset_id do
*/
/**
获取urlmap

*/
func UrlGetSlice() []string {

	urlGetSlice := []string{
		"flow_statistics",
		"monitors",
		"fstrategys",
		"active/fstrategys",
		"pagination/vehicle/fstrategys",
		"flows",
		"pagination/flows", //可以为空
		"flow_type_counts",
		"tflow_dps",
		"tflows",
		"flow_dps",
		"vehicles",
	}
	return urlGetSlice
}

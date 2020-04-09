package model

import "github.com/jinzhu/gorm"

type Asset struct {
	gorm.Model
	VehicleId       						string 						//关联的小v ID
	AssetId  						string 			`gorm:"unique"`//资产id

	IP         			string       			//小v资产 IP
	Mac        				string       		//资产Mac地址
	Name       				string 					//小v资产名称
	TradeMark                    string //资产品牌

	OnlineStatus      		bool 	//在线状态
	LastOnline        uint32 //最近活跃时间

	InternetSwitch                bool//是否允许联网
	ProtectStatus   		bool //是否受小V保护
	LanVisitSwitch                bool//是否可以访问内网

	//GwAssetCategory                 int  //资产类别
	AssetGroup                    string
	AssetLeader					string
}
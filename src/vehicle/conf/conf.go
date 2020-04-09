package conf

import (
	"time"
	"vehicle_system/src/vehicle/logger"
)

const (
	CONF_SOURCE = "conf.ini"
)
var (
	//domain
	ServerHost           string
	LocalHost           string

	//server_port
	ServerPort           uint32

	//mysql
	MysqlUser      string
	MysqlPassword  string
	MysqlDbname    string
	MysqlPort    uint32
	MaxIdleConns    int
	MaxOpenConns    int

	//redis
	SessionRedisAddr string
	SessionRedisPassword string
	SessionRedisDB     int

	//emq
	EmqBrokerUrl string
	EmqClientId string
	EmqKeepAlive time.Duration
	AutoReconnect bool
	ConnectTimeOut time.Duration

	EmqPingTimeOut time.Duration
	EmqCleanSession bool
	EmqUser string
	EmqPassword string
	MaxReconnectInterval time.Duration
	EmqQos uint32
	PublishChanCapa uint32
	PublishChanIdle time.Duration
	SubscribeChanCapa uint32
	SubscribeChanIdle time.Duration

	//jwt
	Expires time.Duration
	SignKey string
)

func Setup()  {
	iniParser := IniParser{}
	if err:=iniParser.Load(CONF_SOURCE);err!=nil{
		logger.Logger.Error("iniParser load ini err:%+v",err)
		return
	}

	//host port
	ServerHost = iniParser.GetString("domain","server_host")
	LocalHost = iniParser.GetString("domain","local_host")
	ServerPort = iniParser.GetUint32("server_port","server_port")

	logger.Logger.Info("server_host:%s,localhost:%s,server_port:%d",ServerHost,LocalHost,ServerPort)
	logger.Logger.Print("server_host:%s,localhost:%s,server_port:%d",ServerHost,LocalHost,ServerPort)

	//mysql
	MysqlUser = iniParser.GetString("mysql","user_name")
	MysqlPassword = iniParser.GetString("mysql","password")
	MysqlDbname = iniParser.GetString("mysql","db_name")
	MysqlPort = iniParser.GetUint32("mysql","mysql_port")
	MaxIdleConns = iniParser.GetInt("mysql","max_idle_conns")
	MaxOpenConns = iniParser.GetInt("mysql","max_open_oonns")
	logger.Logger.Info("server_host:%s,localhost:%s,server_port:%d",ServerHost,LocalHost,ServerPort)

	logger.Logger.Info("user_name:%s,password:%s,db_name:%s,mysql_port:%d,max_idle_conns:%d,max_open_oonns:%d",
		MysqlUser,MysqlPassword,MysqlDbname,MysqlPort,MaxIdleConns,MaxOpenConns)
	logger.Logger.Print("user_name:%s,password:%s,db_name:%s,mysql_port:%d,max_idle_conns:%d,max_open_oonns:%d",
		MysqlUser,MysqlPassword,MysqlDbname,MysqlPort,MaxIdleConns,MaxOpenConns)

	//redis
	SessionRedisAddr = iniParser.GetString("redis","session_redis_address")
	SessionRedisPassword = iniParser.GetString("redis","session_redis_password")
	SessionRedisDB = iniParser.GetInt("redis","session_redis_db")

	logger.Logger.Info("redis_addr:%s,redis_password:%s,redis_db:%d",SessionRedisAddr,SessionRedisPassword,SessionRedisDB)
	logger.Logger.Print("redis_addr:%s,redis_password:%s,redis_db:%d",SessionRedisAddr,SessionRedisPassword,SessionRedisDB)

	//emq
	EmqBrokerUrl = iniParser.GetString("emq","broker_url")
	EmqClientId = iniParser.GetString("emq","client_id")
	EmqKeepAlive = iniParser.GetTimeDuration("emq","keep_alive")
	AutoReconnect = iniParser.GetBool("emq","auto_reconnect")
	ConnectTimeOut = iniParser.GetTimeDuration("emq","connect_time_out")
	EmqPingTimeOut = iniParser.GetTimeDuration("emq","ping_time_out")
	EmqCleanSession = iniParser.GetBool("emq","clean_session")
	EmqUser = iniParser.GetString("emq","username")
	EmqPassword = iniParser.GetString("emq","password")
	MaxReconnectInterval = iniParser.GetTimeDuration("emq","max_reconnect_interval")
	EmqQos = iniParser.GetUint32("emq","qos")
	PublishChanCapa = iniParser.GetUint32("emq","publish_chan_capa")
	PublishChanIdle = iniParser.GetTimeDuration("emq","publish_chan_idle")
	SubscribeChanCapa = iniParser.GetUint32("emq","subscribe_chan_capa")
	SubscribeChanIdle = iniParser.GetTimeDuration("emq","subscribe_chan_idle")

	logger.Logger.Info("broder_url:%s,client_id:%s,keep_alive:%d,auto_reconnect:%v,connect_time_out:%d,ping_time_out:%d,clean_session:%v,username:%s," +
		"password:%s,max_reconnect_interval:%d,qos:%d,publish_chan_capa:%d,publish_chan_idle:%d,subscribe_chan_capa:%d,subscribe_chan_idle:%d\n",
		EmqBrokerUrl,EmqClientId,EmqKeepAlive,AutoReconnect,ConnectTimeOut,EmqPingTimeOut,EmqCleanSession,
		EmqUser,EmqPassword,MaxReconnectInterval,EmqQos,PublishChanCapa,PublishChanIdle,SubscribeChanCapa,SubscribeChanIdle)

	logger.Logger.Print("broder_url:%s,client_id:%s,keep_alive:%d,auto_reconnect:%v,connect_time_out:%d,ping_time_out:%d,clean_session:%v,username:%s," +
		"password:%s,max_reconnect_interval:%d,qos:%d,publish_chan_capa:%d,publish_chan_idle:%d,subscribe_chan_capa:%d,subscribe_chan_idle:%d\n",
		EmqBrokerUrl,EmqClientId,EmqKeepAlive,AutoReconnect,ConnectTimeOut,EmqPingTimeOut,EmqCleanSession,
		EmqUser,EmqPassword,MaxReconnectInterval,EmqQos,PublishChanCapa,PublishChanIdle,SubscribeChanCapa,SubscribeChanIdle)

	Expires = iniParser.GetTimeDuration("jwt","expires")
	SignKey = iniParser.GetString("jwt","sign_key")

	logger.Logger.Info("Expires:%d,SignKey:%s",Expires,SignKey)
	logger.Logger.Print("Expires:%d,SignKey:%s",Expires,SignKey)
}
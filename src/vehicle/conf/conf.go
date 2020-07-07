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
	ServerHost string
	LocalHost  string

	//server_port
	ServerPort uint32

	//push_host
	PushHost string
	PushPort uint32

	//mysql
	MysqlUser     string
	MysqlPassword string
	MysqlDbname   string
	MysqlPort     uint32
	MaxIdleConns  int
	MaxOpenConns  int

	//redis
	SessionRedisAddr     string
	SessionRedisPassword string
	SessionRedisDB       int

	//emq
	EmqBrokerUrl   string
	EmqClientId    string
	EmqKeepAlive   time.Duration
	AutoReconnect  bool
	ConnectTimeOut time.Duration

	EmqPingTimeOut       time.Duration
	EmqCleanSession      bool
	EmqUser              string
	EmqPassword          string
	MaxReconnectInterval time.Duration
	EmqQos               uint32
	PublishChanCapa      uint32
	PublishChanIdle      time.Duration
	SubscribeChanCapa    uint32
	SubscribeChanIdle    time.Duration

	//jwt
	SignKey string

	//fp
	CollectTime  uint32
	ProtoCount   uint64
	CollectBytes uint64

	CollectTimeRate  float64
	ProtoCountRate   float64
	CollectBytesRate float64
	//CollectHostRate  float64
	//CollectTlsRate   float64
	MinRate float64

	//collect_bytes=1048576
	//collect_bytes_rate=0.2

	//权重占比
	MainProtoWeight  float64
	ProtosKindWeight float64
	HostnameWeight   float64
	MacWeight        float64
	TypeWeight       float64
	TlsWeight        float64
	MinRateWeight    float64
)

func Setup() {
	iniParser := IniParser{}
	if err := iniParser.Load(CONF_SOURCE); err != nil {
		logger.Logger.Error("iniParser load ini err:%+v", err)
		return
	}

	//host port
	ServerHost = iniParser.GetString("domain", "server_host")
	LocalHost = iniParser.GetString("domain", "local_host")
	ServerPort = iniParser.GetUint32("server_port", "server_port")

	//push
	PushHost = iniParser.GetString("push", "push_host")
	PushPort = iniParser.GetUint32("push", "push_port")

	logger.Logger.Info("server_host:%s,localhost:%s,server_port:%d,pushhost:%s,push_port:%d", ServerHost, LocalHost, ServerPort, PushHost, PushPort)
	logger.Logger.Print("server_host:%s,localhost:%s,server_port:%d,pushhost:%s,push_port:%d", ServerHost, LocalHost, ServerPort, PushHost, PushPort)

	//mysql
	MysqlUser = iniParser.GetString("mysql", "user_name")
	MysqlPassword = iniParser.GetString("mysql", "password")
	MysqlDbname = iniParser.GetString("mysql", "db_name")
	MysqlPort = iniParser.GetUint32("mysql", "mysql_port")
	MaxIdleConns = iniParser.GetInt("mysql", "max_idle_conns")
	MaxOpenConns = iniParser.GetInt("mysql", "max_open_oonns")
	logger.Logger.Info("server_host:%s,localhost:%s,server_port:%d", ServerHost, LocalHost, ServerPort)

	logger.Logger.Info("user_name:%s,password:%s,db_name:%s,mysql_port:%d,max_idle_conns:%d,max_open_oonns:%d",
		MysqlUser, MysqlPassword, MysqlDbname, MysqlPort, MaxIdleConns, MaxOpenConns)
	logger.Logger.Print("user_name:%s,password:%s,db_name:%s,mysql_port:%d,max_idle_conns:%d,max_open_oonns:%d",
		MysqlUser, MysqlPassword, MysqlDbname, MysqlPort, MaxIdleConns, MaxOpenConns)

	//redis
	SessionRedisAddr = iniParser.GetString("redis", "session_redis_address")
	SessionRedisPassword = iniParser.GetString("redis", "session_redis_password")
	SessionRedisDB = iniParser.GetInt("redis", "session_redis_db")

	logger.Logger.Info("redis_addr:%s,redis_password:%s,redis_db:%d", SessionRedisAddr, SessionRedisPassword, SessionRedisDB)
	logger.Logger.Print("redis_addr:%s,redis_password:%s,redis_db:%d", SessionRedisAddr, SessionRedisPassword, SessionRedisDB)

	//emq
	EmqBrokerUrl = iniParser.GetString("emq", "broker_url")
	EmqClientId = iniParser.GetString("emq", "client_id")
	EmqKeepAlive = iniParser.GetTimeDuration("emq", "keep_alive")
	AutoReconnect = iniParser.GetBool("emq", "auto_reconnect")
	ConnectTimeOut = iniParser.GetTimeDuration("emq", "connect_time_out")
	EmqPingTimeOut = iniParser.GetTimeDuration("emq", "ping_time_out")
	EmqCleanSession = iniParser.GetBool("emq", "clean_session")
	EmqUser = iniParser.GetString("emq", "username")
	EmqPassword = iniParser.GetString("emq", "password")
	MaxReconnectInterval = iniParser.GetTimeDuration("emq", "max_reconnect_interval")
	EmqQos = iniParser.GetUint32("emq", "qos")
	PublishChanCapa = iniParser.GetUint32("emq", "publish_chan_capa")
	PublishChanIdle = iniParser.GetTimeDuration("emq", "publish_chan_idle")
	SubscribeChanCapa = iniParser.GetUint32("emq", "subscribe_chan_capa")
	SubscribeChanIdle = iniParser.GetTimeDuration("emq", "subscribe_chan_idle")

	logger.Logger.Info("broder_url:%s,client_id:%s,keep_alive:%d,auto_reconnect:%v,connect_time_out:%d,ping_time_out:%d,clean_session:%v,username:%s,"+
		"password:%s,max_reconnect_interval:%d,qos:%d,publish_chan_capa:%d,publish_chan_idle:%d,subscribe_chan_capa:%d,subscribe_chan_idle:%d\n",
		EmqBrokerUrl, EmqClientId, EmqKeepAlive, AutoReconnect, ConnectTimeOut, EmqPingTimeOut, EmqCleanSession,
		EmqUser, EmqPassword, MaxReconnectInterval, EmqQos, PublishChanCapa, PublishChanIdle, SubscribeChanCapa, SubscribeChanIdle)

	logger.Logger.Print("broder_url:%s,client_id:%s,keep_alive:%d,auto_reconnect:%v,connect_time_out:%d,ping_time_out:%d,clean_session:%v,username:%s,"+
		"password:%s,max_reconnect_interval:%d,qos:%d,publish_chan_capa:%d,publish_chan_idle:%d,subscribe_chan_capa:%d,subscribe_chan_idle:%d\n",
		EmqBrokerUrl, EmqClientId, EmqKeepAlive, AutoReconnect, ConnectTimeOut, EmqPingTimeOut, EmqCleanSession,
		EmqUser, EmqPassword, MaxReconnectInterval, EmqQos, PublishChanCapa, PublishChanIdle, SubscribeChanCapa, SubscribeChanIdle)

	//jwt
	SignKey = iniParser.GetString("jwt", "sign_key")

	logger.Logger.Info("SignKey:%s", SignKey)
	logger.Logger.Print("SignKey:%s", SignKey)

	//fp

	CollectTime = iniParser.GetUint32("fp", "collect_time")
	ProtoCount = iniParser.GetUint64("fp", "proto_count")
	CollectBytes = iniParser.GetUint64("fp", "collect_bytes")

	logger.Logger.Info("collect_time:%d,proto_count:%d,collect_total:%d", CollectTime, ProtoCount, CollectBytes)
	logger.Logger.Print("collect_time:%d,proto_count:%d,collect_total:%d", CollectTime, ProtoCount, CollectBytes)

	CollectTimeRate = iniParser.GetFloat64("fp", "collect_time_rate")
	ProtoCountRate = iniParser.GetFloat64("fp", "proto_count_rate")
	CollectBytesRate = iniParser.GetFloat64("fp", "collect_bytes_rate")
	//CollectHostRate = iniParser.GetFloat64("fp", "collect_host_rate")
	//CollectTlsRate = iniParser.GetFloat64("fp", "collect_tls_rate")
	MinRate = iniParser.GetFloat64("fp", "min_rate")

	logger.Logger.Info("collect_time_rate:%f,proto_count_rate:%f,collect_total_rate:%f,min_rate:%f",
		CollectTimeRate, ProtoCountRate, CollectBytesRate, MinRate)
	logger.Logger.Print("collect_time_rate:%f,proto_count_rate:%f,collect_total_rate:%f,min_rate:%f",
		CollectTimeRate, ProtoCountRate, CollectBytesRate, MinRate)

	MainProtoWeight = iniParser.GetFloat64("fp", "main_proto_weight")
	ProtosKindWeight = iniParser.GetFloat64("fp", "protos_kind_weight")
	HostnameWeight = iniParser.GetFloat64("fp", "hostname_weight")
	MacWeight = iniParser.GetFloat64("fp", "mac_weight")
	TypeWeight = iniParser.GetFloat64("fp", "type_weight")
	TlsWeight = iniParser.GetFloat64("fp", "tls_weight")
	MinRateWeight = iniParser.GetFloat64("fp", "min_rate_weight")

	logger.Logger.Info("main_proto_weight:%f,protos_kind_weight:%f,hostname_weight:%f,mac_weight:%f,type_weight:%f,tls_weight:%f,MinRateWeight:%f",
		MainProtoWeight, ProtosKindWeight, HostnameWeight, MacWeight, TypeWeight, TlsWeight, MinRateWeight)

	logger.Logger.Print("main_proto_weight:%f,protos_kind_weight:%f,hostname_weight:%f,mac_weight:%f,type_weight:%f,tls_weight:%f,MinRateWeight:%f",
		MainProtoWeight, ProtosKindWeight, HostnameWeight, MacWeight, TypeWeight, TlsWeight, MinRateWeight)

}

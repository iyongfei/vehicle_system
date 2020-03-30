package conf


const (
	CONF_SOURCE = "conf.ini"
)


func init()  {
	iniParser := IniParser{}

	if err:=iniParser.Load(CONF_SOURCE);err!=nil{
		return
	}
}
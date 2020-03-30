package conf

import (
	"gopkg.in/ini.v1"
	"time"
)

type IniParser struct {
	conf_reader *ini.File
}

func (this *IniParser)Load(ini_name string)error  {
	conf,err:=ini.Load(ini_name)

	if err!= nil{
		this.conf_reader = nil
		return err
	}
	this.conf_reader = conf
	return nil
}

func (this *IniParser)GetString(section string,key string)string{
	if this.conf_reader == nil{
		return  ""
	}
	sec:= this.conf_reader.Section(section)

	if sec == nil{
		return  ""
	}
	return sec.Key(key).String()
}


func (this *IniParser)GetBool(section string,key string)bool{
	if this.conf_reader == nil{
		return  false
	}
	sec:= this.conf_reader.Section(section)

	if sec == nil{
		return  false
	}
	value,_:=sec.Key(key).Bool()
	return value
}


func (this *IniParser)GetInt(section string,key string)int  {
	if this.conf_reader == nil{
		return 0
	}
	sec:=this.conf_reader.Section(section)
	if sec == nil{
		return 0
	}

	value,_:=sec.Key(key).Int()
	return value
}



func (this *IniParser)GetUint32(section string,key string)uint32  {
	if this.conf_reader == nil{
		return 0
	}
	sec:=this.conf_reader.Section(section)
	if sec == nil{
		return 0
	}

	value,_:=sec.Key(key).Uint()
	return uint32(value)
}


func (this *IniParser) GetUint64(section string, key string) uint64 {
	if this.conf_reader == nil {
		return 0
	}

	s := this.conf_reader.Section(section)
	if s == nil {
		return 0
	}

	value_int, _ := s.Key(key).Uint64()
	return value_int
}


func (this *IniParser) GetTimeDuration(section string, key string) time.Duration {
	if this.conf_reader == nil {
		return 0
	}

	s := this.conf_reader.Section(section)
	if s == nil {
		return 0
	}

	value_int, _ := s.Key(key).Int64()
	valueDuration := time.Duration(value_int)
	return valueDuration
}

func (this *IniParser) GetInt64(section string, key string) int64 {
	if this.conf_reader == nil {
		return 0
	}

	s := this.conf_reader.Section(section)
	if s == nil {
		return 0
	}

	value_int, _ := s.Key(key).Int64()
	return value_int
}
func (this *IniParser) GetFloat32(section string, key string) float32 {
	if this.conf_reader == nil {
		return 0
	}

	s := this.conf_reader.Section(section)
	if s == nil {
		return 0
	}

	value_float, _ := s.Key(key).Float64()
	return float32(value_float)
}

func (this *IniParser) GetFloat64(section string, key string) float64 {
	if this.conf_reader == nil {
		return 0
	}

	s := this.conf_reader.Section(section)
	if s == nil {
		return 0
	}

	value_float, _ := s.Key(key).Float64()
	return value_float
}

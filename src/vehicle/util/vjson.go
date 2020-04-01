package util

func GetInt(j map[string]interface{}, key string, defaultVal int) int {
	val, ok := j[key].(float64)
	if !ok {
		return defaultVal
	}
	return int(val)
}

func GetFloat(j map[string]interface{}, key string, defaultVal float64) float64 {
	val, ok := j[key].(float64)
	if !ok {
		return defaultVal
	}
	return val
}

func GetString(j map[string]interface{}, key string, defaultVal string) string {
	val, ok := j[key].(string)
	if !ok {
		return defaultVal
	}
	return val
}

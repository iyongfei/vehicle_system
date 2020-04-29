package emq_cmd

import (
	"bytes"
	"strings"
)

func createCmdId(args ...string) string {
	var buffer bytes.Buffer
	for _, arg := range args {
		if strings.Trim(arg, " ") != "" {
			buffer.WriteString(arg)
		}
	}
	return buffer.String()
}

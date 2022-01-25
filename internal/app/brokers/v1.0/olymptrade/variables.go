package olymptrade

import (
	"fmt"
)

const (
	ColorDefault = "\x1b[39m"

	ColorRed   = "\x1b[91m"
	ColorGreen = "\x1b[32m"
	ColorBlue  = "\x1b[94m"

	messageCandleStream = "[{\"t\":2,\"e\":10,\"uuid\":\"%s\",\"d\":[{\"pair\":\"%s\",\"size\":60,\"to\":%d,\"solid\":true}]}]"
	messageBuy          = "[{\"t\":2,\"e\":23,\"uuid\":\"%s\",\"d\":[{\"amount\":%d,\"dir\":\"%s\",\"pair\":\"%s\",\"cat\":\"digital\",\"pos\":0,\"source\":\"platform\",\"account_id\":%s,\"group\":\"%s\",\"timestamp\":%d,\"risk_free_id\":null,\"duration\":60}]}]"
)


func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func blue(s string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, s, ColorDefault)
}